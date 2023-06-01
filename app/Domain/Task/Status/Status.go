package Status

type Status interface {
	Value() Enum
}

func New(enum Enum) Status {
	return &status{enum}
}

func FromString(mode string) (Status, error) {
	switch mode {
	case "PENDING":
		return New(Pending), nil
	case "RUNNING":
		return New(Running), nil
	case "DONE":
		return New(Done), nil
	case "SUCCESSFUL":
		return New(Successful), nil
	case "FAILED":
		return New(Failed), nil
	}
	return nil, NewUnknownError()
}

type status struct {
	value Enum
}

func (s *status) Value() Enum {
	return s.value
}

type Enum string

const (
	Pending    Enum = "PENDING"
	Running    Enum = "RUNNING"
	Done       Enum = "DONE"
	Successful Enum = "SUCCESSFUL"
	Failed     Enum = "FAILED"
)
