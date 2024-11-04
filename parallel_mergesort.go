package main

import (
	"fmt"
	"sync"
)

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Otherwise, use parallel recursion with goroutines.
	var wg sync.WaitGroup
	var left, right []int

	wg.Add(2)
	go func() {
		defer wg.Done()
		left = MergeSort(arr[:len(arr)/2])
	}()
	go func() {
		defer wg.Done()
		right = MergeSort(arr[len(arr)/2:])
	}()
	wg.Wait()

	return merge(left, right)
}

// merge combines two sorted slices into one sorted slice.
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
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

func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	sortedArr := MergeSort(arr)
	fmt.Println("Sorted array:", sortedArr)
}
