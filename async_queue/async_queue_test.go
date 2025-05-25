package async_queue

import (
	"slices"
	"testing"
)

func TestAsyncQueue(t *testing.T) {

	t.Run("test-put-flush", func(t *testing.T) {
		q := NewQueue[int]()

		q.Put(1)
		q.Put(2)
		q.Put(3)
		expected := []int{1, 2, 3}

		if slices.Compare(q.q, expected) != 0 {
			t.Fatalf("Put: %v, want: %v", q.q, expected)
		}

		res := q.Flush()

		if slices.Compare(res, expected) != 0 {
			t.Fatalf("Flush: %v, want: %v", q.q, expected)
		}

		if len(q.q) != 0 {
			t.Fatalf("q not empty after flush")
		}

		q.Put(1)
		q.Put(2)
		q.Put(3)

		if slices.Compare(q.q, expected) != 0 {
			t.Fatalf("refill Put: %v, want: %v", q.q, expected)
		}

		res2 := q.Flush()

		if slices.Compare(res2, expected) != 0 {
			t.Fatalf("refill Flush: %v, want: %v", q.q, expected)
		}

	})
}
