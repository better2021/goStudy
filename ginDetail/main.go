package main

import(
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

/**
* gin实现商品详情
* 1.搭建静态HTML服务
* 2.实现ajax 服务部分
* 3.实现ajax HTML部分
*/
func main()  {
	fmt.Println("http://127.0.0.1:8088/html")
	router := gin.Default()
	router.Static("/html","./html") // 用于存放静态html文件

	// 路由，设定/detail/:id 返回JSON数据
	router.GET("/detail/:id",detailGet)

	// 服务启动
	router.Run("127.0.0.1:8088")
}

type photo struct {
	Id string
	ImagePath string
}

// 返回三项数据；photo，title，price
func detailGet(c * gin.Context){
	// 假装处理商品id
	id := c.Param("id")
	fmt.Println("id is",id)

	// 假装访问数据库
	buffer :=[]photo {
		photo{
			Id:"pic1",
			ImagePath:"img01.png",
		},
		photo{
			Id:"pic2",
			ImagePath:"img02.png",
		},
		photo{
			Id:"pic3",
			ImagePath:"img03.png",
		},
	}

	// 返回JSON格式的数据
	c.JSON(http.StatusOK,gin.H{
		"photos":buffer,
		"title":"商品的标题",
		"price":100,
	})
}