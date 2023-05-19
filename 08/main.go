package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	var temp int
	temp = num
	var rev int
	rev = 0
	var point int
	for {
		if num == 0 {
			break
		}
		point = num % 10
		rev = rev*10 + point
		num = num / 10
	}
	if rev == temp {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
