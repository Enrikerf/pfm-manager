package Step

type Vo interface {
	GetSentence() string
}

type vo struct {
	sentence string
}

func (v *vo) GetSentence() string {
	return v.sentence
}

func NewVo(sentence string) (Vo, error) {
	self := &vo{}
	if len(sentence) > 255 {
		return nil, NewInvalidSentenceLengthError()
	}
	self.sentence = sentence
	return self, nil
}
