package lesson

import (
	"fmt"
	"reflect"
	"runtime"
	"strconv"
)

type person struct {
	Human
	name string
	age int
	sex string
	skills
}

type Human struct {
	name string
	phone string
}

type skills []string

// methon定义method
func (h *Human) say(){
	fmt.Println(h.name,h.phone,"---")
}

func Four(){
	P := person{ Human:Human{name:"coco",phone:"1254545487"},name:"Tom",age:25}
	P.skills = []string{"1","3","5"}
	P.skills = append(P.skills,"go") // 向数组添加元素go
	P.sex = "gril"
	fmt.Println(P,P.name)
	P.say()
	const a = "3"
	stringToint,_ := strconv.Atoi(a) // 字符串转int
	fmt.Println(stringToint + 2)

	/*
	Go语言实现了反射，所谓反射就是能检查程序在运行时的状态
	获取反射值能返回相应的类型和数值
	*/
	var x float64 =3.5
	v := reflect.ValueOf(x)
	fmt.Println("type:",v.Type()) // 获取类型
	fmt.Println("value:",v.Float()) // 获取值

	// 反射要修改值，需要这样写
	var y float32 = 3.3
	q := reflect.ValueOf(&y)
	zz := q.Elem()
	zz.SetFloat(6.8)
	fmt.Println(zz)

	go sayHi("golang")
	sayHi("hello")

}

/*
 go 的并发编程
 goroutine
*/
func sayHi(s string){
	for i:=0;i<1000;i++{
		runtime.Gosched()
		fmt.Println(s,i)
	}
}