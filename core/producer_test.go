package core

import (
	"testing"
)

func Test_NewProducer(t *testing.T) {
	producer, err := NewProducer(nil)
	if nil != producer && nil == err {
		t.Log("PASS: success to create a producer.")
	} else {
		t.Error("UNPASS: failed to create a producer.")
	}
	block, ok := producer.MakeBlock()
	if nil != block && nil == ok {
		t.Log("PASS: success to make a block.")
	} else {
		t.Error("UNPASS: failed to make a block.")
	}
}
