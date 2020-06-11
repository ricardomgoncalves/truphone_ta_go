package countrycode

import (
	countrycodes "github.com/launchdarkly/go-country-codes"
)

func IsValid(countryCode string) bool {
	_, has := countrycodes.GetByAlpha2(countryCode)
	return has
}
