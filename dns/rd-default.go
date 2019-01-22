package dns

import (
	"fmt"

	"github.com/xyzrlee/go/logger"
)

type DefaultRD struct {
	Data []byte
}

func (defaultRD *DefaultRD) Read(reader *reader, length uint16) ResourceData {
	defaultRD.Data = reader.read(int(length))
	return defaultRD
}

func (defaultRD *DefaultRD) String() string {
	logger.Debug("DefaultRD String()")
	return fmt.Sprintf("% x", defaultRD.Data)
}
