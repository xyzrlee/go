package dns

import (
	"github.com/xyzrlee/go/logger"
)

type Header struct {
	ID      uint16
	Flags   *Flags
	QDCount uint16
	ANCount uint16
	NSCount uint16
	ARCount uint16
}

func NewHeader(reader *reader) *Header {
	logger.Debugw("in NewHeader")
	header := Header{}
	header.ID = reader.readUInt16()
	logger.Debugw("", "header.ID", header.ID)
	header.Flags = NewFlags(reader)
	logger.Debugw("", "header.Flags", header.Flags)
	header.QDCount = reader.readUInt16()
	header.ANCount = reader.readUInt16()
	header.NSCount = reader.readUInt16()
	header.ARCount = reader.readUInt16()
	return &header
}
