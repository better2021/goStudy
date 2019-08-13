package lesson

import (
	"fmt"
	"reflect"
)

type A struct {
	B
	Name string
}

type B struct {
	Name string
}

type TZ int

// *定义指针类型，&取地址
func (a *A) print(){ // a A跟结构相互绑定
	a.Name = "AA"
	fmt.Println(a)
}

func(b TZ) console(){
	b = 10 + 5
	fmt.Println(b)
}

// 定义接口类型
type USB interface {
	Name() string
	Connecter
}

type PhoneConnecter struct {
	name string
}

type Connecter interface {
	Connect()
}

func (pc PhoneConnecter) Name() string{
	return pc.name
}


// 反射
type User struct {
	Id int
	Name string
	Age int
}

func (u User) Hello()  {
	fmt.Println(u)
}

func Info(o interface{})  {
	t := reflect.TypeOf(o)
	fmt.Println(t.Name(),"--")

	v := reflect.ValueOf(o)
	fmt.Println(v)
}

func Lesson(){
	a := A{Name:"A",B:B{Name:"B"}}
	fmt.Println(a.Name,a.B.Name)
	a.print()

	var b TZ = 5
	b.console()

	u := User{1,"ok",12}
	Info(u)

	// 定义map
	m1 := make(map[int]string)
	m1[1] = "ad"
	m1[2] = "go"
	m1[3] = "js"
	fmt.Println(m1,len(m1))

	m2 := map[int]string{
		3:"java",
		4:"c++",
		5:"go",
		6:"js",
	}
	fmt.Println(m2)

	for key,val:= range m2{
		fmt.Println(key,val,"---")
	}

	val,ok := m2[5]
	if ok == true{
		fmt.Println(val)
	}else {
		fmt.Println("not find key")
	}
	// 取map的值
	fmt.Println(m2[6],"+++")
	delete(m2,6) // 删除key为6的键值
	fmt.Println(m2[6],"+++")
}