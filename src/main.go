package main

import (
	"Schedule/src/common"
	"Schedule/src/db"
	"Schedule/src/server"
)

func main() {
	//env road
	common.Init()

	//db connect info road
	db.Init()

	//golang run
	server.Init()
}

