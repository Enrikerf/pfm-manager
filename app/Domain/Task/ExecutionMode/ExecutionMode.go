package ExecutionMode

type Mode interface {
	Value() Enum
}

func New(enum Enum) Mode {
	return &mode{enum}
}

func FromString(mode string) (Mode, error) {
	switch mode {
	case "MANUAL":
		return New(Manual), nil
	case "AUTOMATIC":
		return New(Automatic), nil
	}
	return nil, NewUnknownError()
}

type mode struct {
	enum Enum
}

func (s *mode) Value() Enum {
	return s.enum
}

type Enum string

const (
	Manual    Enum = "MANUAL"
	Automatic Enum = "AUTOMATIC"
)
