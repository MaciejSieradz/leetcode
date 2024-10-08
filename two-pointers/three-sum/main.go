package main

import "sort"

func main() {
	res := threeSum([]int{-1, 0, 1, 2, -1, -4})

	for _, val := range res {
		for _, second := range val {
			println(second)
		}
	}
}

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	var result [][]int

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		k := len(nums) - 1
		j := i + 1

		for j < k {
			sum := nums[i] + nums[j] + nums[k]

			if sum == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				for j < len(nums) && nums[j] == nums[j-1] {
					j++
				}
			}

			if sum < 0 {
				j++
			}

			if sum > 0 {
				k--
			}
		}
	}

	return result
}
