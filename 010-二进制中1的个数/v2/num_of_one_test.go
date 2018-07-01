package main

import (
	"testing"
	"fmt"
	"strconv"
)

func TestNumberOfOne(t *testing.T)  {
	got := NumberOfOne(2)
	want := 1

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestStrConv(t *testing.T)  {
	t.Run("int to binary string", func(t *testing.T) {
		var integer int64 = 3
		binaryString := strconv.FormatInt(integer, 2)
		fmt.Printf("%d's binary string is %s \n", integer, binaryString)
	})

	t.Run("binary string to int", func(t *testing.T) {
		binaryString := "0"
		integer, _ := strconv.ParseInt(binaryString, 2, 64)
		fmt.Printf("binary string %s is integer %d \n", binaryString, integer)
	})
}