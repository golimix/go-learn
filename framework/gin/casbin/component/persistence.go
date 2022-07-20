package component

import (
	"fmt"
	"time"

	"github.com/allegro/bigcache"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	DB          *gorm.DB
	GlobalCache *bigcache.BigCache
)

func init() {
	// Connect to DB
	var err error
	DB, err = gorm.Open("mysql", "root:iflytek@tcp(127.0.0.1:3306)/appmanage?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(fmt.Sprintf("failed to connect to DB: %v", err))
	}

	// Initialize cache
	GlobalCache, err = bigcache.NewBigCache(bigcache.DefaultConfig(30 * time.Minute)) // Set expire time to 30 mins
	if err != nil {
		panic(fmt.Sprintf("failed to initialize cahce: %v", err))
	}
}
