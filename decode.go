package netflow

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"cmd/go/testdata/testinternal3"
)

func extractFieldList(buf *bytes.Buffer, count int) (list []Field) {
	list = make([]Field, count)

	for i := 0; i < count; i++ {
		binary.Read(buf, binary.BigEndian, &list[i])
	}

	return
}

func parseDataFlow(data []byte, header *FlowHeader) (interface{}, error) {
	var flow DataFlow

	flow.Id = header.Id
	flow.Length = header.Length
	flow.Data = data

	return flow, nil
}

func parseOptionsTemplateFlowSet(data []byte, header *FlowHeader) (interface{}, error) {
	var flow TemplateOptionsFlow
	var tplOpt TemplateOptions

	flow.Id = header.Id
	flow.Length = header.Length

	buf := bytes.NewBuffer(data)
	headerLen := binary.Size(tplOpt.TemplateId) + binary.Size(tplOpt.ScopeLength) + binary.Size(tplOpt.OptionLength)
	for buf.Len() >= 4 {
		if buf.Len() < headerLen {
			return nil, errorIncompletePacket(headerLen - buf.Len())
		}
		binary.Read(buf, binary.BigEndian, &tplOpt.TemplateId)
		binary.Read(buf, binary.BigEndian, &tplOpt.ScopeLength)
		binary.Read(buf, binary.BigEndian, &tplOpt.OptionLength)

		if buf.Len() < int(tplOpt.ScopeLength)+int(tplOpt.OptionLength) {
			return nil, errorIncompletePacket(int(tplOpt.ScopeLength) + int(tplOpt.OptionLength) - buf.Len())
		}

		scopeCount := int(tplOpt.ScopeLength) / binary.Size(Field{})
		optionCount := int(tplOpt.OptionLength) / binary.Size(Field{})

		tplOpt.Scopes = extractFieldList(buf, scopeCount)
		tplOpt.Options = extractFieldList(buf, optionCount)

		flow.Records = append(flow.Records, tplOpt)
	}

	return flow, nil
}

func parseTemplateFlow(data []byte, header *FlowHeader) (interface{}, error) {
	var flow TemplateFlow
	var tpl Template

	flow.Id = header.Id
	flow.Length = header.Length

	buf := bytes.NewBuffer(data)
	headerLen := binary.Size(tpl.Id) + binary.Size(tpl.FieldCount)

	for buf.Len() >= 4 {
		if buf.Len() < headerLen {
			return nil, errorIncompletePacket(headerLen - buf.Len())
		}
		binary.Read(buf, binary.BigEndian, &tpl.Id)
		binary.Read(buf, binary.BigEndian, &tpl.FieldCount)

		fieldsLen := int(tpl.FieldCount) * binary.Size(Field{})
		if fieldsLen > buf.Len() {
			return nil, errorIncompletePacket(fieldsLen - buf.Len())
		}
		tpl.Fields = extractFieldList(buf, int(tpl.FieldCount))

		flow.Records = append(flow.Records, tpl)
	}
	return flow, nil

}

/* Error functions */

func errorIncompatibleVersion(version uint16) error {
	return fmt.Errorf("incompatible protocol version v%d, only v9 is supported!", version)
}

func errorIncompletePacket(bytes int) error {
	return fmt.Errorf("incomplete packet, missing %d bytes.", bytes)
}

func errorExcessBytes(bytes int) error {
	return fmt.Errorf("excess %d bytes at the end of the packet.", bytes)
}