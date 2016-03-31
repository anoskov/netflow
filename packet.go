package netflow

type Packet struct {
	Version   		uint16
	Count     		uint16
	SysUpTime 		uint32
	UnixSecs  		uint32
	SequenceNumber  	uint32
	SourceId 		uint32
	FlowSets 		[]interface{}
}

func (p *Packet) Templates() (list []*Template) {
	for i := range p.FlowSets {
		switch set := p.FlowSets[i].(type) {
		case TemplateFlow:
			for j := range set.Records {
				list = append(list, &set.Records[j])
			}
		}
	}
	return
}

func (p *Packet) TemplateOptions() (list []*TemplateOptions) {
	for i := range p.FlowSets {
		switch set := p.FlowSets[i].(type) {
		case TemplateOptionsFlow:
			for j := range set.Records {
				list = append(list, &set.Records[j])
			}
		}
	}
	return
}
