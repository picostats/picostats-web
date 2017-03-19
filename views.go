package main

import (
	"fmt"
	"log"

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

func pagePricingView(ctx *iris.Context) {
	pd := newPageData(ctx)
	pd.Title = "Pricing"
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
