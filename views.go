package main

import (
	"fmt"

	"gopkg.in/kataras/iris.v6"
)

func pageView(ctx *iris.Context) {
	titles := map[string]string{
		"home":      "Home",
		"demo":      "Demo",
		"downloads": "Downloads",
	}

	pd := newPageData(ctx)
	page := ctx.Param("page")
	if len(page) == 0 {
		page = "home"
	}

	title, ok := titles[page]
	var status int
	if ok {
		pd.Title = title
		status = iris.StatusOK
	} else {
		pd.Title = "Not Found"
		page = "404"
		status = iris.StatusNotFound
	}

	template := fmt.Sprintf("%s.html", page)

	ctx.RenderWithStatus(status, template, pd)
}
