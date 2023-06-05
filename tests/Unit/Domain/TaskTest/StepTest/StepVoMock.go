package StepTest

import "github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Step"

const defaultSentence = "DEFAULT_SENTENCE"
const invalidSentence = "tooLongSentence-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------"

type voMock struct {
	sentence string
}

func (v *voMock) GetSentence() string {
	return v.sentence
}

func BuildDefaultStepVo() Step.Vo {
	return &voMock{defaultSentence}
}
