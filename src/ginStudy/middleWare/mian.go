package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 中间件
 func IPAuthMiddleWare() gin.HandlerFunc{
 	return func(c *gin.Context) {
		ipList:=[]string{ // IP白名单
			"127.0.0.1","127.0.0.2",
		}
		flag := false
		clientIp:=c.ClientIP()
		for _,host:=range ipList {
			if clientIp==host{
				flag=true
				break
			}
		}
		if !flag{
			c.String(401,"%s,not in ipList",clientIp)
			panic("not in ipList")
		}
	}
 }

func main(){
	f,_:=os.Create("gin.log")	// 创建gin.log日志文件
	// gin.DefaultWriter = io.MultiReader(f)
	gin.DefaultErrorWriter = io.MultiWriter(f) // 错误信息写入log日志文件

	r:=gin.New()
	r.Use(gin.Logger(),gin.Recovery())
	r.Use(IPAuthMiddleWare())

	r.GET("/test/:age", func(c *gin.Context) {
		name:=c.DefaultQuery("name","xiaoming")

		c.String(http.StatusOK,"%s",name,c.Param("age"))

	})

	/*
	优雅关停服务器
	*/
	r.GET("/close", func(c *gin.Context) {
		time.Sleep(10*time.Second)
		c.String(200,"hello test")
	})

	srv:=&http.Server{
		Addr:":8081",
		Handler:r,
	}

	go func() {
		if err:=srv.ListenAndServe();err!=nil&&err!=http.ErrServerClosed{
			log.Fatalf("listen: %s\n",err)
		}
	}()

	quit:=make(chan os.Signal)
	signal.Notify(quit,syscall.SIGINT,syscall.SIGTERM)
	<-quit
	log.Println("shutdown server ...")

	ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	if err:=srv.Shutdown(ctx);err!=nil{
		log.Fatal("server shutDown",err)
	}
	log.Println("sever exiting")

	r.Run(":8081")
}
