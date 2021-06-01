package util

import "sort"

//ID号去重
func IDWeight(ids []int) []int {
	var NewInts []int

	tempMap := make(map[int]bool,len(ids))
	for _,v := range ids {
		if tempMap[v] == false {
			tempMap[v] = true
			NewInts = append(NewInts,v)
		}
	}
	return NewInts
}

//ID号排序，降序
func DescIDArr(ids []int) []int{
	sort.Sort(sort.Reverse(sort.IntSlice(ids)))
	return ids
}