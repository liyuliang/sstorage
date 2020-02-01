package database

func init() {
	Register(func() Table {
		return new(Book)
	})
}

type Book struct {
	Code       string `gorm:"index:code"`
	Number     string `gorm:"index:number"`
	Url        string
	Title      string
	Face       string
	Intro      string
	UpdateTime int64  `gorm:"Column:updateTime"`
}

func (t *Book) TableName() string {
	return "book"
}
