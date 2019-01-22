package dns

func Decode(data []byte) *Message {
	reader := &reader{data: data, offset: 0}
	message := Message{}
	message.Header = decodeHeader(reader)
	message.Questions = decodeQuestions(reader, message.Header.QDCount)
	return &message
}

func decodeHeader(reader *reader) *Header {
	return NewHeader(reader)
}

func decodeQuestions(reader *reader, qdcount uint16) []*Question {
	questions := make([]*Question, 0, 5)
	for i := 0; i < int(qdcount); i++ {
		questions = append(questions, NewQuestion(reader))
	}
	return questions
}
