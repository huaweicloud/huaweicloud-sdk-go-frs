package param

import (
	"encoding/json"
)

type AddExternalFields struct {
	prop map[string]interface{}
}

func NewAddExternalFields() *AddExternalFields {
	return &AddExternalFields{prop: map[string]interface{}{}}
}

func (addExternalFields *AddExternalFields) AddField(key string, value interface{}) {
	addExternalFields.prop[key] = value
}

func (addExternalFields *AddExternalFields) GetString() string {
	data, err := json.Marshal(addExternalFields.prop)
	if nil != err {
		return ""
	}
	return string(data)
}

func (addExternalFields *AddExternalFields) GetValue() map[string]interface{} {
	return addExternalFields.prop
}
