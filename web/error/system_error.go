package error

type ErrSystemError struct {
	Message string
}

func (s ErrSystemError) Error() string {
	return s.Message
}
