package workerpool

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/officialHaze/go-pool/logger"
)

type WorkerJob func() error

type WorkerPool struct {
	size     int
	jobs     chan WorkerJob
	errchan  chan error
	quit     chan interface{}
	closeJob chan string // the job id needs to be passed in the channel to close it
	allerrs  []error
	mu       sync.Mutex
	wg       *sync.WaitGroup
	onceDo   sync.Once
}

// Initializer
func New(size int) *WorkerPool {
	if size <= 0 {
		size = runtime.NumCPU() // default: available CPUs
	}

	return &WorkerPool{
		size:     size,
		jobs:     make(chan WorkerJob),     // unbuffered channel
		errchan:  make(chan error, size*2), // buffered channel
		quit:     make(chan interface{}),
		closeJob: make(chan string), // unbuffered channel
		allerrs:  make([]error, 0),
		wg:       &sync.WaitGroup{},
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.size; i++ {
		wp.wg.Add(1)
		go wp.worker(uuid.NewString()) // Start worker
	}

	// Start error collection
	go wp.collectErrors()
}

func (wp *WorkerPool) Add(job WorkerJob) {
	select {
	case wp.jobs <- job:
	case <-wp.quit:
		// pool shut down
		logger.WARN().Println("Pool shutting down. Cannot add JOB")
		return
	}
}

func (wp *WorkerPool) StopWaitGracefully() []error {
	time.Sleep(30 * time.Millisecond) // graceful wait
	wp.onceDo.Do(func() {
		close(wp.quit)    // send quit signal
		close(wp.jobs)    // shutdown job channel
		wp.wg.Wait()      // wait for all workers to finish
		close(wp.errchan) // close error channel
	})

	// Return the collected errors
	wp.mu.Lock()
	errs := make([]error, len(wp.allerrs))
	copy(errs, wp.allerrs)
	wp.mu.Unlock()

	return errs
}

func (wp *WorkerPool) collectErrors() {
	for err := range wp.errchan {
		wp.mu.Lock()
		wp.allerrs = append(wp.allerrs, err)
		wp.mu.Unlock()
	}
}

func (wp *WorkerPool) worker(id string) {
	defer wp.wg.Done()

	for {
		select {
		case job, ok := <-wp.jobs:
			if !ok {
				// no job
				return
			}
			// execute job
			if err := job(); err != nil {
				// pass to error channel
				select {
				case wp.errchan <- fmt.Errorf("Worker(%s): Failed to execute JOB - %w", id, err):
				default:
					// err channel closed
					logger.WARN().Printf("Worker(%s): Error channel full.\nDropping err - %s", id, err.Error())
					return
				}
			} else {
				logger.SUCCESS().Printf("Worker(%s): JOB executed successfully!", id)
			}
		case <-wp.quit:
			// pool shut down
			logger.WARN().Printf("Worker(%s): Pool shutting down.\nDropping JOB", id)
			return
		}
	}
}
