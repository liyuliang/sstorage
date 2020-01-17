package system

import (
	"sync"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

var mysqlHandler *mysqlDB

type mysqlDB struct {
	username  string
	password  string
	mutex     *sync.RWMutex
	connector *gorm.DB
}

func Mysql() (*gorm.DB) {
	if mysqlHandler == nil {
		mysqlHandler = newMysql()
	}
	return mysqlHandler.conn()
}

func newMysql() *mysqlDB {

	mysqlHandler = new(mysqlDB)
	mysqlHandler.mutex = new(sync.RWMutex)
	mysqlHandler.conn()
	return mysqlHandler
}

func (db *mysqlDB) conn() (*gorm.DB) {

	if db.connector == nil {

		username := Config()[SystemMysqlUserName]
		password := Config()[SystemMysqlPwd]
		host := Config()[SystemMysqlHost]
		port := Config()[SystemMysqlPort]
		database := Config()[SystemMysqlDatabase]
		charset := Config()[SystemMysqlCharset]

		conn, err := gorm.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+database+"?charset="+charset+"&parseTime=true&loc=Local")
		if err != nil {
			panic("failed to connect database " + err.Error())
		}

		//conn.LogMode(mysqlConfig.Debug)
		db.connector = conn
	}

	return db.connector
}

func (db *mysqlDB) CreateTable(table interface{}) (err error) {

	if !db.conn().HasTable(table) {
		err = db.conn().CreateTable(table).Error
	}
	return err
}
