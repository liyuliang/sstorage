package database

func init() {
	Register(func() table {
		return new(Book)
	})
}

type Book struct {
	Code       string
	Number     string
	Url        string
	Title      string
	Face       string
	Desc       string
	UpdateTime int64
}

func (t *Book) TableName() string {
	return "book"
}
