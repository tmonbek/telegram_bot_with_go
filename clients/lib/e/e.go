package e

import "fmt"

func WrapError(msg string, err error) error {
	return fmt.Errorf("%s, %w", msg, err)
}

func WrapErrorf(msg string, err error) error {
	if err == nil {
		return nil
	}

	return WrapError(msg, err)
}
