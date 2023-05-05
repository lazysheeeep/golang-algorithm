package main

import "fmt"

func main() { //暴力法
	var n int
	fmt.Scan(&n)
	var num []int
	for i := 0; i < n; i++ {
		var temp int
		fmt.Scan(&temp)
		num = append(num, temp)
	}
	var target int
	var index1, index2 int
	fmt.Scan(&target)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if num[i]+num[j] == target {
				index1 = i
				index2 = j
			}
		}
	}
	fmt.Print(index1, index2)
}
