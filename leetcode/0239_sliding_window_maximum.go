package main

import "fmt"

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Print(solve(nums, k))
}

func solve(nums []int, k int) []int {
	output := []int{} // Resulting array to store maximum values
	q := make([]int, 0) // Deque to store indices of elements in the current window
	l, r := 0, 0 // Left and right pointers for the sliding window

	for r < len(nums) {
		// Remove elements from the back of the deque if they are smaller than the current element
		for len(q) != 0 && nums[q[len(q)-1]] < nums[r] {
			q = q[:len(q)-1]
		}
		q = append(q, r) // Add the current index to the deque

		// Remove the leftmost element from the deque if it's no longer in the current window
		if l > q[0] {
			q = q[1:]
		}

		// If the window size is equal to k, record the maximum element in the window
		if (r + 1) >= k {
			output = append(output, nums[q[0]])
			l++
		}
		r++
	}
	return output
}
