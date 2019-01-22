package dns

type Flags struct {
	QR     bool
	Opcode uint8
	AA     bool
	TC     bool
	RD     bool
	RA     bool
	RCode  uint8
}

func NewFlags(reader *reader) *Flags {
	data := reader.readUInt16()
	flags := Flags{}
	flags.QR = data&0x8000>>15 == 1
	flags.Opcode = uint8(data & 0x7800 >> 11)
	flags.AA = data&0x0400>>10 == 1
	flags.TC = data&0x0200>>9 == 1
	flags.RD = data&0x0100>>8 == 1
	flags.RA = data&0x0080>>7 == 1
	flags.RCode = uint8(data & 0x000f)
	return &flags
}
