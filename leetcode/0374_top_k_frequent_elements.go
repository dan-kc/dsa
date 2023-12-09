package leetcode

func TopKFrequent(nums []int, k int) (res []int) {
	countMap := map[int]int{}
	countSlice := make([][]int, len(nums)+1)

	for _, num := range nums {
		countMap[num]++
	}

	for num, count := range countMap {
		countSlice[count] = append(countSlice[count], num)
	}

	for i := len(countSlice) - 1; i > 0; i-- {
		res = append(res, countSlice[i]...)
		if len(res) == k {
			return
		}
	}

	return
}
