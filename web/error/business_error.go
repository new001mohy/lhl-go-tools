package error

type BusinessError struct {
	Message string
}

func (b BusinessError) Error() string {
	return b.Message
}
