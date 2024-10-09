package main

func main() {
	println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	res := 0
	j := 0
	k := len(height) - 1

	for j < k {
		y := min(height[j], height[k])
		surface := y * (k - j)
		if surface > res {
			res = surface
		}

		if height[j] <= height[k] {
			j++
		} else {
			k--
		}

	}

	return res
}
