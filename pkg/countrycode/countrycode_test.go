package countrycode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValid(t *testing.T) {
	t.Run("should return valid country code", func(t *testing.T) {
		countryCodes := []string{
			"PT",
			"GB",
			"UK",
			"US",
			"ES",
			"FR",
			"GR",
		}

		for i := 0; i < len(countryCodes); i++ {
			assert.True(t, IsValid(countryCodes[i]), "should had return true for index:", i)
		}
	})

	t.Run("should return invalid for country code", func(t *testing.T) {
		countryCodes := []string{
			"",
			"ASD",
			"123",
			"EN",
		}

		for i := 0; i < len(countryCodes); i++ {
			assert.False(t, IsValid(countryCodes[i]), "should had return false for index:", i)
		}
	})
}
