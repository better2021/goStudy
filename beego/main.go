package main

import (
	"fmt"
	"github.com/astaxie/beego"
)

// RESTful Controller 路由
type RESTfulController struct {
	beego.Controller
}

// 正则路由
type RegExpController struct {
	beego.Controller
}

// get方法
func (this * RESTfulController) Get(){
	this.Ctx.WriteString("hello go 哟")
}

// post方法
func (this * RESTfulController) Post() {
	this.Ctx.WriteString("post go method")
}

func (this * RegExpController) Get(){
	this.Ctx.WriteString(fmt.Sprintln("RegExpController"))
	id := this.Ctx.Input.Param(":id")
	this.Ctx.WriteString(fmt.Sprintln("id is:",id))
}

func main(){
	fmt.Println("hello go")
	beego.Router("/RESTful",&RESTfulController{}) // 访问地址 http://127.0.0.1:8081/RESTful
	// 正则路由，从path中提取参数
	beego.Router("/RegExp1/?:id",&RegExpController{})
	beego.Router("/RegExp2/:id([0-9]+)",&RegExpController{})

	// 启动服务
	beego.Run("127.0.0.1:8081")
}