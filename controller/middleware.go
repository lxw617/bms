package controller

import (
	"fmt"
	"github.com/kataras/iris/v12"
)

func PreHandler(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("id")
	fmt.Println("[MIDDLEWARE] get book", id)
	//ctx.Values().Set("aaa", "hello")
	ctx.Next()
}
