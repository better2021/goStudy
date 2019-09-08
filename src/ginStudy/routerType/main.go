package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func main(){
	f,_:=os.Create("gin.log")	// 创建gin.log日志文件
	// gin.DefaultWriter = io.MultiReader(f)
	gin.DefaultErrorWriter = io.MultiWriter(f) // 错误信息写入log日志文件

	r:=gin.Default()
	r.GET("/get", func(c*gin.Context) {
		c.String(200,"get")
	})

	r.POST("/post", func(c *gin.Context) {
		c.String(200,"post")
	})

	r.PUT("/patch", func(c *gin.Context) {
		c.String(200,"patch")
	})

	r.DELETE("/delete", func(c *gin.Context) {
		c.String(200,"delete")
	})

	// 静态资源文件
	r.StaticFS("/static",http.Dir("static"))

	// 模板渲染
	r.LoadHTMLGlob("template/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",gin.H{
			"title":"index.html",
		})
		// panic("err")
		fmt.Println(add(2,3))

		//mapA:=map[string]int{"apple": 5, "lettuce": 7}
		//mapB,_:=json.Marshal(mapA)
		//fmt.Println(mapB,"json")

		resp,err:= http.Get("http://example.com/")
		if err!=nil{
			fmt.Println(err)
			return
		}
		fmt.Println(resp,"resp")
	})




	// 自动化证书配置
	//r.GET("/test", func(c *gin.Context) {
	//	c.String(200,"hello test")
	//})
	//autotls.Run(r,"www.itpp.tk") // www.itpp.tk 是配置证书的服务器域名地址


	r.Run(":8082")

}

// 函数的返回值也可以在函数中预先定义
func add(a int,b int)(c int,d string){
	c = a + b
	return c,"hello"
}