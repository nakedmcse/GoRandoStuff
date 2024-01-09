package main

import (
	"math/rand"
	"slices"
	"time"
)

func main() {
	count := 10
	list := generateRandomArray(count)
	var list2 [10]int
	list3 := make([]int, count)
	for c := 0; c < len(list); c++ {
		list2[c] = list[c]
	}
	copy(list3, list)
	otherlist := generateRandomArray(count)
	dummy(list3)
	bogoSort(otherlist)
	//list3[1] = 999
}

func dummy(arr []int) []int {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	temp := make([]int, len(arr))
	copy(temp, arr)
	for k := 1; k < 49000; k++ {
		for i, v := range rng.Perm(len(arr)) {
			temp[i] = arr[v]
		}
	}
	return temp
}
func bogoSort(arr []int) []int {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	temp := make([]int, len(arr))
	copy(temp, arr)
	for !slices.IsSorted(temp) {
		for i, v := range rng.Perm(len(arr)) {
			temp[i] = arr[v]
		}
	}
	return temp
}

func generateRandomArray(n int) (numbers []int) {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	for i := 0; i < n; i++ {
		numbers = append(numbers, rng.Int())
	}
	return
}
