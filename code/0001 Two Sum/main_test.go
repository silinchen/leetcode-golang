package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		ret    []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
	}

	for _, tt := range tests {
		ret := twoSum(tt.nums, tt.target)
		if ok := reflect.DeepEqual(ret, tt.ret); !ok {
			t.Errorf("got %v for input %v; expected %v", ret, tt.nums, tt.ret)
		}
	}
}
