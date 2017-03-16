package main

import (
	"gopkg.in/kataras/iris.v6"
)

func homeView(ctx *iris.Context) {
	pd := newPageData(ctx)
	ctx.Render("home.html", pd)
}
