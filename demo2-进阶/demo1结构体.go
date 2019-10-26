package main

import (
	"fmt"
)

type Books struct {
	title   string
	author  string
	subject string
}

func main() {
	/* 1 Go 语言中数组可以存储同一类型的数据，但在结构体中我们可以为不同项定义不同的数据类型。

	结构体是由一系列具有相同类型或不同类型的数据构成的数据集合。

	结构体表示一项记录，比如保存图书馆的书籍记录，每本书有以下属性：type Books struct {
	   title string
	   author string
	   subject string
	   book_id int
	};

	2 // 创建一个新的结构体
	    fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})

	    // 也可以使用 key => value 格式
	    fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})

	    // 忽略的字段为 0 或 空
	   fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})*/
	// 创建一个新的结构体
	fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程"})

	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程"})
	fmt.Println("java是世界上最好的语言！！！")
	/* 3 访问结构体成员*/
	var book1 Books
	var book2 Books
	book1.title = "java最好的语言"
	book2.title = "php最差的语言"
	fmt.Println("book1的标题是：%s", book1.title)
	fmt.Println("book2的标题是：%s", book2.title)
	/** 4 结构体作为函数参数
	你可以像其他数据类型一样将结构体类型作为参数传递给函数。并以以上实例的方式访问结构体变量*/
	sendMsg(book2)

	for a := 0; a < 10; a++ {
		sendMsg(book1)
	}

}
func sendMsg(books Books) {
	fmt.Println("book1的标题是：", books.title)

}
