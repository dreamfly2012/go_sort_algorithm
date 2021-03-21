package main

import (
	"fmt"
	"testing"
)

func TestSelectSort(t *testing.T) {
	list := []int{2, 5, 6, 8, 9, 1, 4, 3, 7, 0}
	SelectSort(list)
	for i := 0; i < len(list); i++ {
		fmt.Println(list[i])
	}
}

func SelectSort(list []int) {
	for i := 0; i < len(list); i++ {
		min := i
		for j := i + 1; j < len(list); j++ {
			if list[j] < list[min] {
				min = j
			}
		}
		list[i], list[min] = list[min], list[i]
	}

}
