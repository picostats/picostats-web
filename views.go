package main

import (
	"fmt"
	"log"

	"gopkg.in/kataras/iris.v6"
)

func pageView(ctx *iris.Context) {
	titlePrefixes := map[string]string{
		"home":      "",
		"demo":      "Demo | ",
		"downloads": "Downloads | ",
		"docs":      "Documentation | ",
	}

	pd := newPageData(ctx)
	page := ctx.Param("page")
	if len(page) == 0 {
		page = "home"
		pd.Page = ""
	} else {
		pd.Page = "/" + page
	}

	title, ok := titlePrefixes[page]
	var status int
	if ok {
		pd.TitlePrefix = title
		status = iris.StatusOK
	} else {
		pd.TitlePrefix = "Not Found"
		page = "404"
		status = iris.StatusNotFound
	}

	template := fmt.Sprintf("%s.html", page)

	ctx.RenderWithStatus(status, template, pd)
}

func pagePricingView(ctx *iris.Context) {
	pd := newPageData(ctx)
	pd.TitlePrefix = "Pricing | "
	pd.UserId = ctx.Param("userId")
	ctx.Render("pricing.html", pd)
}

func paidView(ctx *iris.Context) {
	pf := &PaymentForm{}
	err := ctx.ReadForm(pf)
	if err != nil {
		log.Printf("[views.go] Error reading PaymentForm: %s", err)
	}

	if getMD5Hash(pf.SecurityData+conf.FSPrivateKey) == pf.SecurityHash {
		user := &User{}
		db.First(user, pf.UserID)
		user.MaxWebsites = 0
		db.Save(user)
	}

	ctx.SetStatusCode(iris.StatusOK)
}

func paidCancelView(ctx *iris.Context) {
	pf := &PaymentForm{}
	err := ctx.ReadForm(pf)
	if err != nil {
		log.Printf("[views.go] Error reading PaymentForm: %s", err)
	}

	if getMD5Hash(pf.SecurityData+conf.FSPrivateKey2) == pf.SecurityHash {
		user := &User{}
		db.First(user, pf.UserID)
		user.MaxWebsites = 1
		db.Save(user)
	}

	ctx.SetStatusCode(iris.StatusOK)
}
