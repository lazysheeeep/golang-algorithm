package main

import (
	"fmt"
	"github.com/lazysheeeep/Stack"
	"math"
)

func main() {
	var str string
	fmt.Scan(&str)
	var length int
	length = len(str)
	var max float64
	max = 0
	var result string
	for i := 0; i < length; i++ {
		for j := i; j <= length; j++ { //不会发生溢出，因为字符串最后还有一个字符
			temp := str[i:j]
			stack := Stack.NewStack(len(temp))
			for k := 0; k < len(temp); k++ { //字符串存入到栈
				stack.Push(temp[k])
			}
			//开始判断回文
			var flag bool
			flag = true
			for k := 0; k < len(temp); k++ {
				if temp[k] != stack.Pop() {
					flag = false
					break
				}
			}
			if flag {
				if float64(len(temp)) > max {
					max = math.Max(max, float64(len(temp)))
					result = temp
				}
			}
		}
	}
	fmt.Println(result)
}
