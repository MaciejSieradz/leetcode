package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	res[0] = 1

	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	multiplicator := nums[len(nums)-1]

	for i := len(nums) - 2; i >= 0; i-- {
		res[i] *= multiplicator
		multiplicator *= nums[i]
	}

	return res
}
