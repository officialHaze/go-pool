package workerpool

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"
)

type WorkerJob func() error

type WorkerPool struct {
	size    int
	jobs    chan WorkerJob
	errchan chan error
	quit    chan interface{}
	allerrs []error
	mu      sync.Mutex
	wg      *sync.WaitGroup
	onceDo  sync.Once
}

// Initializer
func New(size int) *WorkerPool {
	if size <= 0 {
		size = runtime.NumCPU() // default: available CPUs
	}

	return &WorkerPool{
		size:    size,
		jobs:    make(chan WorkerJob),     // unbuffered channel
		errchan: make(chan error, size*2), // buffered channel
		quit:    make(chan interface{}),
		allerrs: make([]error, 0),
		wg:      &sync.WaitGroup{},
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.size; i++ {
		wp.wg.Add(1)
		go wp.worker(i + 1) // Start worker
	}

	// Start error collection
	go wp.collectErrors()
}

func (wp *WorkerPool) Add(job WorkerJob) {
	select {
	case wp.jobs <- job:
	case <-wp.quit:
		// pool shut down
		log.Println("⚠️ Pool shutting down. Cannot add JOB")
		return
	}
}

func (wp *WorkerPool) StopWaitGracefully() []error {
	wp.onceDo.Do(func() {
		close(wp.quit)                    // send quit signal
		time.Sleep(30 * time.Millisecond) // graceful wait
		close(wp.jobs)                    // shutdown job channel
		wp.wg.Wait()                      // wait for all workers to finish
		close(wp.errchan)                 // close error channel
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

func (wp *WorkerPool) worker(id int) {
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
				case wp.errchan <- fmt.Errorf("⚠️ Worker(%d): Failed to execute JOB - %w", id, err):
				default:
					// err channel closed
					log.Printf("⚠️ Worker(%d): Error channel full. Dropping err - %v", id, err)
					return
				}
			} else {
				log.Printf("✅ Worker(%d): JOB executed successfully!", id)
			}
		case <-wp.quit:
			// pool shut down
			log.Printf("⚠️ Worker(%d): Pool shutting down. Dropping JOB", id)
			return
		}
	}
}
