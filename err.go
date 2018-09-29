package errors

import (
	stderrors "errors"
	"fmt"
	"io"
	"strings"
)

type Error struct {
	*stack

	Err     error             // 只会是外部错误，保持最早的栈
	Message map[string]string // TODO 大多数应该都用不到
	isTip   bool
	tipMsg  string
}

func New(text string) *Error {
	return &Error{
		stack: callers(),
		Err:   stderrors.New(text),
	}
}

func Wrap(err error) *Error {
	if ee, ok := err.(*Error); ok {
		return ee
	}

	return &Error{
		stack: callers(),
		Err:   err,
	}
}

// if Error is Upper。 must care where e is nil
func (e *Error) IsTip() *Error {
	if e == nil {
		return nil
	}

	e.isTip = true
	return e
}

func (e *Error) WithMessage(key, text string) *Error {
	if e == nil {
		return nil
	}

	if e.Message == nil {
		e.Message = make(map[string]string)
	}

	e.Message[key] = text
	return e
}

func (e *Error) WithTip(text string) *Error {
	if e == nil {
		return nil
	}

	e.tipMsg = text
	return e
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}

	if e.Err != nil {
		msg := ""
		if len(e.Message) > 0 {
			msgs := make([]string, 0, len(e.Message))
			for key, value := range e.Message {
				msgs = append(msgs, fmt.Sprintf("%s: %s", key, value))
			}
			msg = fmt.Sprintf("[%s]", strings.Join(msgs, ", "))
		}

		return e.Err.Error() + msg
	}

	return ""
}

func (e *Error) E() error {
	if e == nil {
		return nil
	}

	return e
}

func (e *Error) Format(s fmt.State, verb rune) {
	msgs := make([]string, 0, len(e.Message))
	for key, value := range e.Message {
		msgs = append(msgs, fmt.Sprintf("%s: %s", key, value))
	}
	msg := fmt.Sprintf("[%s]", strings.Join(msgs, ", "))

	switch verb {
	case 'v':
		if s.Flag('+') {
			io.WriteString(s, e.Error()+" ")
			io.WriteString(s, msg)
			e.stack.Format(s, verb)
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, e.Error())
	case 'q':
		fmt.Fprintf(s, "%q", e.Error())
	}
}

func (e *Error) GetTipMsg() string {
	if e.tipMsg != "" {
		return e.tipMsg
	}

	if e.isTip {
		return e.Error()
	}

	return ""
}
