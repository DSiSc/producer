package policy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var timer uint64 = 10

func Test_NewTimerProducer(t *testing.T) {
	assert := assert.New(t)
	producer, err := NewTimerProducer(timer, nil, nil)
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

func Test_GetTime(t *testing.T) {
	assert := assert.New(t)
	producer, err := NewTimerProducer(timer, nil, nil)
	assert.Nil(err)
	time := producer.GetTime()
	assert.Equal(time, timer, "they should not be equal")
}
