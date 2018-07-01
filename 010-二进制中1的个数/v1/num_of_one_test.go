package main

import "testing"

func TestNumberOfOne(t *testing.T)  {
	got := NumberOfOne(2)
	want := 1

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
