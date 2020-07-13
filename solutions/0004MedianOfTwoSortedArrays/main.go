package main

import (
	"fmt"
	"math"
)

/**
 * 方法一：
 * 合并数组为一个，然后取按奇数或者偶数计算出中位数的位置 i，直接 A[i] 取就可以。
 * 但这种方式，合并数组，需要重新排序数组，排序需要遍历两个数组，时间复杂度为 O(m+n)，大于 O(log (m+n)) 不符合要求。
 *
 * 方法二：求第 K 小的数
 *
 * 方法三：解数学题～～
 */

/**
 * findMedianSortedArrays: 方法二
 *
 * Runtime: 16ms
 * Memory: 5.6MB
 *
 */
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	k := int((m+n)/2) + 1
	if (m+n)%2 == 0 {
		return (findKth(nums1, nums2, m, n, 0, 0, k-1) + findKth(nums1, nums2, m, n, 0, 0, k)) / 2
	} else {
		return findKth(nums1, nums2, m, n, 0, 0, k)
	}
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
func minInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

/**
 * findKth: 寻找第 k 小的数
 */
func findKth(nums1, nums2 []int, m, n, start1, start2, k int) float64 {
	len1 := m - start1
	len2 := n - start2
	halfK := int(k / 2)
	if len1 == 0 {
		return float64(nums2[start2 + k - 1])
	}
	if len2 == 0 {
		return float64(nums1[start1 + k - 1])
	}
	if k == 1 {
		return math.Min(float64(nums1[start1]), float64(nums2[start2]))
	}
	i := start1 + minInt(len1, halfK) - 1
	j := start2 + minInt(len2, halfK) - 1
	if nums1[i] > nums2[j] {
		return findKth(nums1, nums2, m, n, start1, j+1, k-minInt(len2, halfK))
	} else {
		return findKth(nums1, nums2, m, n, i+1, start2, k-minInt(len1, halfK))
	}
}

/**
 * 方法三
 *
 * Runtime: 12ms
 * Memory: 5.6MB
 *
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number}
 */
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		temp := nums1
		nums1 = nums2
		nums2 = temp
		m = len(nums1)
		n = len(nums2)
	}
	iMin, iMax, halfLen := 0, m, int((m+n+1)/2)
	for iMin <= iMax {
		i := int((iMin + iMax) / 2)
		j := halfLen - i
		if i < iMax && nums2[j-1] > nums1[i] {
			iMin = i + 1
		} else if i > iMin && nums1[i-1] > nums2[j] {
			iMax = i - 1
		} else {
			maxLeft := 0
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = maxInt(nums1[i-1], nums2[j-1])
			}
			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}
			minRight := 0
			if i == m {
				minRight = nums2[j]
			} else if j == n {
				minRight = nums1[i]
			} else {
				minRight = minInt(nums1[i], nums2[j])
			}
			return (float64(maxLeft) + float64(minRight)) / 2.0
		}
	}
	return 0.0
}

func main() {
	fmt.Println(findMedianSortedArrays2([]int{1}, []int{2, 3, 4, 5, 6})) // 3.5
	fmt.Println(findMedianSortedArrays2([]int{2}, []int{1, 3})) // 2.0
	fmt.Println(findMedianSortedArrays2([]int{2}, []int{1, 3, 4})) // 2.5
	fmt.Println(findMedianSortedArrays2([]int{1, 2}, []int{3, 4})) // 2.5
	fmt.Println(findMedianSortedArrays2([]int{}, []int{1, 2})) // 1.5

	fmt.Println(findMedianSortedArrays([]int{1}, []int{2, 3, 4, 5, 6})) // 3.5
	fmt.Println(findMedianSortedArrays([]int{2}, []int{1, 3})) // 2.0
	fmt.Println(findMedianSortedArrays([]int{2}, []int{1, 3, 4})) // 2.5
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4})) // 2.5
	fmt.Println(findMedianSortedArrays([]int{}, []int{1, 2})) // 1.5
}
