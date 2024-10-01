package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
}

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)

	for _, str := range strs {
		sortedStr := sortString(str)
		anagrams[sortedStr] = append(anagrams[sortedStr], str)
	}

	var answer [][]string

	for _, value := range anagrams {

		answer = append(answer, value)
	}

	return answer
}

func sortString(s string) string {
	b := []rune(s)

	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})

	return string(b)
}
