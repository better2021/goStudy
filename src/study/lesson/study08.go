package lesson

import (
	"errors"
	"fmt"
	"reflect"
)

func DeferHandle()  {
	// 当执行到defer语句时，暂不执行，会将defer后的语句压入到独立的栈中,当函数执行完毕后，再从该栈按照先入后出的方式出栈执行
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	fmt.Println("main")

	num:=0
	defer fmt.Println(num)
	num = 3
	fmt.Println("num",num)

	var err = errors.New("发生了错误")
	fmt.Println(err)

	s := new(GameService)
	s.Start()
	s.Log("hello")

	// 创建动物名字到实例的映射
	animals := map[string]interface{}{
		"bird":new(bird),
		"pig":new(pig),
	}

	for name,obj := range animals{
		// 断言；判断对选哪个是否飞行动物，行走动物
		f,isFlyer := obj.(Flyer)
		w,isWalker := obj.(Walker)

		fmt.Println("name is",name)

		if isFlyer{
			f.Fly()
		}
		if isWalker{
			w.Walk()
		}
	}

	// 可以实现将接口转换为普通的指针类型，例如将Walker接口转换为*pig类型
	p1 := new(pig)
	var a Walker = p1
	p2 := a.(*pig)
	fmt.Println(p1,p2,"--")

	// 反射操作数据
	var ad int
	fmt.Println(reflect.ValueOf(ad),"---")
	fmt.Println(reflect.TypeOf(ad))
	fmt.Println(reflect.TypeOf(ad).Name())
	fmt.Println(reflect.TypeOf(ad).Kind())

	var number int64 = 100
	rValue := reflect.ValueOf(number)
	fmt.Println(rValue)
	// 假发运算
	fmt.Println(number + rValue.Int())

}

type Service interface {
	Start()
	Log(string)
}

// 日志器
type Logger struct {

}

// 日志输出方法
func (g *Logger) Log(s string){
	fmt.Println("日志",s)
}

// 游戏服务
type GameService struct {
	Logger
}

// 实现游戏服务的Start方法
func (g *GameService) Start(){
	fmt.Println("游戏服务启动")
}

/*
鸟和猪都具有不同的特性，鸟可以飞，猪不能飞，但是二者都可以走，
如果使用结构体实现鸟和猪，让他具备各自的特性Fly()和Walk()方法就让鸟和猪各自实现了飞行动物接口Flyer和行走动物接口Walker
t, ok := i.(T)		// 安全写法：如果接口未实现接口，将会把ok掷为false，t掷为T类型的0值
*/

// 飞行动物接口
type Flyer interface {
	Fly()
}

// 行走动物接口
type Walker interface {
	Walk()
}

// 鸟类
type bird struct {

}

// 鸟类实现飞行动物接口
func (b *bird) Fly() {
	fmt.Println("bird fly")
}

// 鸟类实现行走动物接口
func (b *bird) Walk() {
	fmt.Println("bird walk")
}

// 猪类
type pig struct {

}

// 猪类实现行走动物接口，没有实现飞行接口
func (p *pig) Walk() {
	fmt.Println("pig walk")
}


