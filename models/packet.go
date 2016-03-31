package models

type Packet struct {
	Version   	uint16
	Count     	uint16
	SysUpTime 	uint32
	UnixSecs  	uint32
	SequenceNumber  uint32
	SourceId 	uint32
	FlowSets 	[]interface{}
}