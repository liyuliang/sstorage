package database

import (
	"sync"
	"strings"
	"github.com/liyuliang/sstorage/system"
	"log"
)

func Init() {

	initDatabase()
	initTables()
}

func initDatabase() {
	//TODO
}

func initTables() {
	db := system.Mysql()
	db.SingularTable(true)

	for _, creator := range List() {
		t := creator()

		table := t.TableName()

		if !db.HasTable(table) {
			db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(t)
			log.Println("done.")
		} else {
			log.Printf("Table %s exist.", table)
		}
	}
}

type Table interface {
	TableName() string
}

type Creator func() Table

var _list []Creator

func Register(method Creator) {
	_list = append(_list, method)
}

type Tables map[string]Creator

type parserList struct {
	sync.RWMutex
	creators Tables
}

var list parserList

func List() Tables {

	if len(list.creators) != len(_list) {

		list = parserList{}
		list.creators = make(Tables)
		list.Lock()

		for _, agent := range _list {
			list.creators[agent().TableName()] = agent
		}
		list.Unlock()
	}

	return list.creators
}

func Get(name string) (creator Creator) {
	for _, agent := range List() {
		if strings.ToLower(agent().TableName()) == strings.ToLower(name) {
			creator = agent
			break
		}
	}
	return creator
}
