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
	for step := length / 2; step > 0; step = step / 2 {
		for i := 0; i < step; i++ {
			for j := i + step; j < length && list[j] < list[j-step]; j += step {
				list[j], list[j-step] = list[j-step], list[j]
			}
		}
		fmt.Println(list)
	}
}
