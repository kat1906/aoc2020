package main

import (
	"testing"
)

var (
	inputSliceInts1 = []int{1721, 979, 366, 299, 675, 1456}
)
func TestPartOne(t *testing.T) {
	var test = struct{
		in   []int
		want int
	}{inputSliceInts1, 514579}
	output, err := partOne(test.in)
	if err != nil {
		t.Errorf("%v", err)
	}
	if output != test.want {
		t.Errorf("partOne(%v) = %v, want %v", test.in, output, test.want)
	}
}

func TestPartTwo(t *testing.T) {
	var test = struct{
		in   []int
		want int
	}{inputSliceInts1, 241861950}
	output, err := partTwo(test.in)
	if err != nil {
		t.Errorf("%v", err)
	}
	if output != test.want {
		t.Errorf("partTwo(%v) = %v, want %v", test.in, output, test.want)
	}
}
