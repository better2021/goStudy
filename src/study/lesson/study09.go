package lesson

import (
	"fmt"
	"regexp"
	"runtime"
	"time"
)

func Goroutine()  {
	//go say("hello go")
	//say("world")

	//cpuNum := runtime.NumCPU()
	//fmt.Println(cpuNum,"--")

	// running()

	//ch := make(chan string,5)
	// 开启100个协程
	//for i:=0;i<100;i++{
	//	go timeMore(ch)
	//}
	//
	//for{
	//	time.Sleep(time.Second)
	//}

	a := "I am learning Go language"
	re,_ :=  regexp.Compile("[a-z]{2,4}")
	// 查找符合正则的第一个
	one := re.Find([]byte(a))
	fmt.Println("Find:",string(one))

	// 查找符合正则的全部
	two := re.FindAll([]byte(a),-1)
	fmt.Println("Find:",two)
}


func say(s string)  {
	for i:=0;i<5;i++  {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func running()  {
	for i := 1; i <= 10; i++ {
		go func(i int){
			fmt.Println(i)		// 打印一组无规律数字
		}(i)

	}
	time.Sleep(time.Second)
}

// 耗时操作timeMore，现在有100个并发，限制为5个
func timeMore(ch chan string){
	// 执行前先注册，谢步进去就会阻塞
	ch <- "任务"
	fmt.Println("模拟耗时操作")
	time.Sleep(time.Second) // 模拟耗时操作

	// 任务执行完毕，则管道中销毁一个任务
	<-ch
}