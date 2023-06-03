package CommunicationMode

type Mode interface {
	Value() Enum
}

func New(enum Enum) Mode {
	return &mode{enum}
}

func FromString(mode string) (Mode, error) {
	switch mode {
	case "UNARY":
		return New(Unary), nil
	case "SERVER_STREAM":
		return New(ServerStream), nil
	case "CLIENT_STREAM":
		return New(ClientStream), nil
	case "BIDIRECTIONAL":
		return New(Bidirectional), nil
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
	Unary         Enum = "UNARY"
	ServerStream  Enum = "SERVER_STREAM"
	ClientStream  Enum = "CLIENT_STREAM"
	Bidirectional Enum = "BIDIRECTIONAL"
)
