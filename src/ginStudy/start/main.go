package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type Person struct {
	Name string `form:"name" binding:"required"` // binding:"required" 表示必传
	Address string	`form:"address"`
	//Birthday time.Time `form:"birthday" time_format:"2006-01-02"`
}

func main()  {
	r:=gin.Default()

	r.GET("/ping",func(c*gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"message":"pass GO",
		})
	})

	r.GET("/getParams/:name/:age", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"name":c.Param("name"),
			"age":c.Param("age"),
		})
	})

	// 请求范绑定，只要是/user请求都会返回c.String中的hello go
	r.GET("/user/*action", func(c *gin.Context) {
		c.String(200,"hello go")
	})
	
	r.GET("/test", func(c *gin.Context) {
		name:= c.Query("name")
		age:= c.DefaultQuery("age","26")
		c.JSON(200,gin.H{
			"name":name,
			"age":age,
		})
	})

	r.POST("/test", func(c *gin.Context) {
		bodyBytes,err:=ioutil.ReadAll(c.Request.Body)
		if err!=nil{
			c.String(http.StatusBadRequest,err.Error())
			c.Abort()
		}
		c.String(http.StatusOK,string(bodyBytes))
	})

	r.POST("/te",testing)

	r.Run(":8081")
}


func testing(c *gin.Context)  {
	c.Request.Header.Set("Content-Type","application/json") // 设置请求头为application/json格式
	fmt.Println(c.ClientIP(),"ip")
	fmt.Println(c.ContentType(),"ContentType")
	fmt.Println(c.Request.Method,"method")

	var person Person

	if err:= c.ShouldBind(&person);err!=nil{
		c.String(400,"%c",err.Error())
	}
	fmt.Println(person.Name,"--")
	c.String(200,person.Address)
}
