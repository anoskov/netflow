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
	var set DataFlow

	set.Id = header.Id
	set.Length = header.Length
	set.Data = data

	return set, nil
}

func errorIncompletePacket(bytes int) error {
	return fmt.Errorf("incomplete packet, missing %d bytes.", bytes)
}