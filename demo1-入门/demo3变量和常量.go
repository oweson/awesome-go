package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	/* 1 Go 语言变量
	Go 语言变量名由字母、数字、下划线组成，其中首个字符不能为数字。
	声明变量的一般形式是使用 var 关键字：*/
	var number int64;
	println(number)
	// 可以一次声明多个变量：
	var a, b, c int
	println(a + b + c)
	/* var a string = "Runoob"
	   fmt.Println(a)

	   var b, c int = 1, 2
	   fmt.Println(b, c)*/
	var ab = "runboot"
	var k, j int = 100, 200
	println(ab)
	println(k + j)
	/*变量声明
	第一种，指定变量类型，如果没有初始化，则变量默认为零值。*/
	var vr int
	var string string
	println(vr)
	// int默认值为0
	println(string + "------------------")
	// 字符串默认值为空
	/*默认值：数值类型（包括complex64/128）为 0

	布尔类型为 false

	字符串为 ""（空字符串）

	以下几种类型为 nil：
	var a *int
	var a []int
	var a map[string] int
	var a chan int
	var a func(string) int
	var a error // error 是接口*/

	/* var i int
	   var f float64
	   var b bool
	   var s string
	   fmt.Printf("%v %v %v %q\n", i, f, b, s)*/
	fmt.Printf("%v %v %v %v\n", a, b, c, vr)
	// format 输出；
	var op = true
	println(op)
	// 第二种，根据值自行判定变量类型。

	/*第三种，省略 var, 注意 := 左侧如果没有声明新的变量!!!就产生编译错误，格式：

	v_name := value*/
	vivo := 899999
	println(vivo)
	/*vivo := 999
	println(vivo)*/
	// intVal,intVal1 := 1,2 // 此时不会产生编译错误，因为有声明新的变量，因为 := 是一个声明语句
	f := "f-16"
	fmt.Printf(f)

	/*多变量声明:
	//类型相同多个变量, 非全局变量
	var vname1, vname2, vname3 type
	vname1, vname2, vname3 = v1, v2, v3

	var vname1, vname2, vname3 = v1, v2, v3 // 和 python 很像,不需要显示声明类型，自动推断

	vname1, vname2, vname3 := v1, v2, v3 // 出现在 := 左侧的变量不应该是已经被声明过的，否则会导致编译错误


	// 这种因式分解关键字的写法一般用于声明全局变量
	var (
	    vname1 v_type1
	    vname2 v_type2
	)*/
	var nubia, honor = 2500, 1500
	fmt.Printf("%v %v", nubia, honor)
	var (
		jj  int
		jjj int
	)
	println(jj + jjj)

}
