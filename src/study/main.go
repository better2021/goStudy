package main

import (
	"fmt"
	"github.com/feiyuWeb/beego-gin/src/study/lesson"
	"unicode/utf8"
)

func main() {
	//variable()
	//a,b := 3,4
	//a,b = swap(a,b)
	//fmt.Println(a,b)

	//var array1 [5]int
	//array1[0] = 100
	//array2 := [3]int{1,2,3}
	//array3 := [...]int{1,3,5,6}
	//fmt.Println(array1,array2,array3,array3[1])

	//for i:=0;i<len(array3);i++{
	//	fmt.Println(array3[i])
	//}

	// 数组的遍历
	//sum := 0
	//for i,v:=range array3{ // 通过_省略变量 for _,v:=range array3
	//	fmt.Println(i,v,"--")
	//	sum += v
	//}
	//fmt.Println(sum)

	//数组的切片
	//arr := [...]int{1,2,3,4,5,6,7}
	//s := arr[2:6]
	//fmt.Println(s,len(s),cap(s),"+--") // [3 4 5 6]
	//fmt.Println(arr[:6])
	//fmt.Println(arr[2:])
	//fmt.Println(arr[:])
	//s2 := append(s,10,15)
	//fmt.Println(s2)
	//
	//var numbers = make([]int,3,8) //make定义数组切片 make([]T, length, capacity)
	//println(numbers)

	/*
		Map 是一种无序的键值对的集合
		map[string]string
		[string]表示key的类型，string表示value的类型
	*/
	//m := map[string]string{
	//	"name":"go",
	//	"desc":"haha",
	//	"address":"深圳",
	//}

	// map类型也可以用make()来创建
	//m2 := make(map[string]int)
	//var m3 map[string]int
	//fmt.Println(m,m2,m3)
	//fmt.Println(m["address"])

	// 判断map类型m中是否含有key为desc的元素，有则打印出来对应的value值
	//if name,ok := m["desc"];ok{
	//	fmt.Println(name,"-*-")
	//}else {
	//	fmt.Println("key not find")
	//}
	//m["address"] = "武汉"
	//for k,v := range m{
	//	fmt.Println(k,v,"--")
	//}
	//
	str := "good很好"
	fmt.Println(len(str))
	fmt.Println(utf8.RuneCountInString(str)) // RuneCountInString 获取字符的数量
	//for _,v:= range []rune(str){
	//	fmt.Println(v,"--")
	//	fmt.Printf("%#U 起始于字位置%d\n", v)
	//}

	/*
		结构体
	*/
	//type treeNode struct {
	//	value int
	//	left,right *treeNode
	//}
	//
	//var root treeNode
	//root = treeNode{value:3}
	//root.left = &treeNode{}
	//root.right = &treeNode{5,nil,nil}
	//fmt.Println(root)

	// lesson.Variable()
	// lesson.Lesson()
	// lesson.Run()
	// lesson.Four()
	// lesson.Ostest()
	// lesson.Str()
	lesson.Fn()

}

//func variable(){
//	var a,b int = 5,6
//	var s string ="go"
//	fmt.Printf("%d %q\n",a,b,s)
//}

//const (
//	e = iota
//	f = iota
//)
//
//func swap(x,y int)(int,int){
//	return y,x
//}
