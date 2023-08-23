package error

/////// Interfaces

type Error interface {
	Get() error
	SetSimple(code string, params ...string)
	SetComplete(code, message, expectedVal string, givenVal any)
}

type Errors interface {
	Get() map[string]Error
	GetOne(key string) Error
	GetFirst() Error
	AddSimple(key, code string, params ...string)
	AddComplete(key, code, message, expectedVal string, givenVal any)
	Pick(key string, err Error)
	Delete(key string)
	Count() int
	KeyExists(key string) bool
	Import(src errors)
}

/////// Local types

type error struct {
	Code        string      `json:"code"`
	Message     string      `json:"message,omitempty"`
	ExpectedVal string      `json:"expectedVal,omitempty"`
	GivenVal    interface{} `json:"givenVal,omitempty"`
}

func (i error) Get() error {
	return i
}

func (i error) SetSimple(code string, params ...string) {
	i.Code = code
	if len(params) > 0 {
		i.Message = params[0]
	}
}

func (i error) SetComplete(code, message, expectedVal string, givenVal any) {
	i.Code = code
	i.Message = message
	i.ExpectedVal = expectedVal
	i.GivenVal = givenVal
}

type errors map[string]Error

// Get the errors object
func (i errors) Get() map[string]Error {
	return i
}

// Get one error by key
func (i errors) GetOne(key string) Error {
	if error, ok := i[key]; ok {
		return error
	}
	return nil
}

func (i errors) GetFirst() Error {
	for _, err := range i {
		return err
	}
	return nil
}

// Add an error by using simple parameters (key, code, [detail message])
func (i errors) AddSimple(key, code string, params ...string) {
	err := error{Code: code}
	if len(params) > 0 {
		err.Message = params[0]
	}
	i[key] = err
}

// Add an error by using complete parameters (key, code, detail message, expected value, given value)
func (i errors) AddComplete(key, code, message, expectedVal string, givenVal any) {
	i[key] = error{Code: code, Message: message, ExpectedVal: expectedVal, GivenVal: givenVal}
}

// Add an error by using an existing error (key, error)
func (i errors) Pick(key string, err Error) {
	i[key] = err
}

// Delete an error by key
func (i errors) Delete(key string) {
	delete(i, key)
}

// Get total error count
func (i errors) Count() int {
	return len(i)
}

// Check if a key exists
func (i errors) KeyExists(key string) bool {
	if _, ok := i[key]; ok {
		return true
	}
	return false
}

// Import existing errors
func (i errors) Import(src errors) {
	for idx, val := range src {
		i[idx] = val
	}
}

// Create instance of an error by using simple paramters (code, [message detail])
func NewError(code string, params ...string) error {
	myError := error{Code: code}
	if len(params) > 0 {
		myError.Message = params[0]
	}
	return myError
}

// Create instance of an error by using complete paramters (code, [message detail])
func NewCompletError(code, message, expectedVal string, givenVal any) error {
	return error{Code: code, Message: message, ExpectedVal: expectedVal, GivenVal: givenVal}
}

// Create instance of errors by using simple paramters ([key, code, message detail])
// Wihtout parameters rerturns 0 count of error
// With parameters returns 1 count of error
func NewErrors(params ...string) errors {
	if len(params) > 2 {
		myError := errors{}
		myError.AddSimple(params[0], params[1], params[2:]...)
		return myError
	}
	return errors{}
}

// Create instance of errors by using complete paramters (key, code, message detail, expected value, given value)
// Return 1 count of error
func NewCompleteErrors(key, code, message, expectedVal string, givenVal any) errors {
	return map[string]Error{
		key: error{Code: code, Message: message, ExpectedVal: expectedVal, GivenVal: givenVal},
	}
}

// Create instance of errors by using an existing error (key, error)
func NewErrorsPick(key string, err Error) errors {
	return errors{key: err}
}
