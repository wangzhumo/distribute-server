package mysql

import (
	"fmt"
	"log"
	"sync"

	"com.wangzhumo.distribute/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 实例
var instance *gorm.DB

// 为了避免并发造成问题，加入互斥锁
var dbLock sync.Mutex

// InstanceDB 获取实例
func InstanceDB() *gorm.DB {
	if instance != nil {
		return instance
	}
	// 锁只是为了避免重复创建数据库Engine
	dbLock.Lock()
	defer dbLock.Unlock()
	if instance != nil {
		return instance
	}

	return NewInstance()
}

// NewInstance create a new instance of engine
func NewInstance() *gorm.DB {
	mysqlConnConfig := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8",
		conf.DbMaster.User,
		conf.DbMaster.Psd,
		conf.DbMaster.Host,
		conf.DbMaster.Port,
		conf.DbMaster.Database,
	)

	engin, err := gorm.Open(mysql.Open(mysqlConnConfig), &gorm.Config{})
	if err != nil {
		log.Fatal("mysql can't instance,err", err)
		return nil
	}
	instance = engin
	return instance
}
