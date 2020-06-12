package verify

import (
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/errors"
	"strconv"
	"strings"
)

func StringLength(s string, min, max int) error {
	if strings.ReplaceAll(s, " ", "") == "" {
		return errors.New("name should have a character different of a space")
	}

	if len(s) < min {
		return errors.New("name should be at least " + strconv.Itoa(min) + " characters")
	}

	if len(s) > max {
		return errors.New("name should be at maximum " + strconv.Itoa(max) + " characters")
	}

	return nil
}
