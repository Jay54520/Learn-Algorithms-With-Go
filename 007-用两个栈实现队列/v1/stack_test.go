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

}
