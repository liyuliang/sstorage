package database

func init() {
	Register(func() table {
		return new(Chapter)
	})
}

type Chapter struct {
	Code       string
	Number     string
	Chapter    string
	Url        string
	Title      string
	UpdateTime int64
}

func (t *Chapter) TableName() string {
	return "chapter"
}
