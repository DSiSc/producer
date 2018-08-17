package timer

import (
	"testing"
)

var timer uint64 = 10

func Test_NewTimerProducer(t *testing.T) {
	producer, err := NewTimerProducer(nil, timer)
	if nil != producer && err == nil {
		t.Log("PASS: success to create a producer.")
	} else {
		t.Error("UNPASS: failed to create a producer.")
	}

	if timer == producer.time {
		t.Log("PASS: success to get time of producer.")
	} else {
		t.Error("UNPASS: failed to get time of producer.")
	}
}

func Test_MakeBlock(t *testing.T) {
	producer, _ := NewTimerProducer(nil, timer)
	go func() {
		err := producer.MakeBlock()
		if err != nil{
			t.Log("UNPASS: failed to make a block.")
		} else {
			t.Log("PASS: success to make a block.")
		}
	}()
	producer.StopMakeBlock()
	t.Log("PASS: success to stop make block.")
}
