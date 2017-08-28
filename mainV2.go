package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"college_jouranlv2/Control"
)

func main(){
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	v2 := router.Group("/v2")
	{
		//注册 同时更新权限表
	v2.POST("/user/",Control.Postuser)
	//	登录
	v2.POST("/login/", Control.Getlogin)
	//改签名
	v2.PUT("/user/:id", Control.Putuse)
	//	更新用户资料
	v2.PUT("/user/", Control.Putuser)
	//改密码
	v2.PUT("/pwd/:old/:new/:id", Control.Putpwd)
	//	发布信息
	v2.POST("/info/", Control.Postinfo)
	//	获取信息
	v2.GET("/info/:key", Control.Getinfo)
	//	获取用户信息
	v2.GET("/user/", Control.Getuser)
	//	反馈建议
	v2.POST("/suggest/", Control.Postsuggest)
	//点赞
	v2.POST("/praise/", Control.Postpraise)
	v2.GET("/update/", Control.Getupdate)
	}

	router.Run(":8080")

	Control.CloseEngine()
	fmt.Println("end of service !")

}
