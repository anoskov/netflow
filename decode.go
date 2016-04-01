package netflow

import (
	"bytes"
	"encoding/binary"
)

func extractFieldList(buf *bytes.Buffer, count int) (list []Field) {
	list = make([]Field, count)

	for i := 0; i < count; i++ {
		binary.Read(buf, binary.BigEndian, &list[i])
	}

	return
}
