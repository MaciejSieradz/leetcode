package main

import "sort"

func main() {

}

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	sr := []rune(s)

	sort.Slice(sr, func(i, j int) bool {
		return sr[i] < sr[j]
	})

	tr := []rune(t)
	sort.Slice(tr, func(i, j int) bool {
		return tr[i] < tr[j]
	})

	for i := 0; i < len(sr); i++ {
		if sr[i] != tr[i] {
			return false
		}
	}

	return true
}
