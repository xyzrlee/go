package dns

type Header struct {
	ID      uint16
	Flags   *Flags
	QDCount uint16
	ANCount uint16
	NSCount uint16
	ARCount uint16
}

func NewHeader(reader *reader) *Header {
	header := Header{}
	header.ID = reader.readUInt16()
	header.Flags = NewFlags(reader)
	header.QDCount = reader.readUInt16()
	header.ANCount = reader.readUInt16()
	header.NSCount = reader.readUInt16()
	header.ARCount = reader.readUInt16()
	return &header
}
