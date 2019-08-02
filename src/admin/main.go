package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

/*
基本模型定义gorm.Model，包括字段ID，CreatedAt，UpdatedAt，DeletedAt，
你可以将它嵌入你的模型，或者只写你想要的字段
可以通过"omitempty"参数忽略掉值为空的键
*/

type Admin struct {
	ID        int64     `json:"id,omitempty"`
	Name      string    `json:"name"`
	Address   string    `json:"adress"`
	CreatedAt time.Time `json:"createAt"`
	UpdatedAt time.Time `json:"updateAt"`
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
	router.Use(cors.Default()) // cors.Default() 默认允许所有源跨域

	v1 := router.Group("/api/v1/admin")
	{
		v1.GET("/list", adminList)
		v1.POST("/create", adminCreate)
		v1.PUT("/update/:id", adminUpdate)
		v1.DELETE("/delete/:id", adminDelete)
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
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		panic(err)
	}

	fmt.Println(id, "--")
	// opts 为要修改的元素
	admin := Admin{ID: id} // 修改的条件，根据ID修改
	var opts = Admin{Name: c.PostForm("name"), Address: c.PostForm("address")}
	db.Model(&admin).Update(opts)
	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
		"status":  http.StatusOK,
		"data":    opts,
	})
}

// 删除列表
func adminDelete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		panic(err)
	}

	// Where("id=?", id) 是删除的条件
	fmt.Println(id, "--")
	db.Where("id=?", id).Delete(Admin{})
	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
		"status":  http.StatusOK,
	})
}
