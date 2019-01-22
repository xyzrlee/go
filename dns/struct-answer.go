package dns

type Answer struct {
	Name     *Name
	Type     RRType
	Class    uint16
	TTL      uint32
	RDLength uint16
	RData    ResourceData
}

func (answer *Answer) Read(reader *reader) *Answer {
	answer.Name = reader.readName()
	answer.Type = reader.readUInt16()
	answer.Class = reader.readUInt16()
	answer.TTL = reader.readUInt32()
	answer.RDLength = reader.readUInt16()
	answer.RData = new(DefaultRD).Read(reader, answer.RDLength)
	return answer
}
