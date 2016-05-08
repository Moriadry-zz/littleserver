package node

import (
	"encoding/json"
)

type Message map[string]interface{}

func NewMessage() Message {
	return make(map[string]interface{})
}

func ParseMessage(buf []byte) (Message, err) {
	var f interface{}
	err := json.Unmarshal(buf, &f)
	return f.(map[string]interface{}), err
}

func (m Message) Get(key string) interface{} {
	return m[key]
}

func (m Message) Set(key string, value interface{}) {
	m[key] = value
}

func (m Message) ToBytes() ([]byte, err) {
	return json.Marshal(m)
}
