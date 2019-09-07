package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
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

	r.Run(":8081")

}
