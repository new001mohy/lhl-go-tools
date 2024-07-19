package error

type ErrSystemError struct {
}

func (s ErrSystemError) Error() string {
	return ""
}
