package main

import "fmt"

func main() {
	permute([]int{1, 2, 3})
	fmt.Println("111")
}

var res [][]int
var tmp []int

var lens int

var used []bool

func permute(nums []int) [][]int {
	lens = len(nums)
	res = make([][]int, 0)
	tmp = make([]int, 0)
	used = make([]bool, lens)
	backTrack(nums)

	return res
}

func backTrack(nums []int) {
	if lens == len(tmp) {
		target := make([]int, lens)
		copy(target, tmp)
		res = append(res, target)
		return
	}
	for i := 0; i < lens; i++ {
		if used[i] {
			continue
		}
		used[i] = true
		tmp = append(tmp, nums[i])
		backTrack(nums)
		tmp = tmp[:len(tmp)-1]
		used[i] = false

	}
}
