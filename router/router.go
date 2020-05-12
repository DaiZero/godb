package router

import (
	v1 "dzero/godb/api"
	"dzero/godb/middleware"
	"dzero/godb/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {

	router := gin.Default()
	gin.New()

	router.POST("/api/auth/register", func(ctx *gin.Context) {
		//ctx.JSON(200, gin.H{
		//	"message": "pong",
		//})
		//获取参 数
		name := ctx.PostForm("name")
		telephone := ctx.PostForm("telephone")
		password := ctx.PostForm("password")

		//数据验证
		if len(telephone) != 11 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": http.StatusUnprocessableEntity, "message": "手机号必须为11位"})
			return
		}

		if len(password) < 6 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": http.StatusUnprocessableEntity, "message": "密码不小于6位"})
			return
		}

		if len(name) == 0 {
			name = util.RandName(10)
		}
		ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "注册成功"})
		fmt.Println(name, telephone, password)

	})
	router.POST("/api/auth/login", func(ctx *gin.Context) {
		name := ctx.PostForm("name")
		password := ctx.PostForm("password")
		token, err := util.GenerateToken(name, password)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": ""})
		}
		ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": token, "message": "成功"})
	})

	apiRouter := router.Group("/api/v1")
	apiRouter.Use(middleware.JWTAuth())
	{
		apiRouter.GET("user",v1.GetUserInfo)
	}
	return router
}
