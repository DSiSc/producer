package producer

import (
	"github.com/DSiSc/producer/config"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func mock_NewProducerConfig(time uint, num uint) config.ProducerConfig {
	return config.ProducerConfig{
		PolicyName: "timer",
		PolicyContext: config.ProducerPolicy{
			Timer: uint64(time),
			Num:   uint64(num),
		},
	}
}

func Test_NewProducer(t *testing.T) {
	assert := assert.New(t)
	conf := mock_NewProducerConfig(10, 0)

	producer, err := NewProducer(&conf, nil, nil)
	assert.Nil(err)

	p := reflect.TypeOf(producer)
	method, exist := p.MethodByName("Start")
	assert.NotNil(method)
	assert.True(exist)

	method, exist = p.MethodByName("Stop")
	assert.NotNil(method)
	assert.True(exist)

}
