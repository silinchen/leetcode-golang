package main

import "fmt"

// https://leetcode.com/problems/two-sum/

// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// Example:

//   Given nums = [2, 7, 11, 15], target = 9,
//   Because nums[0] + nums[1] = 2 + 7 = 9,
//   return [0, 1].\

/**
 * 在进行迭代并将元素插入到表中的同时，我们还会回过头来检查表中是否已经存在当前元素所对应的目标元素。
 * 如果它存在，那我们已经找到了对应解，并立即将其返回。
 */
func twoSum(nums []int, target int) []int {
	// map 表，缓存已经遍历过的唯一 item，跟对应的 index，可以使用普通对象替代
	numsMap := make(map[int]int, len(nums))
	for i, num := range nums {
		complement := target - num
		// 取目标值与当前项的差值，如果差值已经出现过，则取差值的索引跟当前值的索引作为结果
		if _, ok := numsMap[complement]; ok {
			return []int{numsMap[complement], i}
		}
		// 判断 map 是否有缓存当前的项，如果没有，缓存它
		if _, ok := numsMap[num]; !ok {
			numsMap[num] = i
		}
	}
	return []int{}
}
func main() {
	nums := []int{2, 7, 11, 15}
	ret := twoSum(nums, 9)
	fmt.Println(ret)
}

/**
 * Runtime: 0~4ms
 * Memory: 3.4MB
 */
