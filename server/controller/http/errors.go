package http

import "fmt"

type PanicError struct {
	Panic any
	Stack []byte
}

func (e *PanicError) Error() string {
	return fmt.Sprintf("panic caught: %v\n\n%s", e.Panic, e.Stack)
}
