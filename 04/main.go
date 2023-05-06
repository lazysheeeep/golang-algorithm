package main

import (
	"fmt"
	"sort"
)

func main() {
	var (
		n1     int
		n2     int
		num1   []int
		num2   []int
		temp   int
		sum    []int
		result float64
	)

	fmt.Scan(&n1)
	for i := 0; i < n1; i++ { //数据存储
		fmt.Scan(&temp)
		num1 = append(num1, temp)
	}

	fmt.Scan(&n2)
	for i := 0; i < n2; i++ {
		fmt.Scan(&temp)
		num2 = append(num2, temp)
	}

	sum = num1
	for i := 0; i < n2; i++ { //数据拼接
		sum = append(sum, num2[i])
	}

	sort.Ints(sum) //排序

	if len(sum)%2 != 0 { //奇数
		result = float64(sum[(len(sum)-1)/2])
		fmt.Println()
	} else {
		result = (float64(sum[(len(sum)/2)]) + float64(sum[len(sum)/2-1])) / 2.0
		fmt.Println(result)
	}
}
