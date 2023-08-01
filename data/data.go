package data

type IS map[string]string
type ISIS map[string]IS
type II map[string]interface{}

type Data struct {
	Meta any `json:"meta,omnitempty"`
	Data any `json:"data"`
	Ref  any `json:"ref,omnitempty"`
}

type Message struct {
	Message string `json:"message"`
}

type Content struct {
	Content any `json:"content"`
}
