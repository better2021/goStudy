package lesson

import (
	"fmt"
)

func Variable(){
	var a,b int = 5,6
	var s string ="go"
	fmt.Println(a,b,s,"--")

	m1 :=map[int]string{1:"a",2:"b",3:"c",4:"d"}
	fmt.Println(m1)
	m2 := make(map[string]int)

	for k,v := range m1{
		m2[v] = k
	}
	fmt.Println(m2)
	ad("哈哈",3,4,5)

	cc := func() {
		fmt.Println("这是一个匿名函数")
	}
	cc() // 执行匿名函数

	//ff := clo(10)
	//fmt.Println(ff(5))

	//def()

	tt := &test{Address:"深圳"}
	tt.Name = "feiyu"
	tt.Age = 18
	tt.Contact.Phone = "13246586235"
	tt.Contact.City = "wuhan"
	tt.human.Sex = "girl"
	fmt.Println(tt)
	AA(tt)
	fmt.Println(tt)

	// 匿名结构体
	qq := struct {
		Name string
		Age int
	}{
		Name:"coco",
		Age:20,
	}
	fmt.Println(qq)
}

// ...int 可以传入多个int参数
func ad(a string,b ...int){
	fmt.Println(a,b)
}

// 闭包
func clo(x int) func(int) int  {
	return func(y int) int {
		return x + y
	}
}

func def()  {
	for i:=0;i<5 ;i++  {
		defer fmt.Println(i) // defer 新定义的后执行，后定义的先执行，后来居上
	}
}

/*
结构体
*/
type human struct {
	Sex string
}

type test struct {
	human
	Name string
	Age int
	Address string
	Contact struct{
		Phone, City string
	}
}

func AA(per *test){
	per.Age = 24
	fmt.Println(per,"--")
}