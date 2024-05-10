package queries

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrKeyNotFound = errors.New("key not found")
)

type InvalidKeySizeError struct {
	expected int
	got      int
}

func NewInvalidKeySizeError(expected, got int) error {
	return &InvalidKeySizeError{
		expected: expected,
		got:      got,
	}
}

func (e *InvalidKeySizeError) Error() string {
	var b strings.Builder
	b.WriteString("invalid key size: expected ")
	b.WriteString(strconv.Itoa(e.expected))
	b.WriteString(", got ")
	b.WriteString(strconv.Itoa(e.got))
	return b.String()
}
