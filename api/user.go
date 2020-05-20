package v1

import (
	"dzero/godb/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 获取用户信息
// @Description GetUserInfo
// @Tags User
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.User
// @Security ApiKeyAuth
// @Router /api/v1/user [post]
func GetUserInfo(c *gin.Context) {
	fmt.Println("GetUserInfo")
	userReturn := service.GetUser()
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": userReturn, "message": "成功"})
}
