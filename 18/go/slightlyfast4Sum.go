package main

import "sort"

func SlightlyFastSumRecursive(totDepth,curDepth,target,curSum int,curCombination,curSlice []int, prtn **[][]int){
  if curDepth == totDepth {
		if curSum==target  {
			rtn := **prtn
			if len(rtn)==0 || !IsSame(curCombination,rtn[len(rtn)-1]) {
				toAppend:=make([]int,totDepth)
				copy(toAppend,curCombination)
				rtn = append(rtn, toAppend)
				*prtn=&rtn
			}
		}
		return
	}

	var last int
	if len(curSlice) != 0{
		last=curSlice[0]-1
	}
	for i:=0;i<len(curSlice);i++{
    nextSum:=curSum+curSlice[i]
    if curSlice[i]>= 0 && nextSum > target {
      return
    }

		if  curSlice[i] != last {
			curCombination[curDepth]=curSlice[i]
			SlightlyFastSumRecursive(totDepth,curDepth+1,target,nextSum,curCombination,curSlice[i+1:],prtn)
			last=curSlice[i]
		}
	}
}

func SlightlyFast4Sum(nums []int,target int) [][]int  {
	rtn:=make([][]int,0)

	inputLen:=len(nums)
	if inputLen < 4{
		return rtn
	}

	if inputLen == 4{
		if Sum(nums) == target {
			return append(rtn,nums)
		} else{
			return rtn
		}
	}

	sort.Ints(nums)
	curCombination:=make([]int,4)

	prtn:=&rtn
	pprtn:=&prtn

	SlightlyFastSumRecursive(4,0,target,0,curCombination,nums,pprtn)

	return **pprtn
}
