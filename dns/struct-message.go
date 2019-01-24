package dns

type Message struct {
	Header      *Header
	Questions   []*Question
	Answers     []*Answer
	Authorities []*Authority
	Additional  []*Additional
}

func Decode(data []byte) *Message {
	reader := reader{data: data, offset: 0}
	message := Message{}
	message.Header = decodeHeader(&reader)
	message.Questions = decodeQuestions(&reader, message.Header.QDCount)
	message.Answers = decodeAnswers(&reader, message.Header.ANCount)
	return &message
}

func decodeHeader(reader *reader) *Header {
	header := new(Header).Read(reader)
	// logger.Debugw("header decoded", "header", header)
	return header
}

func decodeQuestions(reader *reader, qdcount uint16) []*Question {
	questions := make([]*Question, 0, 5)
	for i := 0; i < int(qdcount); i++ {
		questions = append(questions, new(Question).Read(reader))
	}
	// logger.Debugw("questions decoded", "questions", questions)
	return questions
}

func decodeAnswers(reader *reader, ancount uint16) []*Answer {
	answers := make([]*Answer, 0, 5)
	for i := 0; i < int(ancount); i++ {
		answers = append(answers, new(Answer).Read(reader))
	}
	// logger.Debugw("answer decoded", "answers", answers)
	return answers
}
