package family

import "github.com/ricardomgoncalves/truphone_ta_go/internal/errors"

var (
	ErrorFamilyUnknown       = errors.New("unknown error")
	ErrorFamilyLocked        = errors.New("resource locked")
	ErrorFamilyNotFound      = errors.New("not found")
	ErrorFamilyAlreadyExists = errors.New("already exists")
)

var (
	ErrorMemberUnknown       = errors.New("unknown error")
	ErrorMemberLocked        = errors.New("resource locked")
	ErrorMemberNotFound      = errors.New("not found")
	ErrorMemberAlreadyExists = errors.New("already exists")
)
