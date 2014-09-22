package sorting

func MergeSort(arr []int) []int {
	return sort(arr, make([]int, len(arr)), 0, len(arr)-1)
}

func sort(arr []int, aux []int, lo int, hi int) []int {
	if lo >= hi {
		return arr
	}

	mid := lo + (hi-lo)/2
	arr = sort(arr, aux, lo, mid)
	arr = sort(arr, aux, mid+1, hi)
	return merge(arr, aux, lo, mid, hi)
}

func merge(arr []int, aux []int, lo int, mid int, hi int) []int {
	copy(aux, arr)
	i, j := lo, mid+1
	for k := lo; k <= hi; k++ {
		if i > mid {
			arr[k] = aux[j]
			j++
		} else if j > hi {
			arr[k] = aux[i]
			i++
		} else if aux[i] < aux[j] {
			arr[k] = aux[i]
			i++
		} else {
			arr[k] = aux[j]
			j++
		}
	}

	return arr
}
