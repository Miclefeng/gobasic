package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"math/rand"
	"strings"
	"time"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/1/17 下午2:20
 */

type lotteryController struct {
	Ctx iris.Context
}

var userList []string

func NewApp() *iris.Application {
	var app *iris.Application
	app = iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryController{})
	return app
}

func main() {
	var app *iris.Application
	app = NewApp()
	userList = make([]string, 0)
	app.Run(iris.Addr(":8080"))
}

func (c *lotteryController) Get() string {
	var count int
	count = len(userList)
	return fmt.Sprintf("当前参与抽奖的人数：%d\n", count)
}

func (c *lotteryController) PostImport() string {
	var (
		strUsers string
		users    []string
		count1   int
		user     string
		count2   int
	)
	strUsers = c.Ctx.FormValue("users")
	users = strings.Split(strUsers, ",")
	count1 = len(userList)
	for _, user = range users {
		user = strings.TrimSpace(user)
		if len(user) > 0 {
			userList = append(userList, user)
		}
	}
	count2 = len(userList)
	return fmt.Sprintf("当前参与抽奖的人数：%d, 成功导入用户数：%d\n", count2, (count2 - count1))
}

func (c *lotteryController) GetLucky() string {
	var (
		count int
		seed  int
		idx   int
		user  string
	)

	count = len(userList)
	if count > 1 {
		seed = time.Now().Nanosecond()
		idx = rand.New(rand.NewSource(int64(seed))).Intn(count)
		user = userList[idx]
		userList = append(userList[0:idx], userList[idx+1:]...)
		return fmt.Sprintf("当前中奖用户：%s，剩余用户数：%d\n", user, count-1)
	} else if 1 == count {
		user = userList[0]
		userList = userList[:0]
		return fmt.Sprintf("当前中奖用户：%s，剩余用户数：%d\n", user, count-1)
	} else {
		return fmt.Sprintf("已没有参数用户\n")
	}
}
