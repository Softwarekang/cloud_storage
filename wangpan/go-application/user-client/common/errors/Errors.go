package errors

type Errors struct {
	message string
}

func (e *Errors) Error() string {
	return e.message
}

func (e *Errors) New(message string) error {
	e.message = message
	return e
}
