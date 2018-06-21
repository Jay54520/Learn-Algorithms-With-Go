package main

import (
	"testing"
	"reflect"
)

func TestListLinkedListValue(t *testing.T)  {
	lastNode := Node{
		3,
		nil,
	}
	secondNode := Node{
		2,
		&lastNode,
	}
	firstNode := Node{
		1,
		&secondNode,
	}

	got := ListLinkedListValue(&firstNode)
	want := []int{3, 2, 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
