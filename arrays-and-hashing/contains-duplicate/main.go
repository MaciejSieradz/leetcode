package main

func main() {
	nums := []int{1, 2, 3, 1}
	containsDuplicate(nums)
}

func containsDuplicate(nums []int) bool {
	numbers := make(map[int]int)

	for _, number := range nums {
		if numbers[number] != 0 {
			return true
		} else {
			numbers[number] = 1
		}
	}

	return false
}
