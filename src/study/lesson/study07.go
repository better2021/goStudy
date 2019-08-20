package lesson

import "fmt"

func Fn(){
	// 匿名函数，获取最大值，最小值
	x,y := func(i,j int) (max,min int){
		if i >j {
			max = i
			min = j
		}else {
			max = j
			min = i
		}
		return
	}(15,20)
	fmt.Println("最大值：",x)
	fmt.Println("最小值：",y)
}

// 定义一个函数类型
type MyMath func(int,int) int

