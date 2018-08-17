package common

import (
	"github.com/DSiSc/txpool/core/types"
)

type Block struct {
	Header       *Header
	Transactions []*types.Transaction
}
