package main

import (
	"testing"
	"reflect"
)

func TestStackPush(t *testing.T) {

	t.Run("push", func(t *testing.T) {
		var stack Stack
		stack.Push(1)
		stack.Push(2)
		got := stack.Data
		want := []int{1, 2}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("pop", func(t *testing.T) {
		stack := Stack{
			[]int{1, 2},
		}
		var got []int
		got = append(got, stack.Pop())
		got = append(got, stack.Pop())
		want := []int{2, 1}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
