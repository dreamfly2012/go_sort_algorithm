package main

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	list := []int{2, 5, 1, 9, 7, 3, 4, 8, 0, 6}
	MergeSort(list)
	for i := 0; i < len(list); i++ {
		fmt.Println(list[i])
	}
	fmt.Println("test merge sort")
}

func MergeSort(list []int) {
	if len(list) < 2 {
		return
	}
	//分而治之
	middle := len(list) / 2
	left := list[0:middle]
	right := list[middle:]
	merge_sort(left, right)

}

func merge_sort(lista []int, listb []int) {
	var result []int
	for len(lista) != 0 && len(listb) != 0 {
		if lista[0] < listb[0] {
			result = append(result, lista[0])
			lista = lista[1:]
		} else {
			result = append(result, listb[0])
			listb = listb[1:]
		}
	}
	if len(lista) != 0 {
		result = append(result, lista[0])
		lista = lista[1:]

	}
	if len(listb) != 0 {
		result = append(result, listb[0])
		listb = listb[1:]
	}
}
