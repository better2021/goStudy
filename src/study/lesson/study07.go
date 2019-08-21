package lesson

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

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

	// 字符串遍历，处理中文问题
	r := []rune("hello武汉")
	fmt.Println(r[1])

	// 字符串转整数
	n,_ := strconv.Atoi("123")
	fmt.Println(n)

	// 整数转字符串
	str := strconv.Itoa(123456)
	fmt.Println(str)

	// 字符串转[]byte
	var bytes = []byte("golang")
	fmt.Println(bytes)

	// []byte转字符串
	str01 := string(bytes)
	fmt.Println(str01)

	// 查找子字符串是否在指定的字符串中
	var ss = strings.Contains("saefoond","foo")
	fmt.Println(ss)

	// 统计一个字符中有几个指定的字串
	 aa := strings.Count("banana","a")
	fmt.Println(aa)

	 // 不区分大小写的字符串比较(==是区分字母大小写的)
	 fmt.Println(strings.EqualFold("abc","AbC"))

	 // 返回字串在字符串中第一次出现的index值，如果没有返回-1
	 frist :=  strings.Index("NL_abc","b")
	fmt.Println(frist)

	 // 返回子字符最后一次的inex值，如果没有返回-1
	 last :=strings.LastIndex("go golang","o")
	 fmt.Println(last)

	 // 指定字符串替换为另一个字符串
	 re := strings.Replace("go go hello","go","go 语言",-1)
	 fmt.Println(re)

	 // 将字符串的字母进行大小写的转换
	 low := strings.ToLower("Go")
	 upp := strings.ToUpper("Go")
	 fmt.Println(low,upp)

	 // 将字符串的左右两边的空格去掉
	 trim := strings.TrimSpace(" go to shop ")
	 fmt.Println(trim)

	 // 时间函数
	 now := time.Now()
	 fmt.Println(now)
	 fmt.Println(now.Format("2018-10-10"))
	 fmt.Println(now.Unix()) // 时间戳(秒数)
}



