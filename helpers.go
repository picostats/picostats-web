package main

import (
	"gopkg.in/kataras/iris.v6"
)

type PageData struct {
	Conf *Config
}

func newPageData(ctx *iris.Context) *PageData {
	pd := &PageData{}
	pd.Conf = conf
	return pd
}
