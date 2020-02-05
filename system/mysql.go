package system

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/liyuliang/xorm"
	"fmt"
	"log"
)

var mysqlHandler *mysqlDB

type mysqlDB struct {
	connector *xorm.Engine
}

func Mysql() (*xorm.Engine) {
	if mysqlHandler == nil {
		mysqlHandler = newMysql()
	}
	return mysqlHandler.conn()
}

func newMysql() *mysqlDB {

	mysqlHandler = new(mysqlDB)
	mysqlHandler.conn()
	return mysqlHandler
}

func (db *mysqlDB) conn() (*xorm.Engine) {

	if db.connector == nil {

		username := Config()[SystemMysqlUserName]
		password := Config()[SystemMysqlPwd]
		host := Config()[SystemMysqlHost]
		port := Config()[SystemMysqlPort]
		database := Config()[SystemMysqlDatabase]
		charset := Config()[SystemMysqlCharset]

		uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", username, password, host, port, database, charset)

		engine, err := xorm.NewEngine("mysql", uri)

		//conn, err := gorm.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+database+"?charset="+charset+"&parseTime=true&loc=Local")
		if err != nil {
			log.Printf("%s,%s,%s,%s,%s", username, host, port, database, charset)
			panic("failed to connect database :" + err.Error())
		}

		//conn.LogMode(mysqlConfig.Debug)
		db.connector = engine
	}

	return db.connector
}

func (db *mysqlDB) CreateTable(table interface{}) (err error) {
	err = db.conn().Sync2(table)
	//db.conn().SingularTable(true)
	//if !db.conn().HasTable(table) {
	//	err = db.conn().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(table).Error
	//}
	return err
}
