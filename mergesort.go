package main

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]

	return merge(MergeSort(left), MergeSort(right))
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))

	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[i] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
