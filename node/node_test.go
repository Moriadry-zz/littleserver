package node

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	f, err := ParseMessage([]byte(`{"asd":"zzz", "cc" : 1}`))
	if err != nil {
		fmt.Println("Parse error", err)
	} else {
		fmt.Println(f)
	}

}

func TestGet(t *testing.T) {
	f, _ := ParseMessage([]byte(`{"asd":"zzz", "cc" : 1}`))
	fmt.Println(f.Get("sss"))
}

func TestSet(t *testing.T) {
	f := NewMessage()
	f.Set("asd", "cccc")
	fmt.Println(f)
}

func TestMarshal(t *testing.T) {
	f := NewMessage()
	f.Set("dd", "aaa")
	f.Set("cc", 1)
	f1 := NewMessage()
	f1.Set("vv", "aaa")
	f1.Set("gg", 1.1)
	f.Set("f1", f1)
	b, _ := f.ToBytes()
	fmt.Println(string(b))
}
