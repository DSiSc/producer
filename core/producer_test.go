package core

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_NewProducer(t *testing.T) {
	assert := assert.New(t)

	producer, err := NewProducer(nil)
	assert.Nil(err)

	p := reflect.TypeOf(producer)
	method, exist := p.MethodByName("Start")
	assert.NotNil(method)
	assert.True(exist)

	method, exist = p.MethodByName("Stop")
	assert.NotNil(method)
	assert.True(exist)

}
