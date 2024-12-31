package main

import (
	"fmt"
)

type Status int
type StatusErr struct {
	Status  Status
	Message string
	Err     error
}

func (s StatusErr) Error() string {
	return s.Message
}

func (s StatusErr) Unwrap() error {
	return s.Err
}

func (se StatusErr) Is(target error) bool {
	if se2, ok := target.(StatusErr); ok {
		// return slices.Equal(se.Status)
		return se.Status == se2.Status
	}
	return false
}

func ErrorProne() error {
	return StatusErr{
		Status:  12,
		Message: "Error man",
	}
}

func main() {
	e := ErrorProne()
	k, ei := e.(StatusErr)
	if ei {
		fmt.Println(k.Message, k.Status)
	}

	switch j := e.(type) {
	case StatusErr:
		fmt.Println("Error man type", j.Status)
	}
}
