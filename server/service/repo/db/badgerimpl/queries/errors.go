package queries

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrKeyExists = errors.New("key exists")
)

type InvalidKeyPrefix struct {
	expected string
	got      string
}

func NewInvalidKeyPrefix(expected, got string) error {
	return &InvalidKeyPrefix{
		expected: expected,
		got:      got,
	}
}

func (e *InvalidKeyPrefix) Error() string {
	var b strings.Builder
	b.WriteString("invalid key prefix: expected ")
	b.WriteString(strconv.Quote(e.expected))
	b.WriteString(", got ")
	b.WriteString(strconv.Quote(e.got))
	return b.String()
}

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
