package system

import "github.com/liyuliang/utils/format"

var c format.MapData

func init() {
	c = format.Map()
}

func Config() format.MapData {
	return c
}


