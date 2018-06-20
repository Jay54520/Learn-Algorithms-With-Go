package main

import "testing"

func TestReplaceSpace(t *testing.T)  {
	got := ReplaceSpace("We Are Happy")
	want := "We%20Are%20Happy"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
