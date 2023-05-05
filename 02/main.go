package main

import (
	list2 "container/list"
	"fmt"
)

func main() {
	list1 := list2.List{} //初始化两个链表
	list2 := list2.List{}
	list1.Init()
	list2.Init()

	var l1 string
	var l2 string
	l1 = "243" //这里输入两个数字
	l2 = "564"

	for i := 0; i < len(l1); i++ { //逆序存储
		var temp1 int
		temp1 = int(l1[i]) - 48
		list1.PushFront(temp1)
	}
	for i := 0; i < len(l2); i++ {
		var temp2 int
		temp2 = int(l2[i]) - 48
		list2.PushFront(temp2)
	}
	for i := list1.Back(); i != nil; i = i.Prev() {
		fmt.Print(i.Value)
	}
	//这道题目中道崩殂在不知道怎么把any类型进行转换
}
