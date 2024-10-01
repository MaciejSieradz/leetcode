package main

func main() {

}

func twoSum(nums []int, target int) []int {
	numbers := make(map[int]int)

	for index, element := range nums {
		a, b := numbers[target-element]
		if b == true {
			return []int{index, a}
		} else {
			numbers[element] = index
		}
	}

	return []int{}
}
