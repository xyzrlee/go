package dns

type Name struct {
	Names []string
}

func NewName() *Name {
	return &Name{}
}

func (name *Name) Append(str string) *Name {
	name.Names = append(name.Names, str)
	return name
}
