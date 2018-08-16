package core

import (
	"github.com/DSiSc/txpool/core"
	"testing"
)

func Test_NewProducer(t *testing.T) {
	txpool := core.NewTxPool(core.DefaultTxPoolConfig)
	if nil == txpool {
		t.Error("UNPASS: failed to create a txpool")
	}
	producer := NewProducer(txpool)
	if nil != producer {
		t.Log("PASS: success to create a producer.")
	} else {
		t.Error("UNPASS: failed to create a producer.")
	}
}

func Test_MakeBlock(t *testing.T) {
	producer := NewProducer(nil)
	newBlock, ok := producer.MakeBlock()
	if nil == newBlock && nil == ok {
		t.Log("PASS: success to make a block.")
	} else {
		t.Error("UNPASS: failed to make a block.")
	}
}
