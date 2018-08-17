package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetConfigItem(t *testing.T) {
	assert := assert.New(t)
	config := New(ConfigName)
	assert.NotNil(&config)

	parsedStruct := config.GetConfigItem("producer")
	assert.NotNil(parsedStruct)

	pasedItem := config.GetConfigItem("producer.policy")
	assert.NotNil(pasedItem)
	assert.Equal("timer", pasedItem.(string), "they should be equal")

	pasedItem = config.GetConfigItem("producer.pol")
	assert.Nil(pasedItem)
}
