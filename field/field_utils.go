package netflow

import "bytes"

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

