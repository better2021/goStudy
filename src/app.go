package main // 声明文件的package
// 通过import 关键字导入其他非main包
import(
	"fmt"
	"calc"
) 

func main(){  // 入口函数, 无参数无返回值
	fmt.Println("hello go!")
	var b,c int = 1,2
	f := "good"
	var d = add(b,c)
	// var e = calc.Sub(c,b)
	var arr = [3]int{3,5,9}
	fmt.Println(d,f,arr[1]) // 数组 [3]int{3,5,9} [3]数组的长度 in数组中值的类型 {3,5,9}数组具体的值
	fmt.Println(max(3,6))
	fmt.Println("变量的地址：",&c)
	fmt.Println(calc.Add(b,c))
}

func max(num1,num2 int) int {
	var res int
	if(num1>num2){
		res = num1
	}else{
		res = num2
	}
	return res
}

func add(x,y int) int {
	return x + y
}
