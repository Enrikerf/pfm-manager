package Step

type Step interface {
	GetId() Id
	GetSentence() string
}

type step struct {
	id       Id
	sentence string
}

func (s step) GetId() Id {
	return s.id
}

func (s step) GetSentence() string {
	return s.sentence
}

func New(stepVo Vo) Step {
	self := &step{}
	self.id = NewId()
	self.sentence = stepVo.GetSentence()
	return self
}
func Load(id Id, stepVo Vo) Step {
	self := &step{}
	self.id = id
	self.sentence = stepVo.GetSentence()
	return self
}
