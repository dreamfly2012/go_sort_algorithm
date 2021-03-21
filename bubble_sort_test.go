package main

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	list := []int{2, 4, 1, 9, 5, 6, 3, 8, 7, 0}
	BubbleSort(list)
	for i := 0; i < len(list); i++ {
		fmt.Println(list[i])
	}

}

func BubbleSort(list []int) {
	for i := 0; i < len(list); i++ {
		for j := i; j < len(list); j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}
