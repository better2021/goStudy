package main

import(
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

// GET函数
func hello(c * gin.Context){
	c.String(http.StatusOK, "hello gin Get")
}

// POST函数
func helloPost(c * gin.Context){
	c.String(http.StatusOK,"gin Post")
}

// 提取请求中的参数
func fetchId(c * gin.Context){
	id := c.Param("id")
	c.String(http.StatusOK,fmt.Sprintln("id is",id))
}

// 组路由
func action1(c*gin.Context){
	c.String(http.StatusOK,"action1")
}

func action2(c*gin.Context){
	c.String(http.StatusOK,"action2")
}

func main(){
	fmt.Println("hello gin")
	router := gin.Default()

	// RESTful 路由
	router.GET("/getGin",hello)
	router.POST("/postGin",helloPost)

	// 提取请求中的参数
	router.GET("/param/:id",fetchId)

	// 组路由
	group1 := router.Group("/go1")
	{
		group1.GET("/action1",action1)
		group1.GET("/action2",action2)
	}

	// 启动服务
	router.Run("127.0.0.1:8082")
}
