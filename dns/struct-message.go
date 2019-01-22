package dns

type Message struct {
	Header     *Header
	Questions  []*Question
	Answers    []*Answer
	Authoritys []*Authority
	Additional []*Additional
}
