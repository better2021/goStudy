package main

import (
	"fmt"
	"math/rand"
)

func main() {
	random()
	arr := [5]string{"马龙", "迈克尔乔丹", "雷吉米勒", "蒂姆邓肯", "科比布莱恩特"}
	b := maxNum(arr)
	fmt.Println("数组中最长的字符串是", b)
	/*
		切片：切片与数组相比切片的长度是不固定的，可以追加元素，
		在追加时可能使切片的容量增大，所以可以将切片理解成“动态数组
		append()函数:第一个参数表示向哪个切片追加数据，后面表示具体追加的数据。
	*/
	s := []int{1, 2}
	s = append(s, 3, 5)
	fmt.Println(s)
	/*
		切片截取
		s = s[low : high : max] 切片的三个参数的切片截取的意义为
		low为截取的起始下标（含），
		 high为窃取的结束下标（不含high），
		 max为切片保留的原切片的最大下标（不含max）
	*/
	sli := []int{10, 20, 30, 40, 50}
	slice := sli[1:4:5]
	fmt.Println(slice, len(slice), cap(slice))

	/*
		结构体定义需要使用 type 和 struct 语句。
		struct 语句定义一个新的数据类型，结构体有中有一个或多个成员。
		type 语句设定了结构体的名称
		忽略的字段为 0 或 空
	*/
	type Books struct {
		title   string
		author  string
		subject string
		book_id int
	}
	// 创建一个新的结构体
	fmt.Println(Books{"go语言", "dva", "教程", 15})
	// 也可以使用key => value 格式
	fmt.Println(Books{title: "go语言", author: "dva", subject: "教程", book_id: 15})

	// 访问结构体成员
	var book Books = Books{"go语言", "dva", "教程", 15}
	/* book 描述 */
	// book.title = "go语言a"
	// book.author = "dva"
	// book.subject = "教程a"
	// book.book_id = 10
	fmt.Println(book, book.subject, "--")
}

// 生成随机数
func random() {
	num := rand.Int()
	fmt.Println("随机数", num)
}

/*
用方法来实现：有一个字符串数组：
{ "马龙", "迈克尔乔丹", "雷吉米勒", "蒂姆邓肯", "科比布莱恩特" },
请输出最长的字符串
*/

func maxNum(a [5]string) string {
	var max string
	max = a[0]
	for i := 0; i < len(a); i++ { // len(a) 获取数组a的长度,也可获取字符串的长度
		if len(a[i]) > len(max) {
			max = a[i]
		}
	}
	return max
}
