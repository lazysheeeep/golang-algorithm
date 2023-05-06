package main

import (
	"fmt"
	"math"
	"strings"
)

func Easy(str string) { //因为go语言本身字符串截取的限制，所以舍弃字符串了，list也不行，服了
	//如果我不用字符串拼接就好了
	var (
		right int
		left  int
		count float64
		max   float64
		temp  string
	)

	right = 0
	left = 0
	count = 0
	max = 0

	for {
		if right >= len(str) {
			break
		}

		if !strings.Contains(temp, string(str[right])) {
			temp = temp + string(str[right])
			right++
			count++
		} else {
			temp = temp[1:]
			left++
			count--
		}

		max = math.Max(max, count)

	}
	fmt.Println(max)
}
