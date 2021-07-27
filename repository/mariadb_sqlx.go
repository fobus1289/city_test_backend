package repository

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"sync"
)

var onceSqlx = &sync.Once{}
var mariaDBSqlx *sqlx.DB

type Config struct {
	Host     string
	Port     string
	DBName   string
	User     string
	Password string
}

func NewMariaDBSqlX(config *Config) *sqlx.DB {

	onceSqlx.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.User,
			config.Password,
			config.Host,
			config.Port,
			config.DBName,
		)
		var err error
		mariaDBSqlx, err = sqlx.Connect("mysql", dsn)
		if err != nil {
			log.Fatalln(err)
		}
	})

	return mariaDBSqlx
}
