package main

import (
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		nums1  []int
		nums2  []int
		median float64
	}{
		{[]int{1}, []int{2, 3, 4, 5, 6}, 3.5},
		{[]int{2}, []int{1, 3}, 2.0},
		{[]int{2}, []int{1, 3, 4}, 2.5},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{}, []int{1, 2}, 1.5},
	}

	for _, tt := range tests {
		median := findMedianSortedArrays(tt.nums1, tt.nums2)
		if median != tt.median {
			t.Errorf("got %v for input %v, %v; expected %v", median, tt.nums1, tt.nums2, tt.median)
		}
	}
}

func TestFindMedianSortedArrays2(t *testing.T) {
	tests := []struct {
		nums1  []int
		nums2  []int
		median float64
	}{
		{[]int{1}, []int{2, 3, 4, 5, 6}, 3.5},
		{[]int{2}, []int{1, 3}, 2.0},
		{[]int{2}, []int{1, 3, 4}, 2.5},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{}, []int{1, 2}, 1.5},
	}

	for _, tt := range tests {
		median := findMedianSortedArrays2(tt.nums1, tt.nums2)
		if median != tt.median {
			t.Errorf("got %v for input %v, %v; expected %v", median, tt.nums1, tt.nums2, tt.median)
		}
	}
}
