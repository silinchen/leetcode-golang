package main

import "fmt"

/**
 * Runtime: 0~4ms
 * Memory: 3.4MB
 */
func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int, len(nums))
	for i, num := range nums {
		complement := target - num
		if _, ok := numsMap[complement]; ok {
			return []int{numsMap[complement], i}
		}
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

