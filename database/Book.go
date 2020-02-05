package database

func init() {
	Register(func() Table {
		return new(Book)
	})
}

type Book struct {
	Code       string `xorm:"index"`
	Number     string `xorm:"index"`
	Url        string
	Title      string
	Face       string
	Intro      string
	UpdateTime int64  `xorm:"'updateTime'"`
}

func (t *Book) TableName() string {
	return "book"
}
