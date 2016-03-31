package models

type Template struct {
	Id 	   	uint16
	FieldCount 	uint16
	Fields     	[]Field
}

type TemplateOptions struct {
	TemplateId  	uint16
	ScopeLength 	uint16
	OptionLength	uint16
	Scopes		[]Field
	Options		[]Field
}