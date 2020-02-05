package access

import (
	"reflect"
	"strings"
)

func SetField(obj interface{}, name string, value interface{}) {

	object := reflect.ValueOf(obj)
	structValue := object.Elem()
	structType := structValue.Type()

	for i := 0; i < structValue.NumField(); i++ {
		fieldName := structType.Field(i).Name
		if strings.ToLower(fieldName) == strings.ToLower(name) {

			structFieldValue := structValue.FieldByName(fieldName)

			if structFieldValue.IsValid() {

				if structFieldValue.CanSet() {
					structFieldType := structFieldValue.Type()
					val := reflect.ValueOf(value)

					if structFieldType == val.Type() {
						structFieldValue.Set(val)
					}
				}
			}
		}
	}
}

func SetMap(obj interface{}, m map[string]interface{}) {
	for k, v := range m {
		SetField(obj, k, v)
	}
}

func Set(obj interface{}, fill interface{}) {

	targetObj := reflect.ValueOf(obj)
	structElem := targetObj.Elem()
	structType := structElem.Type()

	fillObj := reflect.ValueOf(fill)
	fillStructElem := fillObj.Elem()

	for i := 0; i < structElem.NumField(); i++ {

		fieldName := structType.Field(i).Name
		structField := structElem.FieldByName(fieldName)
		fillStructField := fillStructElem.FieldByName(fieldName)
		if structField.IsValid() && structField.CanSet() && fillStructField.IsValid() {
			if structField.Type() == fillStructField.Type() {
				structField.Set(fillStructField)
			}
		}
	}
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
