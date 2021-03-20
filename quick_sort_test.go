package main

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {

	//arr := []int{2, 3, 5, 7, 9, 0, 1, 4, 8, 6}
	arr := []int{2, 3, 5, 8, 9, 6, 1, 4, 8, 6}
	QuickSort(arr)
	for i := 0; i < 10; i++ {
		fmt.Println(arr[i])
	}
	// t.Fatal("not implemented")
}

func QuickSort(list []int) {
	if len(list) <= 1 {
		return
	}

	mid, i := list[0], 1
	head, tail := 0, len(list)-1

	for head < tail {
		if list[i] > mid {
			list[i], list[tail] = list[tail], list[i]
			tail--
		} else {
			list[i], list[head] = list[head], list[i]
			i++
			head++
		}

	}
	list[head] = mid
	QuickSort(list[:head])
	QuickSort(list[head+1:])
}
