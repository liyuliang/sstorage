package access

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
)

func ToByte(obj interface{}) []byte {

	targetObj := reflect.ValueOf(obj)
	structElem := targetObj.Elem()
	structType := structElem.Type()

	result := `{`
	for i := 0; i < structElem.NumField(); i++ {
		fieldName := structType.Field(i).Name
		structField := structElem.FieldByName(fieldName)
		fieldValue := structField.Interface()
		fieldValueType := structField.Type().String()

		switch fieldValue.(type) {

		default:
			if strings.Contains(fieldValueType, "*") {
				result += `"` + fieldName + `":` + ObjToStr(fieldValue) + `,`
			} else {
				result += `"` + fieldName + `":"` + ValToStr(fieldValue) + `",`
			}

		case []string:
			result += `"` + fieldName + `":[`
			result += ValToStr(fieldValue)
			result += "],"
		case map[string]string:
			result += `"` + fieldName + `":{`
			result += ValToStr(fieldValue)
			result += "},"
		case map[string]interface{}:
			result += `"` + fieldName + `":{`
			result += ValToStr(fieldValue)
			result += "},"
		case []Pairs:
			result += `"` + fieldName + `":{`
			result += ValToStr(fieldValue)
			result += "},"
		}
	}

	result += `}`
	result = strings.Replace(result, ",]", "]", -1)
	result = strings.Replace(result, ",}", "}", -1)
	return []byte(result)
}

func ValToStr(val interface{}) (result string) {

	switch val.(type) {
	case string:
		result = val.(string)
	case int:
		result = strconv.Itoa(val.(int))
	case int32:
		result = strconv.FormatInt(int64(val.(int32)), 10)
	case int64:
		result = strconv.FormatInt(val.(int64), 10)
	case float32:
		result = strconv.FormatFloat(float64(val.(float32)), 'f', -1, 32)
	case float64:
		result = strconv.FormatFloat(val.(float64), 'f', -1, 64)
	case json.Number:
		result = val.(json.Number).String()
	case []string:
		for _, value := range val.([]string) {
			result += `"` + value + `",`
		}
		if len(result)-1 > 0 {
			result = result[:len(result)-1]
		}
	case map[string]string:
		fields := SortMap(val.(map[string]string))

		for _, field := range fields {
			if strings.Index(field.Val, "{") == 0 && strings.Index(field.Val, "}") == (len(field.Val)-1) {
				result += `"` + field.Key + `":` + field.Val + `,`
			} else {
				result += `"` + field.Key + `":"` + field.Val + `",`
			}
		}

	case map[string]interface{}:
		field := make(map[string]string)
		for k, v := range val.(map[string]interface{}) {
			field[k] = ValToStr(v)
		}
		fields := SortMap(field)

		for _, field := range fields {
			if strings.Index(field.Val, "{") == 0 && strings.Index(field.Val, "}") == (len(field.Val)-1) {
				result += `"` + field.Key + `":` + field.Val + `,`
			} else {
				result += `"` + field.Key + `":"` + field.Val + `",`
			}
		}
	case []Pairs:
		for _, field := range val.([]Pairs) {
			if strings.Index(field.Val, "{") == 0 && strings.Index(field.Val, "}") == (len(field.Val)-1) {
				result += `"` + field.Key + `":` + field.Val + `,`
			} else {
				result += `"` + field.Key + `":"` + field.Val + `",`
			}
		}
	}

	result = strings.TrimSpace(result)
	return result
}

func ObjToStr(obj interface{}) (result string) {

	targetObj := reflect.ValueOf(obj)
	structElem := targetObj.Elem()
	structType := structElem.Type()
	result += "{"
	for i := 0; i < structElem.NumField(); i++ {

		fieldName := structType.Field(i).Name
		fieldValue := structElem.FieldByName(fieldName).Interface()

		value := ValToStr(fieldValue)
		if strings.Index(value, "{") == 0 && strings.Index(value, "}") == (len(value)-1) {
			result += `"` + fieldName + `":` + value + `,`
		} else {
			result += `"` + fieldName + `":"` + value + `",`
		}
	}
	result += "}"
	return result
}
