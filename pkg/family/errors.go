package family

import "github.com/ricardomgoncalves/truphone_ta_go/internal/errors"

var (
	ErrorFamilyUnknown       = errors.WithCode(errors.New("unknown error"), 500)
	ErrorFamilyLocked        = errors.WithCode(errors.New("resource locked"), 500)
	ErrorFamilyNotFound      = errors.WithCode(errors.New("not found"), 404)
	ErrorFamilyAlreadyExists = errors.WithCode(errors.New("already exists"), 404)
)

var (
	ErrorMemberUnknown       = errors.WithCode(errors.New("unknown error"), 500)
	ErrorMemberLocked        = errors.WithCode(errors.New("resource locked"), 500)
	ErrorMemberNotFound      = errors.WithCode(errors.New("not found"), 404)
	ErrorMemberAlreadyExists = errors.WithCode(errors.New("already exists"), 404)
)
