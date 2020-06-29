package param

import (
	"encoding/json"
)

const (
	STRING  = "string"
	INTEGER = "integer"
	FLOAT   = "float"
	DOUBLE  = "double"
	BOOLEAN = "boolean"
	LONG    = "long"
)

type CreateExternalFields struct {
	prop map[string]map[string]string
}

func NewCreateExternalFields() *CreateExternalFields {
	return &CreateExternalFields{prop: map[string]map[string]string{}}
}

func (createExternalFields *CreateExternalFields) AddField(fieldName string, fieldType string) {
	createExternalFields.prop[fieldName] = map[string]string{"type": fieldType}
}

func (createExternalFields *CreateExternalFields) GetString() string {
	data, err := json.Marshal(createExternalFields.prop)
	if nil != err {
		return ""
	}
	return string(data)
}

func (createExternalFields *CreateExternalFields) GetValue() map[string]map[string]string {
	return createExternalFields.prop
}
