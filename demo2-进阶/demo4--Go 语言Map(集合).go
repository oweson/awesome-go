package main

import "fmt"

func test1() {
	map1 := make(map[string]string, 5)
	map2 := make(map[string]string)
	map3 := map[string]string{}
	map4 := map[string]string{"a": "1", "b": "2", "c": "3"}
	fmt.Println(map1, map2, map3, map4)
}
func main() {
	/* 1 Map 是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。

	Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，我们无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的。*/
	/* 2 定义 Map
	可以使用内建函数 make 也可以使用 map 关键字来定义 Map:

	声明变量，默认 map 是 nil
	var map_variable map[key_data_type]value_data_type

	 使用 make 函数
	map_variable := make(map[key_data_type]value_data_type)
	如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对*/
	hi := map[int]string{}
	hello := make(map[int]string)
	hi[1] = "ppx"
	hi[2] = "ppa"
	hello[1] = "silly"
	hello[2] = "tom"
	/*使用键输出地图值 */
	for key := range hi {
		fmt.Println(key, "value is ;", hi[key])
	}

	/*查看元素在集合中是否存在 */
	/*如果确定是真实的,则存在,否则不存在 */
	/*fmt.Println(capital) */
	/*fmt.Println(ok) */
	//capital,ok:= hi[2]
	//
	//if (ok) {
	//	fmt.Println("American 的首都是", capital)
	//} else {
	//	fmt.Println("American 的首都不存在")
	//}
}
