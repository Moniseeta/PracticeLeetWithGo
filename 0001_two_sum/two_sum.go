package twosum

/*
1. Two Sum

Source: https://leetcode.com/problems/two-sum/

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].

*/

func twoSum(nums []int, target int) []int {
	mapNumVsIndex := make(map[int]int)
	for i, num := range nums {
		if num < target {
			pair := target - num
			if j, ok := mapNumVsIndex[pair]; ok {
				return []int{j, i}
			} else {
				mapNumVsIndex[num] = i
			}
		}
	}
	return []int{}
}
