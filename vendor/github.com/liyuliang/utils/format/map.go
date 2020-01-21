package format

import (
	"encoding/json"
	"net/url"
)

type MapData map[string]string

func (m MapData) Values() (data []string) {
	for _, value := range m {
		data = append(data, value)
	}
	return data
}

func (m MapData) Keys() (data []string) {

	for key, _ := range m {
		data = append(data, key)
	}
	return data
}

func (m MapData) String() string {
	byteData, _ := json.Marshal(m)
	return string(byteData)
}

func Map() MapData {
	return make(map[string]string)
}

func ToMap(params map[string]string) MapData {
	mapData := Map()
	for key, value := range params {
		mapData[key] = value
	}
	return mapData
}

func (m MapData) ToUrlVals() url.Values {
	v := url.Values{}
	for key, val := range m {
		v.Add(key, val)
	}
	return v
}
