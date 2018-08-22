package producer

import (
	"github.com/DSiSc/producer/config"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func mock_NewProducerConf(time uint, num uint) config.ProducerConf {
	return config.ProducerConf{
		PolicyName: "timer",
		PolicyContext: config.ProducerPolicy{
			Timer: uint64(time),
			Num:   uint64(num),
		},
	}
}

func Test_NewProducer(t *testing.T) {
	assert := assert.New(t)
	conf := mock_NewProducerConf(10, 0)

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
