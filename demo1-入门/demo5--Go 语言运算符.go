package main

import "fmt"

func main() {
	var a, b = 1, 2
	fmt.Print("a add b is:", a+b)
	var c = a * b
	fmt.Printf("第三行 - c 的值为 %d\n", c)
	fmt.Printf("a add b is %d\n ",a+b)
	// 关系运算符
	/*var a int = 21
	  var b int = 10

	  if( a == b ) {
	     fmt.Printf("第一行 - a 等于 b\n" )
	  } else {
	     fmt.Printf("第一行 - a 不等于 b\n" )
	  }
	  if ( a < b ) {
	     fmt.Printf("第二行 - a 小于 b\n" )
	  } else {
	     fmt.Printf("第二行 - a 不小于 b\n" )
	  }

	  if ( a > b ) {
	     fmt.Printf("第三行 - a 大于 b\n" )
	  } else {
	     fmt.Printf("第三行 - a 不大于 b\n" )
	  }*/
	var x int=1
	var y int=2
	if(x<=y){
		fmt.Printf("x is less than y!\n")
		fmt.Printf("%d,%d,%d",a,x,y)
	}
	// 布尔类型
	/* var a bool = true
	   var b bool = false
	   if ( a && b ) {
	      fmt.Printf("第一行 - 条件为 true\n" )
	   }
	   if ( a || b ) {
	      fmt.Printf("第二行 - 条件为 true\n" )
	   }
	   /* 修改 a 和 b 的值 */
	/*a = false
	b = true
	if ( a && b ) {
		fmt.Printf("第三行 - 条件为 true\n" )
	} else {
		fmt.Printf("第三行 - 条件为 false\n" )
	}
	if ( !(a && b) ) {
		fmt.Printf("第四行 - 条件为 true\n" )
	}*/
	var df int=21
	var jl int =19
	if(df>=0 && jl>=0){
		fmt.Println("df is %d and jl is %d !\n",df,jl)
	}
	// 符号
	/*Go 的自增，自减只能作为表达式使用，而不能用于赋值语句。
	a++ // 这是允许的，类似 a = a + 1,结果与 a++ 相同
	a-- //与 a++ 相似
	a = a++ // 这是不允许的，会出现变异错误 syntax error: unexpected ++ at end of statement*/

}
