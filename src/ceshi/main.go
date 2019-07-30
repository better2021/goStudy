package main

import (
	"calc"
	. "fmt" // 用.引入时可以省略前缀，用_引入只会执行包中的init()函数
)

func main() {
	Println("good")
	addres := calc.Add(2, 3)
	subres := calc.Sub(6, 3)
	Println(addres, subres)
}
