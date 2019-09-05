package main

import (
	"fmt"
	"math"
)

var datas = []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}

var goal = 12

type pointerSlice []int

var sliceA pointerSlice

var LoPointer, HiPointer, MidPointer int

func main() {

	sliceA = datas
	max_loop := int(math.Ceil(math.Log2(float64(len(sliceA)))))

	fmt.Printf("\nSlice : %v\n", datas)
	fmt.Printf("Search Value : %d\n", goal)
	fmt.Printf("Search MAX Loop is :%d\n\n", max_loop)

	LoPointer = 0
	HiPointer = len(sliceA) - 1

	fmt.Println("[LoIndex]-[MidIndex]-[HiIndex]")

	binarySearch()

}

func binarySearch() {

	if HiPointer-LoPointer < 0 {
		fmt.Println("")
		fmt.Println("Got Nothing.")
		return
	}

	MidPointer = (HiPointer-LoPointer)/2 + LoPointer
	fmt.Printf("% 5d % 9d % 9d\n", LoPointer, MidPointer, HiPointer)
	if sliceA[MidPointer] > goal {
		HiPointer = MidPointer - 1
	} else if sliceA[MidPointer] < goal {
		LoPointer = MidPointer + 1
	} else {
		fmt.Println("")
		fmt.Println("Value :", sliceA[MidPointer], "at index:", MidPointer)
		return
	}
	binarySearch()
}
