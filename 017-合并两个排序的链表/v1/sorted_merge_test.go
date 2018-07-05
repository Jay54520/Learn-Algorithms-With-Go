package main

import (
	"testing"
	"reflect"
)

func TestSortedMerge(t *testing.T)  {

	t.Run("two three nodes linked lists", func(t *testing.T) {
		l1FourthNode := &Node{20, nil}
		l1ThirdNode := &Node{15, l1FourthNode}
		l1SecondNode := &Node{10, l1ThirdNode}
		l1 := &Node{5, l1SecondNode}

		l2ThirdNode := &Node{20, nil}
		l2SecondNode := &Node{3, l2ThirdNode}
		l2 := &Node{2, l2SecondNode}

		gotNode := SortedMerge(l1, l2)
		var got []int
		for gotNode != nil {
			got = append(got, gotNode.value)
			gotNode = gotNode.next
		}
		want := []int{2, 3, 5, 10, 15, 20, 20}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("two empty linked lists", func(t *testing.T) {
		got := SortedMerge(nil, nil)

		if got != nil {
			t.Errorf("%v should be nil", got)
		}
	})
}
