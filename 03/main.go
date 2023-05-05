package main

import (
	"fmt"
)

func main() { //暴力拆解
	var str string
	fmt.Scan(&str)
	var length int
	length = 1
	var flag bool
	for i := 0; i < len(str); i++ {
		for j := i; j <= len(str); j++ { //取所有的情况
			temp := str[i:j]
			fmt.Println(temp)
			flag = false
			for index := 0; index < len(temp)-1; index++ { //判断是否有相同字符 index是被判断字符
				for k := index + 1; k < len(temp); k++ {
					if temp[k] == temp[index] {
						flag = true
						break
					}
				}
				if flag {
					break
				}
			}
			if !flag {
				if len(temp) > length {
					fmt.Println(temp, "don't has")
					length = len(temp)
				}
			}
		}
	}
	fmt.Println(length)
}
