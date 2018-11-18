package param

import (
	"encoding/json"
)

type SearchReturnFields struct {
	prop []string
}

func NewSearchReturnFields() *SearchReturnFields {
	return &SearchReturnFields{}
}

func (searchReturnFields *SearchReturnFields) AddReturnField(key string) {
	searchReturnFields.prop = append(searchReturnFields.prop, key)
}

func (searchReturnFields *SearchReturnFields) GetString() string {
	data, err := json.Marshal(searchReturnFields.prop)
	if nil != err {
		return ""
	}
	return string(data)
}
