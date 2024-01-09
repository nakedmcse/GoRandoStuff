package main

import (
	"fmt"
	"math/rand"
	"slices"
	"sort"
	"strconv"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	count := 100_000
	list := generateRandomArray(count)
	list2 := make([]int, count)
	list3 := make([]int, count)
	list4 := make([]int, count)
	copy(list2, list)
	copy(list3, list)
	copy(list4, list)
	wg.Add(3)
	go func() {
		defer wg.Done()
		defer timer("quickSort for " + strconv.Itoa(count) + " items")()
		quickSortStart(list)
	}()

	go func() {
		defer wg.Done()
		defer timer("built Slices in sort for " + strconv.Itoa(count) + " items")()
		slices.Sort(list2)
	}()

	go func() {
		defer wg.Done()
		defer timer("built Sorts in sort for " + strconv.Itoa(count) + " items")()
		sort.Ints(list3)
	}()

	wg.Wait()
	fmt.Printf("arrays all equal %t", slicesEqual(list, list2, list3))
}

func slicesEqual[T comparable](arrays ...[]T) bool {
	if len(arrays) < 1 {
		return false
	}
	length := len(arrays[0])
	for _, array := range arrays {
		if len(array) != length {
			return false
		}
		for i, v := range array {
			if v != arrays[0][i] {
				return false
			}
		}
	}
	return true
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func quickSortStart(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func generateRandomArray(n int) (numbers []int) {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	for i := 0; i < n; i++ {
		numbers = append(numbers, rng.Int())
	}
	return
}

func generateBadRandomArray(n int) (numbers []int) {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	var choice int
	for i := 0; i < n; i++ {
		choice = rng.Int()
		if rng.Intn(10) < 8 {
			choice = 0
		}
		numbers = append(numbers, choice)
	}
	return
}
