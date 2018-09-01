package common

import (
	"github.com/DSiSc/craft/types"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

var emptyTx *types.Transaction

func TestNewTransaction(t *testing.T) {
	assert := assert.New(t)
	data := []byte{0xb2, 0x6f, 0x2b, 0x34, 0x2a}
	b := types.Address{
		0xb2, 0x6f, 0x2b, 0x34, 0x2a, 0xab, 0x24, 0xbc, 0xf6, 0x3e,
		0xa2, 0x18, 0xc6, 0xa9, 0x27, 0x4d, 0x30, 0xab, 0x9a, 0x15,
	}
	emptyTx = NewTransaction(
		0,
		b,
		big.NewInt(0), 0, big.NewInt(0),
		data,
		b,
	)
	assert.NotNil(emptyTx)
}

func TestTxHash(t *testing.T) {
	assert := assert.New(t)
	hash := TxHash(emptyTx)
	expect := types.Hash{
		0xee, 0x58, 0xc8, 0x49, 0xe3, 0x4a, 0x63, 0x76, 0x26, 0xc9, 0xbd, 0x11, 0x7c, 0x35, 0xc0, 0x48,
		0xd2, 0xe1, 0xa1, 0x99, 0x4e, 0xf2, 0xd, 0xeb, 0x83, 0x6f, 0x7a, 0x9f, 0x68, 0xf1, 0xc0, 0xd4,
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
