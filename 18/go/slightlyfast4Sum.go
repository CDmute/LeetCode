package main

import "sort"

func SlightlyFastSumRecursive(totDepth, curDepth, target, curSum int, curCombination, curSlice []int, prtn **[][]int) {
	curLen := len(curSlice)
	if curDepth == totDepth {
		if curSum == target {
			rtn := **prtn
			toAppend := make([]int, totDepth)
			copy(toAppend, curCombination)
			rtn = append(rtn, toAppend)
			*prtn = &rtn
		}
		return
	}

	var last int
	if len(curSlice) != 0 {
		last = curSlice[0] - 1
	}
	for i := 0; i < len(curSlice); i++ {
		restLen := totDepth - curDepth
		nextSum := curSum + curSlice[i]
		if curLen-restLen > 0 && curSum+Sum(curSlice[curLen-restLen:]) < target {
			return
		} else {
			continue
		}

		if curSlice[i] >= 0 && nextSum > target {
			return
		}

		if curSlice[i] != last {
			curCombination[curDepth] = curSlice[i]
			SlightlyFastSumRecursive(totDepth, curDepth+1, target, nextSum, curCombination, curSlice[i+1:], prtn)
			last = curSlice[i]
		}
	}
}

func SlightlyFast4Sum(nums []int, target int) [][]int {
	rtn := make([][]int, 0)

	inputLen := len(nums)
	if inputLen < 4 {
		return rtn
	}

	if inputLen == 4 {
		if Sum(nums) == target {
			return append(rtn, nums)
		} else {
			return rtn
		}
	}

	sort.Ints(nums)
	min, max := FindNegativeMaxAndPositiveMax(nums)
	if min > target || max < target {
		return [][]int{}
	}
	curCombination := make([]int, 4)

	prtn := &rtn
	pprtn := &prtn

	SlightlyFastSumRecursive(4, 0, target, 0, curCombination, nums, pprtn)

	return **pprtn
}
