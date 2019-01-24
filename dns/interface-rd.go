package dns

type ResourceData interface {
	Read(*reader, uint16) ResourceData
	String() string
}

func GetRD(rrtype RRType) ResourceData {
	switch rrtype {
	case RRTypeA:
		return new(A)
	default:
		return new(DefaultRD)
	}
}
