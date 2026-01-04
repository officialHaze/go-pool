# go-pool

go-pool is a lightweight, reusable **worker pool library for Go** that simplifies concurrent task execution, goroutine management, and graceful shutdown across projects.

It is designed to standardize concurrency patterns so you don’t need to re-implement worker pools in every application.

---

## Features

- Reusable worker pool abstraction
- Controlled goroutine concurrency
- Centralized pool initialization
- Buffered job queue
- Graceful shutdown support
- Minimal API with zero external dependencies

---

## Installation

```bash
go get github.com/officialHaze/go-pool
```

---

## Quick Start

```bash
package main

import (
	"log"

	"github.com/officialHaze/go-pool"
)

func main() {
	pool := workerpool.New(10) // 10 - Pool size (number of CPUs to use)

	pool.Start()

	pool.Add(func() error {
		log.Println("processing job")
        // Define your JOB here
		return nil
	})

    if errs := pool.StopWaitGracefully(); len(errs) > 0 {
        // Do something with the list of errors
    }
}
```

---

## Lifecycle

- Initialize the pool using New

- Start the worker pool

- Submit tasks/jobs to the pool

- Workers execute tasks concurrently

- Shut down gracefully, allowing in-flight tasks to complete & catch errors that occurred during job execution

---

## Motivation

In many Go projects, concurrency logic often becomes:

- Repetitive and copy-pasted
- Tightly coupled with business logic
- Hard to shut down cleanly
- Difficult to reason about under load

go-pool solves this by providing a **consistent and predictable worker pool** that can be shared across services, APIs, and background workers.

---

## Contributing

Contributions, issues, and feature requests are welcome.

Feel free to open an issue or submit a pull request.

---

## License

This project is licensed under the MIT License.
