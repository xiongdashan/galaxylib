package galaxylib

import (
	"fmt"

	"gorm.io/gorm"
)

type DbScope struct {
	Db *gorm.DB
}

func (c *DbScope) Close() {
	sqlDb, _ := c.Db.DB()
	sqlDb.Close()
}

func (c *DbScope) DB() *gorm.DB {
	return c.Db
}

func (c *DbScope) Add(val interface{}) int64 {
	result := c.Db.Debug().Create(val)
	if result.Error != nil {
		fmt.Println(result.Error)
		return 0
	}
	return result.RowsAffected
}

func (c *DbScope) Get(val interface{}, id uint) {
	c.Db.Debug().First(val, id)
}

func (c *DbScope) Update(val interface{}, attr ...interface{}) int64 {
	ret := c.Db.Model(val).Updates(attr)
	if ret.Error != nil {
		fmt.Println(ret)
		return 0
	}
	return ret.RowsAffected
}

//func (c *DbScope) Update()

func (c *DbScope) List(ret []*interface{}, query interface{}, args ...interface{}) {
	result := c.Db.Where(query, args).Find(&ret)
	if result.Error != nil {
		fmt.Println(result.Error)
		ret = nil
	}
	return
}
