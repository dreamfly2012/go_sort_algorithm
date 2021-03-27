package main

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	list := []int{2, 3, 8, 1, 0, 4, 9, 5, 7, 6}
	HeapSort(list, len(list))
	for i := 0; i < len(list); i++ {
		fmt.Println(list[i])
	}
}

func HeapSort(list []int, length int) {
	buildMaxHeap(list, length)
	for i := length - 1; i >= 0; i-- {
		list[0], list[i] = list[i], list[0]
		list = list[:1]
		length = length - 1
		ajustHeap(list, 0, length)
	}
}

func buildMaxHeap(list []int, length int) {
	for i := length / 2; i >= 0; i-- {
		ajustHeap(list, i, length)
	}
}

func ajustHeap(list []int, pos, length int) {
	left := 2*pos + 1
	right := 2*pos + 2
	largest := pos
	if left < length && list[left] < list[largest] {
		largest = left
	}
	if right < length && list[right] < list[largest] {
		largest = right
	}
	if largest != pos {
		list[pos], list[largest] = list[largest], list[pos]
		ajustHeap(list, largest, length)
	}
}
