package lesson

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"image/color"
	"log"
)

// 定义一个结构体类型
type Student struct {
	Id int
	Name string
	Sex string // 字符串类型
	Age int
	Address string
}

func Run(){
	// 先定义一个普通结构体变量
	var s Student
	// 在定义一个指针变量，保存s的地址
	var p1 *Student
	p1 = &s
	// 通过指针操作成员， p1.id和（*p1）.id完全相同，只能使用.运算符
	 p1.Id = 1
	 p1.Age =  18
	 p1.Sex = "girl"
	 p1.Name = "go语言"
	 p1.Address = "武汉"
	fmt.Println(p1)

	 // 通过news生成一个新的结构体
	 p2 := new(Student)
	p2.Id = 2
	p2.Age =  20
	p2.Sex = "girl2"
	p2.Name = "go语言2"
	p2.Address = "武汉2"
	fmt.Println(p2)
	test01(s)

	a,b := foo(3,6)
	fmt.Println(a,b,a+b)

	// 生成二维码图片
	qrcode.WriteFile("http://www.flysnow.org/",qrcode.Medium,256,"./blog_qrcode.png")

	qr,err:=qrcode.New("https://feiyuweb.me/hexoView/",qrcode.Medium)
	if err != nil {
		log.Fatal(err)
	} else {
		qr.BackgroundColor = color.RGBA{230,217,150,1}
		qr.ForegroundColor = color.Black
		qr.WriteFile(256,"./qrcode.png")
	}
}

func test01(s Student){
	s.Id = 5
	s.Age = 25
	fmt.Println(s)
}

// go的多返回值
func foo(a,b int)(int,int){
	return  a,b
}


