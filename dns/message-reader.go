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
	origin := reader.offset
	reader.offset += length
	result := reader.data[origin:reader.offset]
	return result
}

func (reader *reader) readUInt16() uint16 {
	data := reader.read(2)
	result := binary.BigEndian.Uint16(data)
	return result
}

func (reader *reader) readName() *Name {
	name := NewName()
	length := int(reader.read(1)[0])
	for length > 0 {
		data := reader.read(length)
		str := string(data)
		name.Append(str)
		length = int(reader.read(1)[0])
	}
	return name
}
