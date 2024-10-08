package main

import "fmt"

func main() {
	res := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Printf("first: %d, second: %d", res[0], res[1])
}

func twoSum(numbers []int, target int) []int {
	i := 0
	j := len(numbers) - 1

	for numbers[i]+numbers[j] != target {
		if numbers[i]+numbers[j] > target {
			j--
		} else {
			i++
		}
	}

	return []int{i + 1, j + 1}
}
