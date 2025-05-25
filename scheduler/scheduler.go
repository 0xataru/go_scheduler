package scheduler

import (
	"sync"
	"time"

	"github.com/0xataru/go_scheduler/async_queue"
)

type Task struct {
	ExecuteAt time.Time
	Data      any
	Handler   func(data any) error
}

type Scheduler struct {
	tasks async_queue.Queue[Task]
	mu    sync.Mutex
	stop  chan struct{}
}

func NewScheduler() *Scheduler {
	s := &Scheduler{
		tasks: async_queue.NewQueue[Task](),
		stop:  make(chan struct{}),
	}

	go s.run()
	return s
}

func (s *Scheduler) Schedule(task Task) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.tasks.Put(task)
}

func (s *Scheduler) run() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-s.stop:
			return
		case <-ticker.C:
			s.processTasks()
		}
	}
}

func (s *Scheduler) processTasks() {
	s.mu.Lock()
	tasks := s.tasks.Flush()
	s.mu.Unlock()

	now := time.Now().UTC()
	for _, task := range tasks {
		if task.ExecuteAt.Before(now) || task.ExecuteAt.Equal(now) {
			go task.Handler(task.Data)
		} else {
			s.Schedule(task)
		}
	}
}

func (s *Scheduler) Stop() {
	close(s.stop)
}

func (s *Scheduler) CancelTask(taskID string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Remove task from queue if it exists
	s.tasks.Remove(func(task Task) bool {
		if data, ok := task.Data.(map[string]any); ok {
			if id, ok := data["task_id"].(string); ok {
				return id == taskID
			}
		}
		return false
	})
}
