package data

type IS map[string]string
type ISIS map[string]IS
type II map[string]interface{}

type Data struct {
	Meta any `json:"meta,omitempty"`
	Data any `json:"data"`
	Ref  any `json:"ref,omitempty"`
}

type Message struct {
	Message string `json:"message"`
}

type Content struct {
	Content any `json:"content"`
}
