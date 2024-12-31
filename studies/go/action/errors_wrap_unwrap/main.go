package main

import (
	"errors"
	"fmt"
)

func main() {

	pErr := fmt.Errorf("Parent error this is")
	wrappedErr := fmt.Errorf("Error this is %w", pErr)

	unwappedErr := errors.Unwrap(wrappedErr)

	fmt.Println(pErr.Error())
	// -Parent error this is
	fmt.Println(unwappedErr.Error())
	// Parent error this is
	fmt.Println(wrappedErr.Error())
	// Error this is Parent error this is
}
