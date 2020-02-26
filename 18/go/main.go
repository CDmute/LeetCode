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
		{
			Nums:   []int{7, 7, -1, -5, 2, -2, 8, -8, -10, 0, 1, -4, -1, 4, -6, 5, 4},
			Target: -21,
			Results: [][]int{
				[]int{-10, -8, -5, 2},
				[]int{-10, -8, -4, 1},
				[]int{-10, -8, -2, -1},
				[]int{-10, -6, -5, 0},
				[]int{-10, -6, -4, -1},
				[]int{-10, -5, -4, -2},
				[]int{-8, -6, -5, -2},
			},
		},
		{
			Nums:   []int{-479, -472, -469, -461, -456, -420, -412, -403, -391, -377, -362, -361, -340, -336, -336, -323, -315, -301, -288, -272, -271, -246, -244, -227, -226, -225, -210, -194, -190, -187, -183, -176, -167, -143, -140, -123, -120, -74, -60, -40, -39, -37, -34, -33, -29, -26, -23, -18, -17, -11, -9, 6, 8, 20, 29, 35, 45, 48, 58, 65, 122, 124, 127, 129, 145, 164, 182, 198, 199, 206, 207, 217, 218, 226, 267, 274, 278, 278, 309, 322, 323, 327, 350, 361, 372, 376, 387, 391, 434, 449, 457, 465, 488},
			Target: 1979,
		},
	}

	for _, testCase := range testCases {
		fmt.Println("Expected")
		fmt.Println(testCase.Results)
		fmt.Println("Caculated")
		//fmt.Println(fastFourSum(testCase.Nums, testCase.Target))
		fmt.Println(SlightlyFast4Sum(testCase.Nums, testCase.Target))
	}
}
