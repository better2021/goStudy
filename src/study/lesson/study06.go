package lesson

import (
	"fmt"
	"strconv"
	"strings"
)

func Str()  {
	// func Contains(s, substr string) bool
	// 字符串s中是否包含substr，返回bool值
	fmt.Println(strings.Contains("seafood","foo"))
	fmt.Println(strings.Contains("seafood","bar"))
	fmt.Println(strings.Contains("seafood",""))
	fmt.Println(strings.Contains("", ""))

	// 字符串链接，把slice a通过sep链接起来
	s := []string{"foo","bar","baz"}
	fmt.Println(strings.Join(s,","))

	// 在字符串s中查找sep所在的位置，返回位置值，找不到返回-1
	fmt.Println(strings.Index("chicken","ken")) // 4
	fmt.Println(strings.Index("chicken","dir")) // -1

	// 重复s字符串count次，最后返回重复的字符串
	fmt.Println(strings.Repeat("na",3))

	//func Split(s, sep string) []string
	//把s字符串按照sep分割，返回slice
	fmt.Println(strings.Split("a,b,c",","))
	fmt.Println(strings.Split("f-e-q","-"))
	fmt.Println(strings.Split("xyz",""))

	// func Trim(s string, cutset string) string
	// //在s字符串的头部和尾部去除cutset指定的字符串
	fmt.Println(strings.Trim( "!! hello !","! "))

	//func Fields(s string) []string
	//去除s字符串的空格符，并且按照空格分割返回slice
	fmt.Println(strings.Fields("  foo var baz "))

	// Append 系列函数将整数等转换为字符串后，添加到现有的字节数组中。
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '单')
	fmt.Println(string(str))

	// Format 系列函数把其他类型的转换为字符串
	//a := strconv.FormatBool(false)
	//b := strconv.FormatFloat(123.23, 'g', 12, 64)
	//c := strconv.FormatInt(1234, 10)
	//d := strconv.FormatUint(12345, 10)
	//e := strconv.Itoa(1023)
	//fmt.Println(a, b, c, d, e)

	// Parse 系列函数把字符串转换为其他类型
	a, err := strconv.ParseBool("false")
	checkError(err)
	b, err := strconv.ParseFloat("123.23", 64)
	checkError(err)
	c, err := strconv.ParseInt("1234", 10, 64)
	checkError(err)
	d, err := strconv.ParseUint("12345", 10, 64)
	checkError(err)
	e, err := strconv.Atoi("1023")
	checkError(err)
	fmt.Println(a, b, c, d, e)
}

func checkError(e error){
	if e != nil{
		fmt.Println(e)
	}
}