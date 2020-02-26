package main

import (
	"sort"
)

func DP4Sum(nums []int, target int) [][]int {
	return DPAnySum(nums, target, 4)
}

func printArray(arr []int) {
	for _, x := range arr {
		print(x)
		print(" ")
	}

	println()
}

func FindNegativeMaxAndPositiveMax(nums []int) (int, int) {
	for idx, num := range nums {
		if num > 0 {
			return Sum(nums[0:idx]), Sum(nums[idx:])
		}
	}

	return Sum(nums), 0
}

type DPtrace struct {
	ParentDepth   int
	ParentCombIdx int
	CurElem       int
}

func DPAnySum(nums []int, target, depth int) [][]int {
	rtn := make([][]int, 0)
	numLen := len(nums)
	sort.Ints(nums)

	//store rest avaliable nums
	dpNumCursorStack := make([][]int, depth+1)
	// store last combination
	dpCombinationStack := make([][]*DPtrace, depth+1)
	// store last combination's sum
	dpSumStack := make([][]int, depth+1)

	dpSumStack[0] = []int{0}
	dpCombinationStack[0] = []*DPtrace{&DPtrace{
		ParentDepth:   -1,
		ParentCombIdx: -1,
		CurElem:       0,
	}}
	dpNumCursorStack[0] = []int{-1}

	for i := 1; i <= depth; i++ {
		for combIdx, _ := range dpCombinationStack[i-1] {
			lastNum := nums[0] - 1
			lastCursor := dpNumCursorStack[i-1][combIdx] + 1
			for numIdx, num := range nums[lastCursor:] {
				if num != lastNum {
					lastSum := dpSumStack[i-1][combIdx]
					resetLen := depth - (i - 1)
					if lastSum+Sum(nums[numLen-resetLen:]) < target {
						break
					}
					if num >= 0 && num+lastSum > target {
						break
					}

					curCombinations := dpCombinationStack[i]
					curNumCursor := dpNumCursorStack[i]
					curSums := dpSumStack[i]
					if curCombinations == nil {
						curCombinations = make([]*DPtrace, 0)
						curNumCursor = make([]int, 0)
					}

					curSums = append(curSums, lastSum+num)
					dpSumStack[i] = curSums

					curCombination := &DPtrace{
						ParentDepth:   i - 1,
						ParentCombIdx: combIdx,
						CurElem:       num,
					}
					curCombinations = append(curCombinations, curCombination)
					//curCombinations = append(curCombinations, curCombination)
					dpCombinationStack[i] = curCombinations

					curNumCursor = append(curNumCursor, numIdx+lastCursor)
					dpNumCursorStack[i] = curNumCursor

					lastNum = num
				}
			}
		}
	}

	for idx, sum := range dpSumStack[depth] {
		if sum == target {
			comb := make([]int, depth)

			combCursor := dpCombinationStack[depth][idx]
			for i := depth - 1; i >= 0; i-- {
				comb[i] = combCursor.CurElem
				combCursor = dpCombinationStack[combCursor.ParentDepth][combCursor.ParentCombIdx]
			}

			rtn = append(rtn, comb)
		}
	}

	return rtn
}
