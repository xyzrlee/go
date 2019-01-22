package dns

type Question struct {
	QName  *Name
	QType  RRType
	QClass uint16
}

func (question *Question) Read(reader *reader) *Question {
	question.QName = reader.readName()
	question.QType = reader.readUInt16()
	question.QClass = reader.readUInt16()
	return question
}
