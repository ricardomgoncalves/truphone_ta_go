package errors

import (
	"errors"
	"fmt"
	"strconv"
)

// New is a drop in replacement for the standard library errors module that records
// the location that the error is created.
func New(message string) error {
	return errors.New(message)
}

// Errorf creates a new annotated error and records the location that the
// error is created.  This should be a drop in replacement for fmt.Errorf.
func Errorf(format string, args ...interface{}) error {
	return errors.New(fmt.Sprintf(format, args...))
}

// Iterate errors and calls fn for every underlying error.
func Iterate(err error, fn func(err error) bool) bool {
	if err == nil {
		return false
	}
	e, _ := err.(interface {
		Unwrap() error
	})
	if r := fn(err); !r || e == nil {
		return r
	}
	return Iterate(e.Unwrap(), fn)
}

// Code returns the code representing the kind of an error.
func Code(err error) int {
	for err != nil {
		e, ok := err.(Coded)
		if !ok {
			err = errors.Unwrap(err)
			continue
		}
		return e.Code()
	}
	return 0
}

// IsTemporary returns true if the error is marked as being temporary.
func IsTemporary(err error) bool {
	for err != nil {
		e, ok := err.(Temporary)
		if !ok {
			err = errors.Unwrap(err)
			continue
		}
		return e.IsTemporary()
	}
	return false
}

// Is checks if an error or any of its underlying errors and the instance passed as arguments are the same.
func Is(err error, instance error) bool {
	return Iterate(err, func(err error) bool {
		return err == instance
	})
}

type Coded interface {
	Code() int
}

type Temporary interface {
	IsTemporary() bool
}

// Annotated has extra message to the original error.
type annotated struct {
	error
	message string
}

// Annotate is used to add extra context to an existing error.
func Annotate(other error, message string) error {
	if other == nil {
		return nil
	}
	return annotated{
		error:   other,
		message: message,
	}
}

// Annotatef is used to add extra context to an existing error.
func Annotatef(other error, format string, args ...interface{}) error {
	if other == nil {
		return nil
	}
	return annotated{
		error:   other,
		message: fmt.Sprintf(format, args...),
	}
}

// Returns the string of the error
func (e annotated) Error() string {
	return e.error.Error() + ": " + e.message
}

// Unwrap returns original error.
func (e annotated) Unwrap() error {
	return e.error
}

// Annotated has extra message to the original error.
type wrapped struct {
	error
	outer error
}

// Wrap changes the Cause of the error.
func Wrap(inner, outer error) error {
	return wrapped{
		error: inner,
		outer: outer,
	}
}

// Returns the string of the error
func (e wrapped) Error() string {
	return e.outer.Error() + ": " + e.error.Error()
}

// Unwrap returns original error.
func (e wrapped) Unwrap() error {
	return e.error
}

// Error with a code
type coded struct {
	error
	code int
}

// WithCode returns passed err associated with an error code.
func WithCode(err error, code int) error {
	return &coded{
		error: err,
		code:  code,
	}
}

// Returns the string of the error
func (e coded) Error() string {
	return strconv.Itoa(e.code) + ": " + e.error.Error()
}

// Unwrap returns original error.
func (e coded) Unwrap() error {
	return e.error
}

func (e coded) Code() int {
	return e.code
}

// Temporary error
type temporary struct {
	error
}

// AsTemporary returns passed err annotated as being temporary.
func AsTemporary(err error) error {
	return &temporary{
		error: err,
	}
}

// Returns the string of the error
func (e temporary) Error() string {
	return e.error.Error() + "[temporary]"
}

// Unwrap returns original error.
func (e temporary) Unwrap() error {
	return e.error
}

func (e temporary) IsTemporary() bool {
	return true
}