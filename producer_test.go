package producer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewProducer(t *testing.T) {
	assert := assert.New(t)

	producer, err := NewProducer(nil, nil)
	assert.Nil(err)
	assert.NotNil(producer)
}

func Test_MakeBlock(t *testing.T) {
	assert := assert.New(t)
	producer, _ := NewProducer(nil, nil)
	assert.Panics(func() { producer.MakeBlock() })
}

func Test_assembleBlock(t *testing.T) {
	assert := assert.New(t)
	producer, _ := NewProducer(nil, nil)
	assert.Panics(func() { producer.assembleBlock() })
}
