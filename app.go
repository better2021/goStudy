package main // 声明文件的package
import "fmt" // 通过import 关键字导入其他非main包

func main(){  // 入口函数, 无参数无返回值
	fmt.Println("hello go!")
	var b,c int = 1,2
	f := "good"
	var d = sum(b,c)
	var arr = [3]int{3,5,9}
	fmt.Println(d,f,arr[1]) // 数组 [3]int{3,5,9} [3]数组的长度 in数组中值的类型 {3,5,9}数组具体的值
	fmt.Println(max(3,6))
	fmt.Println("变量的地址：",&c)
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

func sum(a,b int) int {
	return a + b
}
