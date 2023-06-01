package Content

type Content interface {
	GetValue() string
}

type content struct {
	value string
}

func (c *content) GetValue() string {
	return c.value
}

func NewContent(value string) (Content, error) {
	if len(value) > 255 {
		return nil, NewMaxLengthExceed()
	}
	self := &content{}
	self.value = value
	return self, nil
}
