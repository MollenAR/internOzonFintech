package errorTypes

import (
	"fmt"
)

type ErrTryAgainLater struct {
	Reason string
}

func (e ErrTryAgainLater)Error() string {
	return fmt.Sprintf("Reason: %s", e.Reason)
}

// -----------------------------------------------------------

type ErrWrongOriginalUrl struct {
	Reason string
}

func (e ErrWrongOriginalUrl)Error() string {
	return fmt.Sprintf("Reason: %s", e.Reason)
}

// -----------------------------------------------------------

type ErrWrongShortUrl struct {
	Reason string
}

func (e ErrWrongShortUrl)Error() string {
	return fmt.Sprintf("Reason: %s", e.Reason)
}

// -----------------------------------------------------------

type ErrWrongUsage struct {
	Reason string
}

func (e ErrWrongUsage)Error() string {
	return fmt.Sprintf("Reason: %s", e.Reason)
}
