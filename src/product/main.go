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

var db *gorm.DB

// 定义数据模型(结构体定义)
type Product struct {
	ID        int       `gorm:"primary_key" json:"id"`
	Name      string    `gorm:"not null;unique" json:"name"` // 设置字段为非空并唯一
	Price     int       `json:"price"`
	Address   string    `json:"address"`
	Desc      string    `json:"desc"`
	CreatedAt time.Time `json:"createAt"`
	UpdatedAt time.Time `json:"updateAt"`
}

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:709463253@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("数据库连接失败", err.Error())
	} else {
		fmt.Println("数据库已连接!")
	}
	db.AutoMigrate(&Product{})

	// 检查模型`Product`表是否存在
	hasTable := db.HasTable(&Product{})
	fmt.Println(hasTable, "---")
	if !hasTable {
		// 为模型`Product`创建表,CHARSET=utf8设置数据库的字符集为utf8
		db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8").CreateTable(&Product{})
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
		v1.GET("/product", productList)
		v1.POST("/product", productCreate)
		v1.PUT("/product/:id", productUpdate)
		v1.DELETE("/product/:id", prodectDelete)
	}

	router.Run("127.0.0.1:8080")
}

// 获取数据列表
func productList(c *gin.Context) {
	var product []Product
	db.Find(&product)
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"status":  http.StatusOK,
		"data":    product,
	})
}

// 创建列表
func productCreate(c *gin.Context) {
	/*
	 gin还提供了更加高级方法，c.Bind，
	 它会更加content-type自动推断是bind表单还是json的参数
	 json格式application/json或者表单格式x-www-form-urlencoded
	*/
	data := &Product{}
	err := c.Bind(data) // c.Bind 可以获取json格式参数
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Create(data)
	c.JSON(http.StatusOK, gin.H{
		"message": "创建成功",
		"status":  http.StatusOK,
		"data":    data,
	})
}

// 更新列表
func productUpdate(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	fmt.Println(id, "---")

	product := Product{ID: id} // 修改条件，根据ID修改
	// 需要更新的元素
	data := &Product{}
	err := c.Bind(data) // c.Bind 可以获取json格式参数
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&product).Update(data)
	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
		"status":  http.StatusOK,
		"data":    data,
	})
}

// 删除列表
func prodectDelete(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	fmt.Println(id, "--")

	db.Where("id=?", id).Delete(Product{})
	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
		"status":  http.StatusOK,
		"data":    id,
	})
}
