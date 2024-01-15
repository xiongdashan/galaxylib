package galaxylib

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var DbInstance *gorm.DB

// test env
func init() {
	dsn := mysql.Config{
		Addr:                 GalaxyCfgFile.MustValue("database", "server_address"),
		User:                 GalaxyCfgFile.MustValue("database", "user_name"),
		Passwd:               GalaxyCfgFile.MustValue("database", "password"),
		Net:                  "tcp",
		DBName:               GalaxyCfgFile.MustValue("database", "db_name"),
		Params:               map[string]string{"charset": "utf8", "parseTime": "True", "loc": "Local"},
		Timeout:              5 * time.Second,
		AllowNativePasswords: true,
	}

	db, err := gorm.Open("mysql", dsn.FormatDSN())

	if err != nil {
		//fmt.Println(err)
		panic(err.Error())
	}

	DbInstance = db
}
