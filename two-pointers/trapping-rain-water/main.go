package main

func main() {
	println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

func trap(height []int) int {

	l, r := 0, len(height)-1
	leftMax, rightMax := height[l], height[r]
	res := 0

	for l < r {
		if leftMax < rightMax {
			l++
			leftMax = max(leftMax, height[l])
			res += leftMax - height[l]
		} else {
			r--
			rightMax = max(rightMax, height[r])
			res += rightMax - height[r]
		}
	}

	return res
}
