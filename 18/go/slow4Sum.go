package main

import "sort"

func SlowSumRecursive(totDepth,curDepth,target int,curCombination,curSlice []int,prtn **[][]int) {
	if curDepth == totDepth {

		if Sum(curCombination)==target  {
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
		//sum := Sum(curCombination[0:curDepth])
		if  curSlice[i] != last {
			curCombination[curDepth]=curSlice[i]
			SlowSumRecursive(totDepth,curDepth+1,target,curCombination,curSlice[i+1:],prtn)
			last=curSlice[i]
		}
	}
}


func slowFourSum(nums []int,target int) [][]int  {
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

	SlowSumRecursive(4,0,target,curCombination,nums,pprtn)

	return **pprtn
}
