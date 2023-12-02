package sort

func sort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = sort(arr, low, p-1)
		arr = sort(arr, p+1, high)
	}
	return arr
}

func QuickSort(arr []int) []int {
	return sort(arr, 0, len(arr)-1)
}

func partition(arr []int, low, high int) ([]int, int) {
	p := arr[low]
	i := low  // Search for elemets bigger than pivot
	j := high // Search for elements smaller than pivot
	for i < j {
		// Increment i until you find a bigger element
		for arr[i] <= p {
			i++
		}
		// Increment j until you find a smaller element
		for arr[j] > p {
			j--
		}
		// Swap them
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[low], arr[j] = arr[j], arr[low]
	// j is the new position of pivot
	return arr, j
}
