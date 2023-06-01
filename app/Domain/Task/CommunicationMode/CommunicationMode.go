package CommunicationMode

type Mode string

const (
	Unary         Mode = "UNARY"
	ServerStream  Mode = "SERVER_STREAM"
	ClientStream  Mode = "CLIENT_STREAM"
	Bidirectional Mode = "BIDIRECTIONAL"
)

func FromString(mode string) (Mode, error) {
	switch mode {
	case "UNARY":
		return Unary, nil
	case "SERVER_STREAM":
		return ServerStream, nil
	case "CLIENT_STREAM":
		return ClientStream, nil
	case "BIDIRECTIONAL":
		return Bidirectional, nil
	}
	return "", NewUnknownError()
}
