package CreateTask

type Command struct {
	Host              string
	Port              string
	CommandSentences  []string
	CommunicationMode string
	ExecutionMode     string
}
