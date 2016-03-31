package models

type DataFlow struct {
	Data	[]byte
	FlowHeader
}

type FlowHeader struct {
	Id 	uint16
	Length 	uint16
}
