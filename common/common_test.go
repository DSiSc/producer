package common

import (
	"github.com/DSiSc/craft/types"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

var to = &types.Address{
	0xb2, 0x6f, 0x2b, 0x34, 0x2a, 0xab, 0x24, 0xbc, 0xf6, 0x3e,
	0xa2, 0x18, 0xc6, 0xa9, 0x27, 0x4d, 0x30, 0xab, 0x9a, 0x15,
}
var from = &types.Address{
	0x12, 0x6f, 0x2b, 0x34, 0x2a, 0xab, 0x24, 0xbc, 0xf6, 0x3e,
	0xa2, 0x18, 0xc6, 0xa9, 0x27, 0x4d, 0x30, 0xab, 0x9a, 0x15,
}

var data = []byte{
	0xb2, 0x6f, 0x2b, 0x34, 0x2a, 0xab, 0x24, 0xbc, 0xf6, 0x3e,
}

var emptyTx = &types.Transaction{

	Data: types.TxData{
		AccountNonce: 0,
		Price:        new(big.Int).SetUint64(10),
		GasLimit:     100,
		Recipient:    to,
		From:         from,
		Amount:       new(big.Int).SetUint64(50),
		Payload:      data,
	},
}

func TestTxHash(t *testing.T) {
	assert := assert.New(t)
	hash := TxHash(emptyTx)
	expect := types.Hash{
		0x1d, 0xcf, 0x7, 0xba, 0xfc, 0x42, 0xb0, 0x8d, 0xfd, 0x23, 0x9c, 0x45, 0xa4, 0xb9, 0x38, 0xd,
		0x8d, 0xfe, 0x5d, 0x6f, 0xa7, 0xdb, 0xd5, 0x50, 0xc9, 0x25, 0xb1, 0xb3, 0x4, 0xdc, 0xc5, 0x1c,
	}
	assert.Equal(expect, hash)
}

func TestSum(t *testing.T) {
	b := []byte{
		0xb2, 0x6f, 0x2b, 0x34, 0x2a, 0xab, 0x24, 0xbc, 0xf6, 0x3e,
		0xa2, 0x18, 0xc6, 0xa9, 0x27, 0x4d, 0x30, 0xab, 0x9a, 0x15,
	}
	c := Sum(b)
	except := []byte{
		0x5b, 0x49, 0xff, 0xa2, 0x47, 0x5e, 0x11, 0x40, 0x92, 0xb,
		0x4, 0x6, 0x43, 0xc5, 0x7f, 0x8e, 0x5b, 0xb1, 0x86, 0x4, 0xeb,
		0xcc, 0x62, 0x43, 0x71, 0x39, 0x76, 0xfd, 0xa4, 0x8b, 0xc3, 0x7,
	}
	assert.Equal(t, except, c)
}

func TestCopyBytes(t *testing.T) {
	b := []byte{
		0xb2, 0x6f, 0x2b, 0x34, 0x2a, 0xab, 0x24, 0xbc, 0xf6, 0x3e,
		0xa2, 0x18, 0xc6, 0xa9, 0x27, 0x4d, 0x30, 0xab, 0x9a, 0x15,
	}
	c := CopyBytes(b)
	assert.Equal(t, b, c)
}