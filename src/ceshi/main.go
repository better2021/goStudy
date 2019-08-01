package main

import (
	"calc"
	"fmt"
	. "fmt" // 用.引入时可以省略前缀，用_引入只会执行包中的init()函数
	"time"
)

// 分组式赋值
var (
	a string = "分组式赋值"
	b int    = 2
	c bool   = false
)

var e, f = 10, "str" // go语言会自动推断类型

const name = "慕课网"          // 常量的隐式声明(无type声明)
const title string = "go语言" // 常量的显式声明
// 常量的分组声明
const (
	cat    string = "猫"
	dog           = "狗"
	person string = "人"
	num    int    = 10
)

func main() {
	x, y := 7, 8 // := 简写的赋值方式相当于 var x,y int = 7,8
	Println("good")
	addres := calc.Add(2, 3)
	subres := calc.Sub(6, 3)
	Println(addres, subres, time.Now())
	Println(calc.Hello("go"))
	Println(a, c, b, e, f, x, y)
	Println(len(name), title, cat, dog, person, num)

	if x > y {
		Println("x大于y")
	} else {
		Println("x小于y")
	}

	var number interface{}
	number = 3.14
	switch number.(type) {
	case int:
		Println("整形")
	case string:
		Println("字符串型")
	case bool:
		Println("布尔型")
	default:
		Println("都不是")
	}

	// for {
	// 	Println("for循环没有条件会进行无限循环")
	// 	time.Sleep(1 * time.Second) // 1秒钟打印一次
	// }

	// goto 语句可以跳过代码块
	goto One
	Println("中间语句")

	for i := 0; i <= 5; i++ {
		Println("会打印5次")
		time.Sleep(2 * time.Second)
		if i == 2 { // 当i等于2时，break语句会终止循环
			break
		}
	}

One:
	Println("跳过了for语句")

	// range可以达到forEach的效果
	fruit := []string{"香蕉", "苹果", "梨子"}
	for key, value := range fruit {
		Println(key, value)
	}
	Println(fruit[1])

	Println(calc.Params(100, "多返回值"))

	// go的并发，gorunting
	for i := 0; i < 100; i++ {
		//gorunting
		go printHello(i)
	}
	time.Sleep(time.Second)
}

func printHello(i int) {
	fmt.Println("Hello go 并发", i)
}
