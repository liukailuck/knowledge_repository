package dbhelp

import (
	"fmt"
	"github.com/astaxie/beego/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"sync"
	"time"
)

var (
	connStr     string
	iniConfiger config.Configer
	err         error
	engine      *xorm.Engine
	lock        sync.Mutex
)

func init() {
	if iniConfiger, err = config.NewConfig("ini", "../conf/app.ini"); err != nil {
		panic(err)
	}
	connStr = iniConfiger.String("mysql::connStr")
}

func NewSingleDbEngine() *xorm.Engine {
	if engine != nil {
		return engine
	}
	lock.Lock()
	defer lock.Unlock()

	if engine != nil {
		return engine
	}

	if engine, err = xorm.NewEngine("mysql", connStr); err != nil {
		fmt.Println("mysql connect fail")
		panic(err)
	}

	if err = engine.Ping(); err != nil {
		fmt.Println("ping timeout")
		panic(err)
	}
	engine.ShowSQL(true)
	engine.SetMapper(core.SameMapper{})
	engine.Logger().SetLevel(core.LOG_DEBUG)
	engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")

	return engine
}
