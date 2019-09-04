package main

import (
	"fmt"
	"math"
)

var datas = []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}

var goal = 23

type pointerSlice []int

var sliceA pointerSlice

var LoPointer, HiPointer, MidPointer int

func main() {

	fmt.Printf("\nSlice : %v\n\n", datas)

	sliceA = datas
	max_loop := int(math.Ceil(math.Log2(float64(len(sliceA)))))

	LoPointer = 0
	HiPointer = len(sliceA) - 1

	fmt.Println("[LoIndex]-[MidIndex]-[HiIndex]")
	for index := 0; index < max_loop; index++ {

		MidPointer = (HiPointer-LoPointer)/2 + LoPointer

		fmt.Printf("% 5d % 9d % 9d\n", LoPointer, MidPointer, HiPointer)
		if sliceA[MidPointer] > goal {
			HiPointer = MidPointer - 1
		} else if sliceA[MidPointer] < goal {
			LoPointer = MidPointer + 1
		} else {
			fmt.Println("")
			fmt.Println("Value :", sliceA[MidPointer], "at index:", MidPointer)
			break
		}
	}
}
