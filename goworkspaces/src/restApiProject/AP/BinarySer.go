package main

import "fmt"

func BinaArySearch(aa []int, target int) int {
	//not found
	start := 0
	last := len(aa) - 1
	mid := start + last/2
	for start <= last {

		value := mid
		if aa[mid] == target {

			fmt.Println("index", value)
			return value
		}
		if aa[value] > target {
			last = mid - 1
			mid = (start + mid) / 2
			fmt.Println("index", last)
		}
		start = mid + 1
		mid = (start + mid) / 2
		fmt.Println(start)
	}
	return -1
}
func main() {
	as := []int{1, 2, 3, 4, 7, 12, 56, 60}
	aa := as
	f := BinaArySearch(aa, 60)
	fmt.Println("fund", f)
}
