package structs

import (
	"errors"
	"fmt"
	"strings"
)

type Errors struct {
	errors []error
}

func Errors() Errors {
	return Errors{make([]error, 0)}
}

func (e *Errors) Error() error {
	if len(e.errors) == 0 {
		return nil
	} else {
		return errors.New(e.String())
	}
}

func (e *Errors) Append(errs ...error) {
	for _, err := range errs {
		if err != nil {
			e.errors = append(e.errors, err)
		}
	}
}

func (e *Errors) String() string {
	return fmt.Sprintf("%d errors:\n\t%s", len(e.errors), strings.Join(e.errors, "\n\t"))
}
