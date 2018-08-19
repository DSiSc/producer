package store

import (
	"github.com/DSiSc/producer/common"
)

type Ledger interface {
	AddBlock(block *common.Block) error
}
