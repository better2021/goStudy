package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/unknwon/com"
)

// 定义数据模型(结构体)
type Movie struct {
	ID        int       `gorm:"primary_key" json:"id"`
	Name      string    `gorm:"not null;unique" json:"name"`
	Year      string    `json:"year"`
	Desc      string    `json:"desc"`
	CreatedAt time.Time `json:"createAt"`
	UpdatedAt time.Time `json:"updateAt"`
}

var db *gorm.DB

// init 初始化函数会最先执行
func init() {
	var err error
	db, err = gorm.Open("mysql", "root:709463253@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("数据库连接失败", err.Error())
	} else {
		fmt.Println("数据库已连接")
	}

	// 关联数据表
	db.AutoMigrate(&Movie{})

	// 检查模型`Movie`的表是否存在
	hasTable := db.HasTable(&Movie{})
	fmt.Println(hasTable, "--")
	if !hasTable {
		// 为模型`Product`创建表,CHARSET=utf8设置数据库的字符集为utf8
		db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8").CreateTable(&Movie{})
	}
}

func main() {
	/*
		cors.Default() 默认允许所有源跨域
		跨域要写在路由前面，需要先执行
	*/
	router := gin.Default()
	router.Use(cors.Default())

	v1 := router.Group("/api/v1")
	{
		v1.GET("/movie", movieList)
		v1.POST("/movie", movieCreate)
		v1.PUT("/movie/:id", movieUpdate)
		v1.DELETE("/movie/:id", movieDelete)
	}

	router.Run("127.0.0.1:8081")

}

// 获取电影列表
func movieList(c *gin.Context) {
	var movie []Movie
	name := c.Query("name")
	pageNum := com.StrTo(c.Query("pageNum")).MustInt()
	pagesize := 10

	fmt.Println("name", name, pageNum, pagesize)
	/*
		迷糊搜索，name为搜索的条件，根据电影的名称name来搜索
		Offset 其实条数
		Limit	 每页的条数
		Order("id desc") 根据id倒序排序
		总条数 Count(&count)
	*/

	var count int
	db.Model(&movie).Where("name LiKE?", "%"+name+"%").Count(&count)
	db.Offset((pageNum-1)*pagesize).Limit(pagesize).Order("id desc").Where("name LiKE?", "%"+name+"%").Find(&movie)
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"status":  http.StatusOK,
		"data":    movie,
		"page":    pageNum,
		"total":   count, // 总条数
	})
}

// 创建列表
func movieCreate(c *gin.Context) {
	/*
	 gin还提供了更加高级方法，c.Bind，
	 它会更加content-type自动推断是bind表单还是json的参数
	 json格式application/json或者表单格式x-www-form-urlencoded
	*/
	data := &Movie{}
	err := c.Bind(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data, "--")
	db.Create(data)
	c.JSON(http.StatusOK, gin.H{
		"message": "创建成功",
		"status":  http.StatusOK,
		"data":    data,
	})
}

//  更新列表
func movieUpdate(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	fmt.Println(id, "--")

	movie := &Movie{ID: id} // 修改条件，根据ID修改
	// 需要更新的元素
	data := &Movie{}
	err := c.Bind(data) // c.Bind 可以获取json格式参数
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(movie).Update(data)
	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
		"status":  http.StatusOK,
		"data":    data,
	})
}

// 删除列表
func movieDelete(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	fmt.Println(id, "--")

	db.Where("id=?", id).Delete(Movie{})
	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
		"status":  http.StatusOK,
		"data":    id,
	})
}
