package Port

import "regexp"

const validPortRegexChecker = "^((6553[0-5])|(655[0-2][0-9])|(65[0-4][0-9]{2})|(6[0-4][0-9]{3})|([1-5][0-9]{4})|([0-5]{0,5})|([0-9]{1,4}))$"

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
	match, err := regexp.MatchString(validPortRegexChecker, value)
	if err != nil || !match {
		return nil, NewInvalidPortError()
	}
	self.value = value
	return self, nil
}
