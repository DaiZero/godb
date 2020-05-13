package main

import (
	"dzero/godb/router"
	"dzero/godb/util"
)

func main() {
	util.InitDB()
	router := router.InitRouter()
	router.Run()
}
