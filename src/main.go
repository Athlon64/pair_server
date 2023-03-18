package main

import (
	_ "pair_server/models"
	_ "pair_server/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

