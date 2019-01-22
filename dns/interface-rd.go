package dns

type ResourceData interface {
	Read(*reader, uint16) ResourceData
	String() string
}
