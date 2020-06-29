package param

import (
	"encoding/json"
)

const (
	ASC  = "asc"
	DESC = "desc"
)

type SearchSort struct {
	prop []map[string]string
}

func NewSearchSort() *SearchSort {
	return &SearchSort{}
}

func (searchSort *SearchSort) AddSortField(key string, sortType string) {
	sortField := map[string]string{}
	sortField[key] = sortType
	searchSort.prop = append(searchSort.prop, sortField)
}

func (searchSort *SearchSort) GetString() string {
	data, err := json.Marshal(searchSort.prop)
	if nil != err {
		return ""
	}
	return string(data)
}

func (searchSort *SearchSort) GetValue() []map[string]string {
	return searchSort.prop
}
