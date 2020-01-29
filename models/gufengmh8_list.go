package models

import (
	"github.com/liyuliang/sstorage/database"
	"access"
)

func init() {
	Register(func() Model {
		return new(gufengmh8_list)
	})
}

type gufengmh8_list struct {
	Code     string
	Database string
	Desc     string
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

	b := new(database.Book)
	access.Set(b, m)

	var chapters []*database.Chapter
	for _, url := range m.Pages {

		chapter := new(database.Chapter)
		chapter.Url = url
		chapters = append(chapters, chapter)
	}

	return []string{

	}
}
