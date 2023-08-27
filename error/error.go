package error

// extended error

import "fmt"

type XError struct {
	Code        string `json:"code"`
	Message     string `json:"message,omitempty"`
	ExpectedVal string `json:"expectedVal,omitempty"`
	GivenVal    any    `json:"givenVal,omitempty"`
}

func (i XError) Error() string {
	return fmt.Sprintf("code: %v; message: %v; expected value: %v; given value: %v", i.Code, i.Message, i.ExpectedVal, i.GivenVal)
}
