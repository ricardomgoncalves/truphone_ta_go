package family

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstants(t *testing.T) {
	assert.Equal(t, "unknown error", ErrorFamilyUnknown.Error())
	assert.Equal(t, "resource locked", ErrorFamilyLocked.Error())
	assert.Equal(t, "not found", ErrorFamilyNotFound.Error())
	assert.Equal(t, "already exists", ErrorFamilyAlreadyExists.Error())

	assert.Equal(t, "unknown error", ErrorMemberUnknown.Error())
	assert.Equal(t, "resource locked", ErrorMemberLocked.Error())
	assert.Equal(t, "not found", ErrorMemberNotFound.Error())
	assert.Equal(t, "already exists", ErrorMemberAlreadyExists.Error())
}
