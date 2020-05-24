package maxqueue_test

import (
	"hsecode.com/stdlib/maxqueue/int"
	"testing"
)

func TestErrors(t *testing.T) {
	queue := maxqueue.New()
	_, err := queue.Max()
	if err == nil {
		t.Fatal("Max empty error doesn't work")
	}
	_, err1 := queue.Pop()
	if err1 == nil {
		t.Fatal("Pop empty error doesn't work")
	}
}

func TestPushPop(t *testing.T) {
	queue := maxqueue.New()
	queue.Push(0)
	queue.Push(1)
	queue.Push(2)
	value, _ := queue.Pop()
	if value != 0 {
		t.Fatal("Pop doesn't work 1st in")
	}
	value1, _ := queue.Pop()
	if value1 != 1 {
		t.Fatal("Pop doesn't work 2nd in")
	}
}

func TestMax(t *testing.T) {
	queue := maxqueue.New()
	queue.Push(1)
	queue.Push(3)
	queue.Push(2)
	value, _ := queue.Max()
	if value != 3 {
		t.Fatal("Max in q.in doesn't work")
	}
	p1, _ := queue.Pop()
	p2, _ := queue.Pop()
	value1, _ := queue.Max()
	if value1 != 2 {
		t.Fatalf("Max after pop doesn't work: %d, %d, %d", value1, p1, p2)
	}

	queue.Push(4)
	value2, _ := queue.Max()
	if value2 != 4 {
		t.Fatal("Max after new push doesn't work")
	}

}
