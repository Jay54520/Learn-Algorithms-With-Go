package main

import "testing"

func TestSearch(t *testing.T)  {

	array := [][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15},
	}

	t.Run("查找存在的数字", func(t *testing.T) {
		got := Search(array, 5)
		want := true

		if want != got {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
