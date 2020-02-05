package database

func init() {
	Register(func() Table {
		return new(Chapter)
	})
}

type Chapter struct {
	Code       string `xorm:"index"`
	Number     string `xorm:"index"`
	Chapter    string `xorm:"index"`
	Url        string
	Title      string
	Imgs       string `xorm:"'imgs'"`
	UpdateTime int64  `xorm:"'updateTime'"`
}

func (t *Chapter) TableName() string {
	return "chapter"
}
