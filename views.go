package main

import (
	"fmt"

	"gopkg.in/kataras/iris.v6"
)

func pageView(ctx *iris.Context) {
	pd := newPageData(ctx)
	page := ctx.Param("page")
	fmt.Println(page)
	ctx.Render("home.html", pd)
}
