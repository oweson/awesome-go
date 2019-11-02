package main

import "fmt"

func main() {
	/* 1 Go 语言切片是对数组的抽象。

	Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，
	Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。*/
	/*切片初始化
	s :=[] int {1,2,3 }
	直接初始化切片，[]表示是切片类型，{1,2,3}初始化值依次是1,2,3.其cap=len=3*/
	var s = []int{1, 2, 3}
	fmt.Println(s)
	/* 2 切片不需要说明长度。

	或使用make()函数来创建切片:

	var slice1 []type = make([]type, len)

	也可以简写为

	slice1 := make([]type, len)
	也可以指定容量，其中capacity为可选参数。

	make([]T, length, capacity)
	这里 len 是数组的长度并且也是切片的初始长度。*/
	var oweson []int = make([]int, 10)
	fmt.Println(oweson)
}
