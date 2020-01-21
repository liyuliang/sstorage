package format

import "math"

func SliceDelDuplicate(slice []string) []string {
	result := make([]string, 0, len(slice))

	temp := map[string]struct{}{}

	for _, item := range slice {
		if _, ok := temp[item]; !ok {

			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func InArray(items []string, item string) bool {
	for _, a := range items {
		if a == item {
			return true
		}
	}
	return false
}

func SliceChunk(slice []interface{}, size int) (data [][]interface{}) {

	l := len(slice)
	groupLen := int(math.Ceil(float64(l/size))) + 1

	for i := 0; i < groupLen; i++ {

		start := i * size
		end := start + size

		var newSlice []interface{}

		newSlice = slice[start:end]

		data = append(data, newSlice)
	}
	return data
}
