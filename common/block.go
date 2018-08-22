package common

import (
	"github.com/DSiSc/txpool/common"
	"github.com/DSiSc/txpool/types"
)

type Block struct {
	Header       *Header
	Transactions []*types.Transaction
}

func (b *Block) Hash() common.Hash {
	return b.Header.Hash()
}
