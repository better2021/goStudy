package lesson

import (
	"fmt"
	"os"
)

func Ostest()  {
	// dir()

	//file()

	//read()

	 //dele()
}

// 文件夹操作
func dir()  {
	// 创建文件名为name的目录，权限设置0777 （0777为最高权限）
	os.Mkdir("test",0777)
	// 根据path创建多级子目录，例如：test/godir/te
	os.MkdirAll("test/godir/te",0777)
	// 删除名为test的目录，当目录下有文件或者其他目录是会出错
	err := os.Remove("test")
	if err !=nil{
		fmt.Println(err)
	}
	// 根据path删除多级子目录，如果path是单个名称，那么该目录下的子目录全部删除
	os.RemoveAll("test")
}

// 文件操作
func file(){
	txtFile := "hello.txt"
	// 根据提供的文件名创建新的文件，返回一个文件对象，默认权限是0666的文件，返回的文件对象是可读写的。
	file,err:= os.Create(txtFile)
	if err !=nil{
		fmt.Println(err)
		return
	}
	defer file.Close()
	for i:=0;i<10;i++{
		// 写入string信息到文件
		file.WriteString("test hello go\r\n")
		// 写入byte类型的信息到文件
		file.Write([]byte("write byte test\r\n"))
	}
}

// 读文件
func read(){
	readFile := "hello.txt"
	// 读取hello.txt文件的内容
	fl,err := os.Open(readFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fl.Close()
	buf := make([]byte,1024)
	n,_:=fl.Read(buf)
	os.Stdout.Write(buf[:n])
}

// 删除文件
func dele(){
	err:= os.Remove("hello.txt")
	if err !=nil{
		fmt.Println(err)
		return
	}
}