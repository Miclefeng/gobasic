package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
	"math/rand"
	"os"
	"time"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/1/30 下午3:52
 */

var logger *log.Logger

// 存储所有红包
var packageList map[uint32][]uint64

type lotteryController struct {
	Ctx iris.Context
}

func initLog() {
	var f *os.File
	f, _ = os.Create("/var/logs/lottery_demo.log")
	logger = log.New(f, "", log.Ldate|log.Lmicroseconds)
}

func newApp() (app *iris.Application) {
	app = iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryController{})
	return
}

func main() {
	var app *iris.Application
	app = newApp()
	packageList = make(map[uint32][]uint64)
	app.Run(iris.Addr(":8080"))
}

// 发红包
// GET http://localhost:8080/set?uid=1&money=100&num=100
func (c *lotteryController) GetSet() string {
	var (
		uid        int
		errUid     error
		money      float64
		errMoney   error
		num        int
		errNum     error
		moneyTotal uint64
		leftMoney  uint64
		leftNum    uint64
		random     *rand.Rand
		randMax    float64
		list       []uint64
		i          uint64
		randMoney  int64
		curMoney   uint64
		id         uint32
	)
	uid, errUid = c.Ctx.URLParamInt("uid")
	money, errMoney = c.Ctx.URLParamFloat64("money")
	num, errNum = c.Ctx.URLParamInt("num")
	if (errUid != nil || errMoney != nil || errNum != nil) {
		return fmt.Sprintf("参数格式错误：errUid=%s,errMoney=%s,errNum=%s", errUid, errMoney, errNum)
	}
	// 金额以分记数
	moneyTotal = uint64(money * 100)
	if uid < 1 || moneyTotal < 1 || num < 1 {
		return fmt.Sprintf("参数值错误：uid=%d,money=%d,num=%d", uid, money, num)
	}
	leftMoney = moneyTotal
	leftNum = uint64(num)
	// 分配随机数
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
	randMax = 0.25             // 最大中奖比例
	list = make([]uint64, num) // 存储已分配的红包金额

	for leftNum > 0 {
		// 只剩一份，分配剩余全部金额
		if 1 == leftNum {
			list[num-1] = leftMoney
			break
		}
		// 每人最少获取1分
		if leftMoney == leftNum {
			for i = 0; i < leftNum; i++ {
				list[uint64(num)-leftNum] = 1
			}
			break
		}

		// 每次对剩余金额的1%-25%随机，最小1%，最大就是剩余金额25%（需要给剩余的名额留下1分钱的生存空间）
		randMoney = int64(float64(leftMoney-leftNum) * randMax)
		curMoney = uint64(random.Int63n(randMoney))
		if curMoney < 1 {
			curMoney = 1
		}
		list[uint64(num)-leftNum] = curMoney
		leftMoney -= curMoney
		leftNum--
	}
	// 生成
	id = random.Uint32()
	packageList[id] = list
	// 返回抢红包的URL
	return fmt.Sprintf("/get?id=%d&uid=%d&num=%d\n", id, uid, num)
}

// 抢红包
// GET http://localhost:8080/get?id=1&uid=1
func (c *lotteryController) GetGet() string {
	var (
		id     int
		errId  error
		uid    int
		errUid error
		list   []uint64
		ok     bool
		random *rand.Rand
		i      uint64
		money  uint64
	)
	id, errId = c.Ctx.URLParamInt("id")
	uid, errUid = c.Ctx.URLParamInt("uid")
	if errId != nil || errUid != nil {
		return fmt.Sprintf("参数格式错误：errId=%s,errUid=%s", errId, errUid)
	}
	if id < 1 || uid < 1 {
		return fmt.Sprintf("参数值错误：uid=%d,id=%d", uid, id)
	}
	// 判断红包是否存在
	list, ok = packageList[uint32(id)]
	if !ok || len(list) < 1 {
		return fmt.Sprintf("红包不存在,id=%d\n", id)
	}
	// 分配随机数
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
	// 分配一个红包
	i = uint64(random.Intn(len(list)))
	money = list[i]

	if len(list) > 1 {
		if int(i) == len(list)-1 {
			packageList[uint32(id)] = packageList[uint32(id)][:i]
		} else if i == 0 {
			packageList[uint32(id)] = packageList[uint32(id)][1:]
		} else {
			packageList[uint32(id)] = append(packageList[uint32(id)][:i], packageList[uint32(id)][i+1:]...)
		}
	} else {
		delete(packageList, uint32(id))
	}
	return fmt.Sprintf("恭喜你抢到一个红包，金额为:%d\n", money)
}
