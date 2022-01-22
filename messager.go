package cmdstream

type Messager interface {
	SendMessage(message []byte) error
}
