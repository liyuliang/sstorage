package access

import "sort"

type Pairs struct {
	Key string
	Val string
}

func SortMap(m map[string]string) (pairs []Pairs) {

	keys := []string{}
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		v := m[k]
		if k != "" && v != "" {
			pairs = append(pairs, Pairs{
				Key: k,
				Val: v,
			})
		}
	}

	return pairs
}
