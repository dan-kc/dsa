package sort

func QuickSort(arr *[]int) {
	sort(arr, 0, len(*arr)-1)
}

func sort(arr *[]int, low, high int) {
	if low < high {
		p := partition(arr, low, high)
		sort(arr, low, p-1)
		sort(arr, p+1, high)
	}
}

func partition(arr *[]int, low, high int) int {
	p := (*arr)[low]
	i := low
	j := high

	for i < j {
		for (*arr)[i] <= p {
			i++
		}
		for (*arr)[j] > p {
			j--
		}
		if i < j {
			(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
		}
	}

	(*arr)[low], (*arr)[j] = (*arr)[j], (*arr)[low]
	return j
}
