package dns

type Header struct {
	ID      uint16
	Flags   *Flags
	QDCount uint16
	ANCount uint16
	NSCount uint16
	ARCount uint16
}

func (header *Header) Read(reader *reader) *Header {
	header.ID = reader.readUInt16()
	header.Flags = (&Flags{}).Read(reader)
	header.QDCount = reader.readUInt16()
	header.ANCount = reader.readUInt16()
	header.NSCount = reader.readUInt16()
	header.ARCount = reader.readUInt16()
	return header
}
