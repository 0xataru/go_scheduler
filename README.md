# Go Scheduler

[![Go Report Card](https://goreportcard.com/badge/github.com/0xataru/go_scheduler)](https://goreportcard.com/report/github.com/0xataru/go_scheduler)
[![GoDoc](https://godoc.org/github.com/0xataru/go_scheduler?status.svg)](https://godoc.org/github.com/0xataru/go_scheduler)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8?logo=go)](https://golang.org)
[![Build Status](https://img.shields.io/badge/build-passing-brightgreen)](https://github.com/0xataru/go_scheduler/actions)

A lightweight and efficient task scheduler for Go applications that allows scheduling tasks for future execution.

## Features

- Schedule tasks for future execution
- Cancel scheduled tasks
- Asynchronous task processing
- Thread-safe operations
- Generic queue implementation for task management

## Installation

```bash
go get github.com/0xataru/go_scheduler
```

## Usage

### Basic Example

```go
package main

import (
    "fmt"
    "time"
    
    "github.com/0xataru/go_scheduler/scheduler"
)

func main() {
    // Create a new scheduler
    s := scheduler.NewScheduler()
    defer s.Stop()

    // Create a task
    task := scheduler.Task{
        ExecuteAt: time.Now().Add(5 * time.Second),
        Data: map[string]any{
            "task_id": "task-1",
            "message": "Hello, World!",
        },
        Handler: func(data any) error {
            if m, ok := data.(map[string]any); ok {
                fmt.Println(m["message"])
            }
            return nil
        },
    }

    // Schedule the task
    s.Schedule(task)

    // Wait for task execution
    time.Sleep(6 * time.Second)
}
```

### Canceling Tasks

```go
// Schedule a task
taskID := "task-1"
task := scheduler.Task{
    ExecuteAt: time.Now().Add(5 * time.Second),
    Data: map[string]any{
        "task_id": taskID,
    },
    Handler: func(data any) error {
        fmt.Println("Task executed")
        return nil
    },
}
s.Schedule(task)

// Cancel the task before execution
s.CancelTask(taskID)
```

## Components

### Scheduler

The main component that manages task scheduling and execution. It provides methods for:
- Scheduling new tasks
- Canceling scheduled tasks
- Graceful shutdown

### Async Queue

A generic queue implementation that provides thread-safe operations for task management.

## Requirements

- Go 1.24 or higher

## License

This project is dual-licensed under both the MIT License and Apache License 2.0.
