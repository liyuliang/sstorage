package models

import (
	"github.com/liyuliang/sstorage/database"
	"access"
	"fmt"
	"time"
	"strings"
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

	fields := getTableFields(t)

	sql := fmt.Sprintf(
		`INSERT INTO %s (%s)
SELECT '%s', '%s', '%s', '%s', '%s', '%s', %d FROM dual WHERE NOT EXISTS
(SELECT %s FROM %s 
WHERE site = '%s' AND number = '%s' AND chapter = '%s'
)`,
		t.TableName(),
		fields,

		m.Code,
		m.Number,
		m.Chapter,
		filterVal(m.Title),
		filterVal(m.Url),
		strings.Join(m.Imgs, ","),
		time.Now().Unix(),

		fields,
		t.TableName(),

		m.Code,
		m.Number,
		m.Chapter,
	)
	return []string{
		sql,
	}
}

func (m *gufengmh8_page) Extends() (jobs []*Job) {
	return
}
