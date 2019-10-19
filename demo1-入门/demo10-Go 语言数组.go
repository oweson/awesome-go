package main

import "fmt"

func main() {
	/* 1 Go 语言数组声明需要指定元素类型及元素个数*/
	var intArr [100]int64
	fmt.Print(intArr[0])
	/* 2
	数组初始化：
	var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	初始化数组中 {} 中的元素个数不能大于 [] 中的数字。

	如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小：

	 var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	该实例与上面的实例是一样的，虽然没有设置数组的大小。;
	以上实例读取了第五个元素。数组元素可以通过索引（位置）来读取（或者修改），索引从0开始，第一个元素索引为 0，第二个索引为 1，以此类推。*/
	var balance = [10]int{1, 2, 3, 4, 5, 6}
	fmt.Print(balance[1])
	var guess = [...]int{1, 2, 3}
	fmt.Print(len(guess))
	/* 3  var n [10]int  n 是一个长度为 10 的数组
	var i,j int

	为数组 n 初始化元素
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* 设置元素为 i + 100
	}

	/输出每个数组元素的值
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j] )
	}*/
	var arr [100]int
	var i, j int
	for i = 0; i < 100; i++ {
		arr[i] = i + 100
	}
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, arr[j])
	}

}
