package main

import (
	"project/Schedule/src/common"
	"project/Schedule/src/db"
	"project/Schedule/src/server"
)

func main() {
	common.Init()
	db.Init()
	server.Init()
}

