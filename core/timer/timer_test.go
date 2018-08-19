package timer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var timer uint64 = 10

func Test_NewTimerProducer(t *testing.T) {
	assert := assert.New(t)
	producer, err := NewTimerProducer(nil, timer)
	assert.NotNil(producer)
	assert.Nil(err)
	assert.Equal(timer, producer.time, "they should not be equal")
}

func Test_Start(t *testing.T) {
	assert := assert.New(t)
	producer := &TimerProducer{
		txpool: nil,
		time:   timer,
		ledger: nil,
	}
	go func() {
		err := producer.Start()
		assert.Nil(err)
	}()

	err := producer.Stop()
	assert.Nil(err)
}
