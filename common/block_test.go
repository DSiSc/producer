package common

import (
	"github.com/DSiSc/txpool/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Hash(t *testing.T) {
	assert := assert.New(t)
	var head Header
	block := &Block{
		Header:       &head,
		Transactions: make([]*types.Transaction, 0),
	}
	hash := block.Hash()
	assert.NotNil(hash)
}
