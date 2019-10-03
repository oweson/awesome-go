package main

import (
	"fmt"
)

/* 1 函数返回两个数的最大值 */
func max(num1 int64, num2 int64) int64 {
	/* 声明局部变量 */
	var result int64

	if (num1 > num2) {
		result = num1
	} else {
		result = num2
	}
	return result
}

/* 2 两数的和*/
func add(num1, num2 int64) int64 {
	var result int64
	if (num1 != 0 && num2 != 0) {
		result = num1 + num2
		return result
	} else {
		return 0
	}

}
func swap(x, y string) (string, string) {
	return y, x
}
func swapInt(x, y int) (int, int) {
	return y, x

}
func main() {
	var a, b int64 = 100, 100
	i := max(a, b)
	fmt.Print(i)
	i2 := add(a, b)
	fmt.Println(i2)
	// 3 函数返回多个值
	i3, i4 := swapInt(100, 200)
	fmt.Println(i3, i4)
}
