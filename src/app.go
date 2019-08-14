package main // 声明文件的package
// 通过import 关键字导入其他非main包
import (
	"fmt"
	"github.com/feiyuWeb/beego-gin/src/calc"
)

/*
环境变量 GOPATH的配置很重要
GOPATH 下的src是存放项目源文件的
go env 命令可查看GOPATH的路径
例如：GOPATH=C:\Users\feiyu\go;E:\lesson\beego-gin;
GOPATH  可以写多路劲用分号隔开；
C:\Users\feiyu\go 下面有有个文件bin，pkg，src
bin：用于存放.exe可执行文件
pkg：用于存放package包文件，go get XXXX 获取的包会自动存在GOPATH的第一个路劲的pkg下
src：用于存放自己编写的go源代码文件
*/

func main() { // 入口函数, 无参数无返回值
	fmt.Println("hello go!")
	var b, c int = 1, 2
	f := "good"
	var d = add(b, c)
	var e = calc.Sub(c, b)
	var arr = [3]int{3, 5, 9}
	fmt.Println(d, e, f, arr[1]) // 数组 [3]int{3,5,9} [3]数组的长度 in数组中值的类型 {3,5,9}数组具体的值
	fmt.Println(max(3, 6))
	fmt.Println("变量的地址：", &c)
	fmt.Println(calc.Add(b, c))
}

/*
同一行声明多个变量和赋值
例如： var a,b,c int = 1,2,3 或者a,b := 1,2
全局变量的声明必须要var关键词，局部变量则可以省略
也就是说 a,b := 1,2 只能用于函数的局部变量
max,add以小写的字母开头的方法就为私有方法只能在本包中自己调用，不可导出给其他包调用
*/

func max(num1, num2 int) int {
	var res int
	if num1 > num2 {
		res = num1
	} else {
		res = num2
	}
	return res
}

func add(x, y int) int {
	return x + y
}
