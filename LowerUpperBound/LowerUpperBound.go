package main

import "fmt"

/*
LowerBound 그리고 UpperBound 는 이분탐색에서 파생 알고리즘이다.

LowerBound : 정렬된 배열에서 target 보다 같거나 큰 값이 처음 나오는 index를 찾는 알고리즘
UpperBound : 정렬된 배열에서 target 보다 처음으로 큰 값이 나오는 index를 찾는 알고리즘
*/

func main() {
	// Example array
	var array []int = []int{10, 20, 30, 30, 30, 40, 50, 60, 70}

	// LowerBound result, output : 2
	fmt.Println("LowerBound :", LowerBound(array, 30))
	// UpperBound result, output : 5
	fmt.Println("UpperBound :", UpperBound(array, 30))
}

// LowerBound function
func LowerBound(array []int, target int) int {
	low, high, mid := 0, len(array)-1, 0
	for low <= high {
		mid = (low + high) / 2
		if array[mid] >= target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

// UpperBound function
func UpperBound(array []int, target int) int {
	low, high, mid := 0, len(array)-1, 0
	for low <= high {
		mid = (low + high) / 2
		if array[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}
