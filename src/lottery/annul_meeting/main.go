package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/1/17 下午2:20
 */

type lotteryController struct {
	Ctx *iris.Context
}

func NewApp() *iris.Application {
	var app *iris.Application
	app = iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryController{})
	return app
}

func main() {
	var app *iris.Application
	app = NewApp()

	app.Run(iris.Addr(":8080"))
}

func (ctrl *lotteryController) Get() string {
	return "Hello iris!"
}
