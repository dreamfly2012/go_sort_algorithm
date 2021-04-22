package main

import (
	"fmt"
	"testing"
)

func TestShellSort(t *testing.T) {
	list := []int{2, 6, 8, 1, 9, 0, 3, 4, 7, 5}
	ShellSort(list)
	for i := 0; i < len(list); i++ {
		fmt.Println(list[i])
	}
}

func ShellSort(list []int) {
	length := len(list)
	if length < 2 {
		return
	}
	for step := length / 2; step > 0; step = step / 2 {
		for i := step; i < length; i++ {
			j := i
			for j >= step && list[j] < list[j-step] {
				list[j], list[j-step] = list[j-step], list[j]
				j -= step
			}
		}
	}
}
