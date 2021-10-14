package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Fooer interface {
	Foo()
}

type A struct{ Field string }

func (a *A) Foo() {}

type B struct{ Field float64 }

func (b *B) Foo() {}

type Wrapper struct {
	Type  string
	Value Fooer
}

func (w Wrapper) MarshalJSON() ([]byte, error) {
	switch w.Value.(type) {
	case *A:
		w.Type = "A"
	case *B:
		w.Type = "B"
	default:
		return nil, fmt.Errorf("invalid type: %T", w.Value)
	}
	type W Wrapper
	return json.Marshal(W(w))
}

func (w *Wrapper) UnmarshalJSON(b []byte) error {
	var W struct {
		Type  string
		Value json.RawMessage
	}
	if err := json.Unmarshal(b, &W); err != nil {
		return err
	}
	var value interface{}
	switch W.Type {
	case "A":
		value = new(A)
	case "B":
		value = new(B)
	default:
		return fmt.Errorf("invalid type: %s", W.Type)
	}
	if err := json.Unmarshal(W.Value, &value); err != nil {
		return err
	}
	w.Type, w.Value = W.Type, value.(Fooer)
	return nil
}

func main() {
	d := json.NewEncoder(os.Stdout)
	if err := d.Encode(Wrapper{Value: &A{"Value"}}); err != nil {
		log.Fatal(err)
	}
	if err := d.Encode(Wrapper{Value: &B{0.5}}); err != nil {
		log.Fatal(err)
	}
	var w Wrapper
	if err := json.Unmarshal([]byte(`{"Type":"A","Value":{"Field":"Value"}}`), &w); err != nil {
		log.Fatal(err)
	}
	log.Printf("Type: %s Value: %#v", w.Type, w.Value)
	if err := json.Unmarshal([]byte(`{"Type":"B","Value":{"Field":0.5}}`), &w); err != nil {
		log.Fatal(err)
	}
	log.Printf("Type: %s Value: %#v", w.Type, w.Value)
	err := json.Unmarshal([]byte(`{"Type":"X","Value":{"Field":true}}`), &w)
	if err == nil {
		log.Fatal("Error expected")
	}
	log.Printf("Got error: %s ", err)

}
