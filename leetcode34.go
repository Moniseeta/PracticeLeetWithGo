package main

import "fmt"

/* 34. Find First and Last Position of Element in Sorted Array
Medium
Topics
Companies

Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.

If target is not found in the array, return [-1, -1].

You must write an algorithm with O(log n) runtime complexity.



Example 1:

Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]

Example 2:

Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]

Example 3:

Input: nums = [], target = 0
Output: [-1,-1]



Constraints:

    0 <= nums.length <= 105
    -109 <= nums[i] <= 109
    nums is a non-decreasing array.
    -109 <= target <= 109

*/

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 || nums[len(nums)-1] < target || nums[0] > target {
		return []int{-1, -1}
	}

	leftIndex := getFirstIndex(nums, target)
	if leftIndex == -1 {
		return []int{-1, -1}
	}
	rightIndex := getLastIndex(nums, leftIndex, target)
	return []int{leftIndex, rightIndex}

}

func getFirstIndex(nums []int, target int) int {
	if nums[0] == target {
		return 0
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target && nums[mid-1] < target {
			return mid
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func getLastIndex(nums []int, leftIndex int, target int) int {
	if nums[len(nums)-1] == target {
		return len(nums) - 1
	}
	left, right := leftIndex, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] == target && nums[mid+1] > target {
			return mid
		} else {
			left = mid + 1
		}
	}
	return right
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	indx := searchRange(nums, target)
	fmt.Printf("%d %d", indx[0], indx[1])
}
