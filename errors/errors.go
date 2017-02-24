package errors

import (
	"fmt"
	"strings"
)

type Errors []error

func Errors(i int) Errors {
	return make([]error, i)
}

func (errors *Errors) Error() error {
	return nil
}

func (errors *Errors) Append(err error) Errors {
	return append(errors, err)
}

func (errors *Errors) IsNil() bool {
	for _, err := range *errors {
		if err != nil {
			return false
		}
	}
	return true
}

func (errors *Errors) String() string {
	return fmt.Sprintf("%d errors:\n\t%s", len(errors), strings.Join(errors, "\n\t"))
}
