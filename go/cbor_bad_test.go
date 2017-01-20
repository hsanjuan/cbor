package cbor

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/ugorji/go/codec"
)

type Op interface{}

type State interface{}

type OpT struct {
	St State
}

type StateT struct {
	Msg string
}

func TestWorksNotWithCBOR(t *testing.T) {
	stateToEnc := State(StateT{"hello"})
	stateToDec := State(StateT{"temp"})

	buf := new(bytes.Buffer)
	enc := NewEncoder(buf)
	err := enc.Encode(stateToEnc)
	if err != nil {
		t.Fatal(err)
	}

	dec := NewDecoder(buf)
	err = dec.Decode(&stateToDec)
	if err != nil {
		t.Fatal(err)
	}

	st, ok := stateToDec.(StateT)
	if !ok {
		t.Log(reflect.TypeOf(stateToDec))
		t.Fatal("should cast st")
	}

	if st.Msg != "hello" {
		t.Fatal("decoded Msg should be hello")
	}
	t.Log(st.Msg)
}

func TestWorksWithMsgPack(t *testing.T) {
	stateToEnc := State(StateT{"hello"})
	stateToDec := State(StateT{"temp"})

	buf := new(bytes.Buffer)
	enc := codec.NewEncoder(buf, &codec.MsgpackHandle{})
	err := enc.Encode(stateToEnc)
	if err != nil {
		t.Fatal(err)
	}

	dec := codec.NewDecoder(buf, &codec.MsgpackHandle{})
	err = dec.Decode(&stateToDec)
	if err != nil {
		t.Fatal(err)
	}

	st, ok := stateToDec.(StateT)
	if !ok {
		t.Fatal("should cast st")
	}

	if st.Msg != "hello" {
		t.Fatal("decoded Msg should be hello")
	}
	t.Log(st.Msg)
}
