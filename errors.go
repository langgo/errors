package errors

type TipMessageError interface {
	GetMessage() string
}

func WithTipMessage(err error, message string) error {
	return &tipMessageError{
		error:   err,
		message: message,
	}
}

type tipMessageError struct {
	error
	message string
}

func (tm *tipMessageError) GetMessage() string {
	return tm.message
}

func (tm *tipMessageError) Cause() error {
	return tm.error
}
