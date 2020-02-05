package models

import (
	"github.com/liyuliang/sstorage/database"
	"access"
	"fmt"
	"time"
	"strings"
	"log"
)

func init() {
	Register(func() Model {
		return new(gufengmh8_page)
	})
}

type gufengmh8_page struct {
	Code     string
	Database string
	Chapter  string
	Number   string
	Url      string
	Title    string
	Imgs     []string
}

func (m *gufengmh8_page) Name() string {
	return "gufengmh8_page"
}

func (m *gufengmh8_page) Sqls() []string {

	t := new(database.Chapter)
	access.Set(t, m)

	t.Imgs = strings.Join(m.Imgs, ",")

	field, exist := CheckStructExistEmptyVal(t)
	if exist {
		log.Printf("Field %s of table chapter is empty, can't insert into db", field)
		return []string{}
	}

	fields := getTableFields(t)

	sql := fmt.Sprintf(
		`INSERT INTO %s (%s)
SELECT '%s', '%s', '%s', '%s', '%s', '%s', %d FROM dual WHERE NOT EXISTS
(SELECT %s FROM %s 
WHERE code = '%s' AND number = '%s' AND chapter = '%s'
)`,
		t.TableName(),
		fields,

		t.Code,
		t.Number,
		t.Chapter,
		filterVal(t.Url),
		filterVal(t.Title),
		t.Imgs,
		time.Now().Unix(),

		fields,
		t.TableName(),

		t.Code,
		t.Number,
		t.Chapter,
	)

	return []string{
		sql,
	}
}

func (m *gufengmh8_page) Extends() (jobs []*Job) {
	return
}
