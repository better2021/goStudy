package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	/*
		图片上传
		c.Request.FormFile("file") 获取上传的文件
	*/

	router.POST("/upload", func(c *gin.Context) {
		file, header, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "文件上传失败",
			})
			return
		}
		//获取文件名
		filename := header.Filename
		fmt.Println(filename, "---")

		content, err := ioutil.ReadAll(file)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "文件读取失败"})
			return
		}
		fmt.Println(string(content))
		filepath := "http://127.0.0.1:8000/upload/tmp/" + filename
		//以json格式返回文件存放路径
		c.JSON(http.StatusOK, gin.H{
			"message":  "上传成功",
			"filepath": filepath,
		})
	})

	router.Run("127.0.0.1:8000")
}
