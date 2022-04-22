package MagicSpider

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

/**
使用DB存储爬虫的运行状态，该功能只有在分布式的情况下才会开启
使用连接池配置
 */

type StaDB struct {
	gorm.Model
	SpiderName string `gorm:"size:255;column: spiderName"`
	IP string `gorm:"size:255;column: ip"`
	Status string `gorm:"column: status;size:255;"`
	StartTime time.Time `gorm:"column:startTime"`
	EndTime time.Time `gorm:"column: endTime"`
	EndReason string `gorm:"column: endReason; size:255;"`
}

func ConnectStaDB() (*gorm.DB, error){
	db, err := gorm.Open(mysql.Open(S.StaDB), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// 检查表是否存在，如果存在
	exist := db.Migrator().HasTable(&StaDB{})
	if !exist {
		// 如果不存在
		return nil, errors.New("table staDB dont have")
	}
	return db, err
}
