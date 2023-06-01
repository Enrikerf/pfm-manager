package ExecutionMode

type Mode string

const (
	Manual    Mode = "MANUAL"
	Automatic Mode = "AUTOMATIC"
)

func FromString(mode string) (Mode, error) {
	switch mode {
	case "MANUAL":
		return Manual, nil
	case "AUTOMATIC":
		return Automatic, nil
	}
	return "", NewUnknownError()
}
