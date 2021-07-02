package galaxylib

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	Conn string
}

func GalaxyDB() *DB {
	conn := GalaxyCfgFile.MustValue("db", "conn")
	return &DB{conn}
}

//OpenDb 打开数据库并执行
func (d *DB) OpenDb(f func(*gorm.DB)) {
	//fmt.Println(conn)
	db, err := gorm.Open(postgres.Open(d.Conn), &gorm.Config{}) //gorm.Open("postgres", d.Conn)
	//db.LogMode(true)
	if err != nil {
		panic(err)
	}
	defer func() {
		sqlDb, err := db.DB()
		if sqlDb != nil && err != nil {
			sqlDb.Close()
		}
	}()
	f(db)
}

func (d *DB) Add(val interface{}) (ret int64) {

	d.OpenDb(func(db *gorm.DB) {
		rst := db.Create(&val)
		if rst.Error != nil {
			fmt.Println(rst.Error)
			return
		}
		ret = rst.RowsAffected
	})
	return
}
