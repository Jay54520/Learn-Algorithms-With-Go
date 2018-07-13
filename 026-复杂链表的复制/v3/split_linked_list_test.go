package main

import "testing"

func TestSplitLinkedList(t *testing.T)  {
	t.Run("four nodes", func(t *testing.T) {
		fourthNode := &Node{
			4,
			nil,
			nil,
		}
		thirdNode := &Node{
			3,
			fourthNode,
			nil,
		}
		secondNode := &Node{
			2,
			thirdNode,
			nil,
		}
		linkedList := &Node{
			1,
			secondNode,
			nil,
		}

		gotLinkedList1, gotLinkedList2 := splitLinkedList(linkedList)

		wantSecond1 := &Node{
			3,
			nil,
			nil,
		}
		wantLinkedList1 := &Node{
			1,
			wantSecond1,
			nil,
		}

		wantSecond2 := &Node{
			4,
			nil,
			nil,
		}
		wantLinkedList2 := &Node{
			2,
			wantSecond2,
			nil,
		}

		if !linkedListEqual(gotLinkedList1, wantLinkedList1) {
			t.Errorf("got %v want %v", gotLinkedList1, wantLinkedList1)
		}

		if !linkedListEqual(gotLinkedList2, wantLinkedList2) {
			t.Errorf("got %v want %v", gotLinkedList2, wantLinkedList2)
		}

	})

	t.Run("empty", func(t *testing.T) {
		var linkedList *Node
		gotLinkedList1, gotLinkedList2 := splitLinkedList(linkedList)
		if gotLinkedList1 != nil {
			t.Errorf("%v should be nil", gotLinkedList1)
		}
		if gotLinkedList2 != nil {
			t.Errorf("%v should be nil", gotLinkedList2)
		}
	})
}

