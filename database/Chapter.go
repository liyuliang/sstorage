package database

func init() {
	Register(func() Table {
		return new(Chapter)
	})
}

type Chapter struct {
	Code       string `gorm:"index:code"`
	Number     string `gorm:"index:number"`
	Chapter    string `gorm:"index:chapter"`
	Url        string
	Title      string
	Imgs       string `gorm:"Column:imgs"`
	UpdateTime int64  `gorm:"Column:updateTime"`
}

func (t *Chapter) TableName() string {
	return "chapter"
}
