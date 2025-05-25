package scheduler

import (
	"sync"
	"testing"
	"time"
)

func TestScheduler(t *testing.T) {
	s := NewScheduler()
	defer s.Stop()

	var wg sync.WaitGroup
	var executed bool

	// Schedule task for 1 second from now
	task := Task{
		ExecuteAt: time.Now().Add(time.Second),
		Data:      nil,
		Handler: func(data any) error {
			executed = true
			wg.Done()
			return nil
		},
	}

	wg.Add(1)
	s.Schedule(task)

	// Wait for task to execute
	wg.Wait()

	if !executed {
		t.Error("Task was not executed")
	}
}

func TestSchedulerMultipleTasks(t *testing.T) {
	s := NewScheduler()
	defer s.Stop()

	var wg sync.WaitGroup
	executionOrder := make([]int, 0)
	var mu sync.Mutex

	// Schedule multiple tasks with different delays
	for i := 0; i < 3; i++ {
		wg.Add(1)
		task := Task{
			ExecuteAt: time.Now().Add(time.Duration(i+1) * time.Second),
			Data:      i,
			Handler: func(data any) error {
				mu.Lock()
				executionOrder = append(executionOrder, data.(int))
				mu.Unlock()
				wg.Done()
				return nil
			},
		}
		s.Schedule(task)
	}

	// Wait for all tasks to execute
	wg.Wait()

	// Verify execution order
	expectedOrder := []int{0, 1, 2}
	for i, v := range executionOrder {
		if v != expectedOrder[i] {
			t.Errorf("Expected task %d to execute at position %d, got %d", expectedOrder[i], i, v)
		}
	}
}
