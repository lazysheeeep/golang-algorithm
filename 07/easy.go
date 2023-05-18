package main

import "fmt"

func solution(num int32) {
	var rev int32
	rev = 0
	var point int32
	for {
		if num == 0 {
			break
		}
		point = num % 10
		rev = rev*10 + point
		num = num / 10
	}
	fmt.Println(rev)
}
