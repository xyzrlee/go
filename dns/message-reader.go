package dns

import (
	"encoding/binary"

	"github.com/xyzrlee/go/logger"
)

type reader struct {
	data   []byte
	offset int
}

func (reader *reader) read(length int) []byte {
	if reader.offset+length > len(reader.data) {
		logger.Panicw("offset exceeded", "offset", reader.offset, "length", length, "size", len(reader.data))
	}
	result := reader.data[reader.offset:length]
	reader.offset += length
	return result
}

func (reader *reader) readUInt16() uint16 {
	data := reader.read(2)
	result := binary.BigEndian.Uint16(data)
	return result
}
