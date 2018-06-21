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
		"start",
		&secondNode,
	}

	got := ListLinkedListValue(&firstNode)
	want := []interface{}{3, 2, "start"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
