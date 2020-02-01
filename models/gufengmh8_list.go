package models

import (
	"github.com/liyuliang/sstorage/database"
	"access"
	"time"
	"github.com/liyuliang/utils/format"
	"fmt"
	"log"
)

func init() {
	Register(func() Model {
		return new(gufengmh8_list)
	})
}

type gufengmh8_list struct {
	Code     string
	Database string
	Intro    string
	Face     string
	Number   string
	Url      string
	Title    string
	Pages    []string
}

func (m *gufengmh8_list) Name() string {
	return "gufengmh8_list"
}

func (m *gufengmh8_list) Sqls() []string {

	t := new(database.Book)
	access.Set(t, m)

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
WHERE code = '%s' AND number = '%s'
)`,
		t.TableName(),
		fields,

		t.Code,
		t.Number,
		t.Url,
		t.Title,
		t.Face,
		t.Intro,
		time.Now().Unix(),

		fields,
		t.TableName(),

		t.Code,
		t.Number,
	)

	return []string{
		sql,
	}
}

func (m *gufengmh8_list) Extends() (jobs []*Job) {

	j := new(Job)
	j.Type = "page" //TODO
	j.Token = format.Int64ToStr(time.Now().UnixNano())

	for _, url := range m.Pages {

		if url != "" {
			j.Urls = append(j.Urls, url)
		}
	}

	jobs = append(jobs, j)
	return jobs
}
