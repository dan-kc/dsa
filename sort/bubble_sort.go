package sort

func BubbleSort(a *[]int) {
	for i := 0; i < len(*a)-1; i++ {
		for j := 0; j < j-i-1; j++ {
			if (*a)[j] > (*a)[j+1] {
				(*a)[j+1], (*a)[j] = (*a)[j], (*a)[j+1]
			}
		}
	}
}
