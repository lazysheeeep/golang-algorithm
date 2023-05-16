package main

import (
	"fmt"
)

func main() {
	var num string
	fmt.Scan(&num)
	if num < string(-2147483648) || num > string(2147483647) {
		fmt.Println(0)
		return
	}

	var flag bool //如果为负数则flag为true
	flag = false
	if num[0] == '-' {
		flag = true
	}

	if flag { //输出小于0时
		fmt.Printf("%c", num[0])
		for i := len(num) - 1; i > 0; i-- {
			fmt.Printf("%c", num[i])
		}
	} else {
		for i := len(num) - 1; i >= 0; i-- {
			fmt.Printf("%c", num[i])
		}
	}

}
