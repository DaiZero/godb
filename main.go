package main

import (
	"dzero/godb/router"
	"dzero/godb/util"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)



func main()  {
	util.InitDB()
	router:=router.InitRouter()
	router.Run()
}


