package producer

import (
	"github.com/DSiSc/txpool"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewProducer(t *testing.T) {
	assert := assert.New(t)

	producer := NewProducer(nil, nil)
	assert.Nil(producer)
}

func Test_MakeBlock(t *testing.T) {
	assert := assert.New(t)
	producer := NewProducer(nil, nil)
	assert.Panics(func() { producer.MakeBlock() })

	tx := txpool.NewTxPool(txpool.DefaultTxPoolConfig)
	producer = NewProducer(tx, nil)
	assert.Panics(func() { producer.MakeBlock() })
	/*
		block, err := producer.MakeBlock()
		assert.Nil(err)
		assert.NotNil(block)
	*/
}

func Test_assembleBlock(t *testing.T) {
	assert := assert.New(t)
	producer := NewProducer(nil, nil)
	assert.Panics(func() { producer.assembleBlock() })
	/*
		tx := txpool.NewTxPool(txpool.DefaultTxPoolConfig)
		producer = NewProducer(tx, nil)
		block, err := producer.assembleBlock()
		assert.Nil(err)
		assert.NotNil(block)
	*/
}
