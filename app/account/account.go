/*
 * @Author: mknight(tianyh)
 * @Mail: 824338670@qq.com
 * @Date: 2022-09-16 14:18:13
 * @LastEditTime: 2022-09-30 09:59:10
 * @FilePath: \trcell\app\account\account.go
 */
package account

import (
	"time"
	"trcell/app/account/accrouter"
	"trcell/pkg/appconfig"
	"trcell/pkg/webserve"

	"github.com/bluele/gcache"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type AccountApp struct {
	webServe  *webserve.WebServe
	stopWeb   chan bool
	accountDB *gorm.DB
	appConfig *appconfig.AccountConfig
	dataCache gcache.Cache
}

func NewAccount(appCfg *appconfig.AccountConfig) *AccountApp {
	s := &AccountApp{
		stopWeb:   make(chan bool),
		webServe:  webserve.NewWebServe(),
		appConfig: appCfg,
		dataCache: gcache.New(1024).LRU().Build(),
	}
	s.SetupRouters()
	return s
}

func (acc *AccountApp) GetDB() *gorm.DB {
	return acc.accountDB
}

// func (acc *AccountApp) OpenLocalDB(dbFile string) bool {
// 	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
// 	if err != nil {
// 		panic("open dbfile faile")
// 	}
// 	acc.accountDB = db
// 	return true
// }
func (acc *AccountApp) OpenMysqlDB(connStr string) bool {
	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})
	if err != nil {
		panic("open dbfile faile")
	}
	sqlDB, _ := db.DB()
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(10)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
	acc.accountDB = db
	return true
}
func (acc *AccountApp) GetRouter() *gin.Engine {
	return acc.webServe.GetRouter()
}

func (acc *AccountApp) Run(webAddr string, releaseMode int32) {
	acc.webServe.Run(webAddr, acc.stopWeb, releaseMode)
}

func (acc *AccountApp) Stop() {
	acc.stopWeb <- true
}

func (acc *AccountApp) SetupRouters() {
	accrouter.SetupAccountApi(acc)
}

func (acc *AccountApp) GetAccountDB() *gorm.DB {
	return acc.accountDB
}

func (acc *AccountApp) GetCache() gcache.Cache {
	return acc.dataCache
}
