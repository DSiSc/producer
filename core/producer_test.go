package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewProducer(t *testing.T) {
	assert := assert.New(t)

	_, err := NewProducer(nil)
	//assert.NotNil(producer)
	assert.Nil(err)
	/*
		p := reflect.TypeOf(producer)
		method, exist := p.MethodByName("Start")
		assert.NotNil(method)
		assert.True(exist)

		method, exist = p.MethodByName("Stop")
		assert.NotNil(method)
		assert.True(exist)
	*/
}
