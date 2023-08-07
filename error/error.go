package error

/////// Interfaces

type Error interface {
	// Get() error
	SetSimple(code string, params ...string)
	SetComplete(code, message, expectedVal string, givenVal any)
}

type Errors interface {
	Get() map[string]error
	GetOne(key string) Error
	AddSimple(key, code string, params ...string)
	AddComplete(key, code, message, expectedVal string, givenVal any)
	Delete(key string)
	Count() int
	KeyExists(key string) bool
	Import(src errors)
}

/////// Local types

type error struct {
	Code        string      `json:"code"`
	Message     string      `json:"message,omnitempty"`
	ExpectedVal string      `json:"expectedVal,omnitempty"`
	GivenVal    interface{} `json:"givenVal,omnitempty"`
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

type errors map[string]error

func (i errors) Get() map[string]error {
	return i
}

func (i errors) GetOne(key string) Error {
	if error, ok := i[key]; ok {
		return error
	}
	return nil
}

func (i errors) AddSimple(key, code string, params ...string) {
	err := error{Code: code}
	if len(params) > 0 {
		err.Message = params[0]
	}
	i[key] = err
}

func (i errors) AddComplete(key, code, message, expectedVal string, givenVal any) {
	i[key] = error{Code: code, Message: message, ExpectedVal: expectedVal, GivenVal: givenVal}
}

func (i errors) Delete(key string) {
	delete(i, key)
}

func (i errors) Count() int {
	return len(i)
}

func (i errors) KeyExists(key string) bool {
	if _, ok := i[key]; ok {
		return true
	}
	return false
}

func (i errors) Import(src errors) {
	for idx, val := range src {
		i[idx] = val
	}
}

func NewError(code string, params ...string) error {
	myError := error{Code: code}
	if len(params) > 0 {
		myError.Message = params[0]
	}
	return myError
}

func NewCompletError(code, message, expectedVal string, givenVal any) error {
	return error{Code: code, Message: message, ExpectedVal: expectedVal, GivenVal: givenVal}
}

func NewErrors(params ...string) errors {
	if len(params) > 2 {
		myError := errors{}
		myError.AddSimple(params[0], params[1], params[2:]...)
		return myError
	}
	return map[string]error{}
}

func NewCompleteErrors(key, code, message, expectedVal string, givenVal any) errors {
	return map[string]error{
		key: {Code: code, Message: message, ExpectedVal: expectedVal, GivenVal: givenVal},
	}
}

func NewErrorsPick(key string, err error) errors {
	return map[string]error{key: err}
}
