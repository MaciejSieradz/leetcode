package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Print(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}

func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	freq := make([][]int, len(nums)+1)

	for _, n := range nums {
		count[n] = 1 + count[n]
	}

	for key, value := range count {
		fmt.Printf("key: %d, value: %d\n", key, value)
	}

	for n, c := range count {
		freq[c] = append(freq[c], n)
	}

	var res []int

	for i := len(freq) - 1; i > 0; i-- {
		for _, n := range freq[i] {
			res = append(res, n)
			if len(res) == k {
				return res
			}
		}
	}

	return []int{}
}

func topKFrequentSlow(nums []int, k int) []int {
	numbers := make(map[int]int)

	for _, number := range nums {
		num, ok := numbers[number]

		if !ok {
			numbers[number] = 1
		} else {
			numbers[number] = num + 1
		}
	}

	keys := make([]int, 0, len(numbers))

	for key := range numbers {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return numbers[keys[i]] > numbers[keys[j]]
	})

	return keys[:k]
}
