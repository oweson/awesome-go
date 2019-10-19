package main

import "fmt"

/*全局变量
在函数体外声明的变量称之为全局变量，全局变量可以在整个包甚至外部包（被导出后）使用。

全局变量可以在任何函数中使用，以下实例演示了如何使用全局变量：
Go 语言程序中全局变量与局部变量名称可以相同，但是函数内的局部变量会被优先考虑*/
/* 声明全局变量,有默认值 */
var g int
var df41 int

func main() {
	/* 1 Go 语言中变量可以在三个地方声明：

	函数内定义的变量称为局部变量
	函数外定义的变量称为全局变量
	函数定义中的变量称为形式参数*/
	/* 2 在函数体内声明的变量称之为局部变量，它们的作用域只在函数体内，参数和返回值变量也是局部变量。

	以下实例中 main() 函数使用了局部变量 a, b, c：*/
	/* 声明局部变量 */
	var a, b, c int

	/* 初始化参数 */
	a = 10
	b = 20
	c = a + b

	fmt.Printf("结果： a = %d, b = %d and c = %d\n", a, b, c)
	var x, y, z int = 1, 2, 3
	fmt.Printf("结果是：x=%d,y=%d,z=%d\n", x, y, z)
	fmt.Printf("全局变量df41的默认值是：%d", df41)
	/*形式参数会作为函数的局部变量来使用。实例如下：

	实例
	package main

	import "fmt"

	声明全局变量
	var a int = 20;

	func main()
	main 函数中声明局部变量
		var a int = 10
		var b int = 20
		var c int = 0

		fmt.Printf("main()函数中 a = %d\n",  a);
		c = sum( a, b);
		fmt.Printf("main()函数中 c = %d\n",  c);*/

}
