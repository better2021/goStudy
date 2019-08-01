package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// 可以通过"omitempty"参数忽略掉值为空的键
type Admin struct {
	gorm.Model
	ID      int64  `json:"id,omitempty"`
	Name    string `json:"name"`
	Address string `json:"adress"`
}

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:709463253@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("数据库连接失败", err.Error())
	} else {
		fmt.Println("数据库已连接！")
	}
	db.AutoMigrate(&Admin{})
}

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1/admin")
	{
		v1.GET("/list", adminList)
		v1.POST("/create", adminCreate)
		v1.PUT("/update", adminUpdate)
		v1.DELETE("/delete/:ID", adminDelete)
	}

	router.Run("localhost:8088")
}

// 获取数据列表
func adminList(c *gin.Context) {
	var admins []Admin
	db.Find(&admins)
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"status":  http.StatusOK,
		"data":    admins,
	})
}

// 创建列表
func adminCreate(c *gin.Context) {
	admins := Admin{Name: c.PostForm("name"), Address: c.PostForm("address")}
	db.Create(&admins)
	c.JSON(http.StatusOK, gin.H{
		"message": "创建成功",
		"status":  http.StatusOK,
		"data":    admins,
	})
}

// 更新列表
func adminUpdate(c *gin.Context) {
	id, err := strconv.ParseInt(c.PostForm("id"), 10, 0)
	if err != nil {
		panic(err)
	}

	fmt.Println(id, "--")
	admin := Admin{ID: id}
	var opts = Admin{ID: id, Name: c.PostForm("name"), Address: c.PostForm("address")}
	db.Model(&admin).Update(opts)
	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
		"status":  http.StatusOK,
		"data":    admin,
	})
}

// 删除列表
func adminDelete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		panic(err)
	}

	fmt.Println(id, "--")
	admin := Admin{ID: id}
	db.Delete(&admin)
	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
		"status":  http.StatusOK,
	})
}
