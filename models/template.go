package models

import (
	"bytes"
)

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

/* Methods */

func (tpl *Template) DecodeFlowSet(set *DataFlow) (list []FlowData) {
	var record FlowData
	buf := bytes.NewBuffer(set.Data)

	if set.Id != tpl.Id {
		return
	}


	for i := 0; buf.Len() >= 4; i++ {
		record.Values = extractFieldValues(buf, tpl.Fields)
		list = append(list, record)
	}

	return
}

/* Functions */

func extractFieldValues(buf *bytes.Buffer, fields []Field) (values [][]byte) {
	values = make([][]byte, len(fields))
	for i, f := range fields {
		if buf.Len() < int(f.Length) {
			break
		}
		values[i] = buf.Next(int(f.Length))
	}
	return
}