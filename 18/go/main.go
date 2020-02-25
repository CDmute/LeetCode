package main

import (
	"fmt"
)

func Sum(s []int) (sum int) {
	for _, x := range s {
		sum += x
	}
	return sum
}

func IsSame(origin, target []int) bool {
	for idx, element := range origin {
		if element != target[idx] {
			return false
		}
	}

	return true
}

type TestCase struct {
	Nums    []int
	Target  int
	Results [][]int
}

func main() {
	fmt.Println("Begin test case")
	testCases := []TestCase{
		{
			Nums:   []int{1, -2, -5, -4, -3, 3, 3, 5},
			Target: -11,
			Results: [][]int{
				[]int{-5, -4, -3, 1},
			},
		},
		{
			Nums:   []int{1, 0, -1, 0, -2, 2},
			Target: 0,
			Results: [][]int{
				[]int{-2, -1, 1, 2},
				[]int{-2, 0, 0, 2},
				[]int{-1, 0, 0, 1},
			},
		},
	}

	for _, testCase := range testCases {
		fmt.Println("Expected")
		fmt.Println(testCase.Results)
		fmt.Println("Caculated")
		//fmt.Println(fastFourSum(testCase.Nums, testCase.Target))
    	fmt.Println(SlightlyFast4Sum(testCase.Nums,testCase.Target))
	}
}
