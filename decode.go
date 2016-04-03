package netflow

import (
	"bytes"
	"encoding/binary"
	"fmt"
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