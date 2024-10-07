package main

func main() {
	println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}

func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)

	for _, num := range nums {
		numSet[num] = true
	}
	longest := 0

	for _, n := range nums {
		_, ok := numSet[n-1]

		if !ok {
			length := 0
			for numSet[n+length] == true {
				length++
			}
			longest = max(length, longest)
		}
	}

	return longest
}
