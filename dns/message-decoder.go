package dns

func Decode(data []byte) *Message {
	reader := reader{data: data, offset: 0}
	message := Message{}
	message.Header = NewHeader(&reader)
	return &message
}
