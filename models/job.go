package models

import (
	"github.com/liyuliang/sstorage/database"
	"strings"
	"reflect"
)

type Job struct {
	Token string
	Type  string
	Urls  []string
}


func getTableFields(t database.Table) string {
	fields := StructKeys(t)
	fields = filterFieldId(fields)
	fieldsStr := strings.Join(fields, ",")
	return fieldsStr
}

func StructKeys(obj interface{}) (keys []string) {
	targetObj := reflect.ValueOf(obj)
	structElem := targetObj.Elem()
	structType := structElem.Type()

	for i := 0; i < structElem.NumField(); i++ {

		fieldName := structType.Field(i).Name
		keys = append(keys, fieldName)
	}
	return keys
}

func filterVal(val string) string {
	val = strings.Replace(val, `'`, `\'`, -1)
	return val
}

func filterFieldId(fields []string) (result []string) {
	for _, field := range fields {
		field = strings.ToLower(field)
		if field != "id" {
			result = append(result, field)
		}
	}
	return result
}
