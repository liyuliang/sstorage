package models

import (
	"sync"
	"strings"
)

type Model interface {
	Name() string
	Sqls() []string
}

type Creator func() Model

var _list []Creator

func Register(method Creator) {
	_list = append(_list, method)
}

type Models map[string]Creator

type parserList struct {
	sync.RWMutex
	creators Models
}

var list parserList

func List() Models {

	if len(list.creators) != len(_list) {

		list = parserList{}
		list.creators = make(Models)
		list.Lock()

		for _, agent := range _list {
			list.creators[agent().Name()] = agent
		}
		list.Unlock()
	}

	return list.creators
}

func Get(name string) (creator Creator) {
	for _, agent := range List() {
		if strings.ToLower(agent().Name()) == strings.ToLower(name) {
			creator = agent
			break
		}
	}
	return creator
}
