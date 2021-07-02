package dblib

import (
	"fmt"

	"github.com/xiongdashan/galaxylib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type IDbFactory interface {
	Db() *gorm.DB
	Tran()
	Commit()
	Rollback()
	Close()
}

type DbFactory struct {
	db   *gorm.DB
	conn string
}

func DefaultDbFactory() *DbFactory {
	d := &DbFactory{}
	d.conn = galaxylib.GalaxyCfgFile.MustValue("db", "conn")
	return d
}

func NewDbFactory(conn string) *DbFactory {
	d := &DbFactory{}
	d.conn = conn
	return d
}

func (d *DbFactory) Db() *gorm.DB {
	var err error
	if d.db == nil {
		d.db, err = gorm.Open(postgres.Open(d.conn), &gorm.Config{})
	}
	if err != nil {
		fmt.Println(err)
	}
	return d.db
}

func (d *DbFactory) Tran() {
	d.db = d.db.Begin()
}

func (d *DbFactory) Commit() {
	d.db.Commit()
}

func (d *DbFactory) Rollback() {
	d.db.Rollback()
}

func (d *DbFactory) Close() {
	if d.db != nil {
		sqlDb, _ := d.db.DB()
		sqlDb.Close()
	}
}
