package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	/* 1 行分隔符
	在 Go 程序中，一行代表一个语句结束。每个语句不需要像 C 家族中的其它语言一样以分号 ; 结尾，因为这些工作都将由 Go 编译器自动完成。

	如果你打算将多个语句写在同一行，它们则必须使用 ; 人为区分，但在实际开发中我们并不鼓励这种做法。

	以下为两个语句：*/
	fmt.Println("tomcat")
	fmt.Printf("rose")
	/* 2 标识符==java
	标识符用来命名变量、类型等程序实体。
	一个标识符实际上就是一个或是多个字母(A~Z和a~z)数字(0~9)、下划线_组成的序列，但是第一个字符必须是字母或下划线而不能是数字。*/
	/* 3 字符串连接==java*/
	/* 4 关键字：25
	break	default	func	interface	select
	case	defer	go	map	struct
	chan	else	goto	package	switch
	const	fallthrough	if	range	type
	continue	for	import	return	var*/
	/* 5 程序一般由关键字、常量、变量、运算符、类型和函数组成。

	程序中可能会使用到这些分隔符：括号 ()，中括号 [] 和大括号 {}。

	程序中可能会使用到这些标点符号：.、,、;、: 和 …。*/
	/* 6 Go 语言的空格:
	Go 语言中变量的声明必须使用空格隔开，如：

	var age int;
	语句中适当使用空格能让程序更易阅读。

	无空格：

	fruit=apples+oranges;
	在变量与运算符间加入空格，程序看起来更加美观，如：

	fruit = apples + oranges; */

	var age int
	// 默认值0类似java全局的变量初始化0；
	println(age)
	age = 101
	println(age)

	println("*********************************************************")
	var name string
	name = "皮皮程"
	println(name)
}
