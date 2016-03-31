package models

type DataFlow struct {
	FlowHeader

	Data	[]byte
}

type TemplateFlow struct {
	FlowHeader

	Records []Template
}

type TemplateOptionsFlow struct {
	FlowHeader

	Records []TemplateOptions
}

type FlowHeader struct {
	Id 	uint16
	Length 	uint16
}
