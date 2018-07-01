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

	t.Run("|", func(t *testing.T) {
		var int1 int64 = 1
		var int2 int64 = 2
		fmt.Printf("%d's binary string is %s \n", int1, strconv.FormatInt(int1, 2))
		fmt.Printf("%d's binary string is %s \n", int2, strconv.FormatInt(int2, 2))
		fmt.Printf("%s | %s is %s \n", strconv.FormatInt(int1, 2), strconv.FormatInt(int2, 2), strconv.FormatInt(int1 | int2, 2))
	})

	t.Run("&", func(t *testing.T) {
		var int1 int64 = 1
		var int2 int64 = 4
		fmt.Printf("%d's binary string is %s \n", int1, strconv.FormatInt(int1, 2))
		fmt.Printf("%d's binary string is %s \n", int2, strconv.FormatInt(int2, 2))
		fmt.Printf("%s & %s is %s \n", strconv.FormatInt(int1, 2), strconv.FormatInt(int2, 2), strconv.FormatInt(int1 & int2, 2))
	})

	t.Run("^", func(t *testing.T) {
		var int1 int64 = 1
		var int2 int64 = 2
		fmt.Printf("%d's binary string is %s \n", int1, strconv.FormatInt(int1, 2))
		fmt.Printf("%d's binary string is %s \n", int2, strconv.FormatInt(int2, 2))
		fmt.Printf("%d ^ %d is %s \n", int1, int2, strconv.FormatInt(int1 ^ int2, 2))
	})

	t.Run("<< 1", func(t *testing.T) {
		var int1 int64 = 4
		int2 := int1 << 1
		fmt.Printf("%d's binary string is %s \n", int1, strconv.FormatInt(int1, 2))
		fmt.Printf("%s << 1 is %s \n", strconv.FormatInt(int1, 2), strconv.FormatInt(int2, 2))
	})

	t.Run(">> 1", func(t *testing.T) {
		var int1 int64 = 4
		int2 := int1 >> 1
		fmt.Printf("%d's binary string is %s \n", int1, strconv.FormatInt(int1, 2))
		fmt.Printf("%s >> 1 is %s \n", strconv.FormatInt(int1, 2), strconv.FormatInt(int2, 2))
	})

	t.Run("位移越界", func(t *testing.T) {
		flag := 1
		times := 0
		for flag != 0 {
			flag <<= 1
			times += 1
		}
		fmt.Printf("%d << 1 %d times is 0", flag, times)
	})
}