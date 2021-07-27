package repository

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sync"
)

var onceGorm = &sync.Once{}
var mariaDBGorm *gorm.DB

func NewMariaDBGorm(config *Config) *gorm.DB {

	onceGorm.Do(func() {
		var err error
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.User,
			config.Password,
			config.Host,
			config.Port,
			config.DBName,
		)

		mariaDBGorm, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			SkipDefaultTransaction: true,
			AllowGlobalUpdate:      true,
		})

		if err != nil {
			log.Fatalln(err)
		}

	})

	return mariaDBGorm
}
