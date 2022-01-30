package validation

import (
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/pkg/errors"
	"regexp"
)

const maxLength = 2048
const minLength = 2

func ValidateOriginalUrl(originalUrl string) error {
	err := validation.Validate(originalUrl, validation.Required, validation.Length(minLength, maxLength), is.URL)
	if err != nil {
		return errors.Wrap(err, "")
	}

	return nil
}

func ValidateShortUrl(shortUrls string) error {
	err := validation.Validate(shortUrls, validation.Required, validation.Match(regexp.MustCompile("^[a-zA-Z0-9_]{10}$")))
	if err != nil {
		return errors.Wrap(err, "")
	}

	return nil
}
