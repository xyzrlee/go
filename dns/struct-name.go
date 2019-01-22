package dns

import "strings"

type Name struct {
	Names []string
}

func NewName() *Name {
	return &Name{}
}

func (name *Name) AppendString(str string) *Name {
	name.Names = append(name.Names, str)
	return name
}

func (name *Name) AppendName(xname *Name) *Name {
	name.Names = append(name.Names, xname.Names...)
	return name
}

func (name *Name) String() string {
	return strings.Join(name.Names, ".") + "."
}
