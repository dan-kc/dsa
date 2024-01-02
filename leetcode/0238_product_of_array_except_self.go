package main

import "fmt"

// Input: nums = [1,2,3,4]
// Output: [24,12,8,6]
func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(solve(nums))
}

func solve(nums []int) []int {
	// Pre = [1,1,2,6]
	pre := make([]int, 4)
	pre[0] = 1
	for i := 1; i < len(nums); i++ {
		pre[i] = nums[i-1] * pre[i-1]
	}

	// Post = [24,12,4,1]
	post := make([]int, 4)
	post[len(post)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		post[i] = nums[i+1] * post[i+1]
	}

	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = pre[i] * post[i]
	}

	return res
}
