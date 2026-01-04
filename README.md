# go-pool

go-pool is a lightweight, reusable **worker pool library for Go** that simplifies concurrent task execution, goroutine management, and graceful shutdown across projects.

It is designed to standardize concurrency patterns so you don’t need to re-implement worker pools in every application.

---

## Features

- Reusable worker pool abstraction
- Controlled goroutine concurrency
- Centralized pool initialization
- Buffered job queue
- Context-aware task execution
- Graceful shutdown support
- Minimal API with zero external dependencies

---

## Motivation

In many Go projects, concurrency logic often becomes:

- Repetitive and copy-pasted
- Tightly coupled with business logic
- Hard to shut down cleanly
- Difficult to reason about under load

go-pool solves this by providing a **consistent and predictable worker pool** that can be shared across services, APIs, and background workers.

---

## Installation

```bash
go get github.com/officialHaze/go-pool
```
