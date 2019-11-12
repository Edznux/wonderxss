package interfaces

type NotificationSystem interface {
	SendMessage(data string, destination string) error
}
