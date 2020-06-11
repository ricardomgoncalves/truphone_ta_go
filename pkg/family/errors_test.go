package family

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstants(t *testing.T) {
	assert.Equal(t, "500: unknown error", ErrorFamilyUnknown.Error())
	assert.Equal(t, "500: resource locked", ErrorFamilyLocked.Error())
	assert.Equal(t, "400: bad request", ErrorFamilyBadRequest.Error())
	assert.Equal(t, "404: not found", ErrorFamilyNotFound.Error())
	assert.Equal(t, "409: already exists", ErrorFamilyAlreadyExists.Error())

	assert.Equal(t, "500: unknown error", ErrorMemberUnknown.Error())
	assert.Equal(t, "500: resource locked", ErrorMemberLocked.Error())
	assert.Equal(t, "404: not found", ErrorMemberNotFound.Error())
	assert.Equal(t, "409: already exists", ErrorMemberAlreadyExists.Error())
}
