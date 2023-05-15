package main

import (
	"math"
)

func solution(str string) string {
	var (
		result string
		length int
		end    int
		mid    float64
		right  int
		left   int
	)
	result = ""
	length = len(str)
	//这里设计一个2*length-1是为了每一个地方都可以做中轴线
	//如果不设计这个，相当于我们截取出来的字符串只能为奇数个
	//比如如果是abba，实际上他能做中轴线的地方有'a','','b','','b','',a七种情况
	end = 2*length - 1
	for i := 0; i < end; i++ {
		mid = float64(i) / 2.0      //''作为中轴线时为x.5
		right = int(math.Ceil(mid)) //右指针向上取整
		left = int(math.Floor(mid)) //左指针向下取整,如果是字符做中轴线，则取原来的数字
		for {
			if left < 0 || right >= length {
				break
			}
			if str[left] != str[right] {
				break
			}
			right++
			left--
		}
		var l int
		l = right - left - 1
		if l > len(result) {
			result = str[left+1 : right] //right不减一，因为截取不会截取到右边的
		}
	}
	return result
}
