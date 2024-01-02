package leetcode

func TopKFrequent(nums []int, k int) (res []int) {
	countMap := map[int]int{}
	countSlice := make([][]int, len(nums))

	for _, num := range nums {
		countMap[num]++
	}

	for key, val := range countMap {
		countSlice[val] = append(countSlice[val], key)
	}

	for i := len(countSlice) - 1; i >= 0; i-- {
		res = append(res, countSlice[i]...)
		if len(res) == k {
			return
		}
	}

	return
}
