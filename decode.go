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