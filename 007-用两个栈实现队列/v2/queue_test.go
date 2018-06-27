package main

import (
	"testing"
	"reflect"
)


func TestQueue(t *testing.T) {

	t.Run("enqueue", func(t *testing.T) {
		var queue Queue
		queue.enqueue(1)
		queue.enqueue(2)

		got := queue.Data
		want := []int{1, 2}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("enqueue", func(t *testing.T) {
		queue := Queue{[]int{1, 2}}

		got := []int{
			queue.dequeue(),
			queue.dequeue(),
		}
		want := []int{1, 2}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
