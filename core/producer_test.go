package core

import (
	"reflect"
	"testing"
)

func Test_NewProducer(t *testing.T) {
	producer, err := NewProducer(nil)
	if nil != producer && nil == err {
		t.Log("PASS: success to create a producer.")
	} else {
		t.Error("UNPASS: failed to create a producer.")
	}

	p := reflect.TypeOf(producer)
	method, exist := p.MethodByName("Start")
	if nil != &method && exist {
		t.Log("PASS: success to create a producer with method Start().")
	} else {
		t.Error("UNPASS: failed to create a producer with method Start().")
	}

	method, exist = p.MethodByName("Stop")
	if nil != &method && exist {
		t.Log("PASS: success to create a producer with method Stop().")
	} else {
		t.Error("UNPASS: failed to create a producer with method Stop().")
	}
}
