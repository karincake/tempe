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

func (i XError) SetSimple(code string, params ...string) {
	i.Code = code
	if len(params) > 0 {
		i.Message = params[0]
	}
}

func (i XError) SetComplete(code, message, expectedVal string, givenVal any) {
	i.Code = code
	i.Message = message
	i.ExpectedVal = expectedVal
	i.GivenVal = givenVal
}
