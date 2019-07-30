package calc

/*
变量的可见性规则
大写字母开头的变量是可导出的，也就是其他包可以读取的，是公用的变量
小写字母开头的变量就是不可导出的，是私有变量
例如下面的Sub和Hello以大写字母开头，所以是共有变量可导出被其他包调用
*/

func Sub(a, b int) int {
	return a - b
}

func Hello(msg string) string {
	return "hello" + msg
}
