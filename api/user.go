package v1

import (
	"dzero/godb/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserInfo(c *gin.Context)  {
    fmt.Println("GetUserInfo")
	userReturn := service.GetUser()
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": userReturn, "message": "成功"})
}

