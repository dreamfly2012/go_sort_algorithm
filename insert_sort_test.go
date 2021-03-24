package main

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	list := []int{3, 4, 1, 8, 5, 2, 9, 7, 6, 0}
	InsertSort(list)
	for i := 0; i < len(list); i++ {
		fmt.Println(list[i])
	}
}

func InsertSort(list []int) {
	for i := 0; i < len(list); i++ {
		for j := 0; j < i; j++ {
			if list[i] < list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}

	}
	fmt.Println("insert sort")
}
