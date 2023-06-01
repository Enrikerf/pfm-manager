package Host

import "regexp"

const validHostRegexChecker = "^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\\-]*[A-Za-z0-9])$"

type Vo interface {
	GetValue() string
}

type vo struct {
	value string
}

func (v *vo) GetValue() string {
	return v.value
}

func NewVo(value string) (Vo, error) {
	self := &vo{}
	match, err := regexp.MatchString(validHostRegexChecker, value)
	if err != nil || !match {
		return nil, NewInvalidHostError()
	}
	self.value = value
	return self, nil
}
