package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

/*
结构体到数据表的名称映射规则为结构体名称的复数
结构体的字段到数据表字段的默认映射规则是用下划线分隔每个大写字母开头的单词
例如：ID 对应的数据库字段是id，Name 对应的数据库字段是name
`json:"address"` 设置返回字段为address 小写
*/

type Person struct {
	gorm.Model
	Id      int64  `gorm:"primary_key;AUTO_INCREMENT" json:"id"` // 字段`ID`为默认主键
	Name    string `json:"name"`
	Email   string `gorm:"unique;not null" json:"email"` // 设置会员号（member number）唯一并且不为空
	Age     int    `json:"age"`
	Address string `json:"address"`
}

/*
@param root => 用户数据库的账号
@param 709463253 => 用户数据库的密码
@param test => 用户数据库表名
为了处理time.Time，您需要包括parseTime作为参数
*/
func init() {
	var err error
	db, err = gorm.Open("mysql", "root:709463253@/test?charset=utf8&parseTime=True&loc=Local") // 连接数据库
	if err != nil {
		fmt.Println("数据库连接失败", err.Error())
	} else {
		fmt.Println("数据库已连接！")
	}
	// defer db.Close()
	db.AutoMigrate(&Person{})
}

/*RESTful 增删改查接口 start*/
// 获取列表接口
func userList(c *gin.Context) {
	var persons []Person
	db.Find(&persons)
	fmt.Println(persons)
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"status":  http.StatusOK,
		"data":    persons,
	})
}

// 创建接口
func createUser(c *gin.Context) {
	person := &Person{Name: c.PostForm("name"), Email: c.PostForm("email"), Age: 18, Address: c.PostForm("address")}
	db.Create(person) // 创建一条数据
	db.Save(person)   // 保存在数据库中
	c.JSON(http.StatusOK, gin.H{
		"message": "创建成功",
		"status":  http.StatusOK,
		"data":    person,
	})
}

// 更新接口
func updateUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 0) // strconv.ParseInt方法用于类型转换，ParseInt转换为整形
	if err != nil {
		panic(err)
	}

	fmt.Println(id, "---")
	person := &Person{Id: id}
	db.Model(person).Updates(Person{Name: c.PostForm("name"), Email: c.PostForm("email"), Age: 18, Address: c.PostForm("address")})
	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
		"status":  http.StatusOK,
		"data":    person,
	})
}

// 删除接口
func deleteUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		panic(err)
	}

	fmt.Println(id, "---")
	person := &Person{Id: id}
	db.Delete(person)
	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
		"status":  http.StatusOK,
	})
}

/*RESTful 增删改查接口 end*/

func main() {
	fmt.Println("hello gin")
	router := gin.Default()

	/*RESTful 增删改查接口 start*/
	v1 := router.Group("/api/v1/users")
	{
		v1.GET("/", userList)
		v1.POST("/", createUser)
		v1.PUT("/:id", updateUser)
		v1.DELETE("/:id", deleteUser)
	}
	/*RESTful 增删改查接口 end*/

	router.GET("/", homePage)

	// RESTful 路由
	router.GET("/getGin", hello) // 访问地址：http://127.0.0.1:8082/getGin
	router.POST("/postGin", helloPost)

	// 提取请求中的参数
	router.GET("/param/:id", fetchid)

	// 组路由
	group1 := router.Group("/go1")
	{
		group1.GET("/action1", action1)
		group1.GET("/action2", action2)
	}

	// 启动服务
	router.Run("127.0.0.1:8082")
}

// GET函数
func hello(c *gin.Context) {
	c.String(http.StatusOK, "hello gin Get")
}

// POST函数
func helloPost(c *gin.Context) {
	fmt.Println(c.Request.Body, "---")
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"message": string(value),
	})
}

// 提取请求中的参数
func fetchid(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, fmt.Sprintln("id is", id))
}

// 组路由
func action1(c *gin.Context) {
	c.String(http.StatusOK, "action1")
}

func action2(c *gin.Context) {
	c.String(http.StatusOK, "action2")
}

func homePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello,let is go", // 后面的逗号，是必须要带的
	})
}
