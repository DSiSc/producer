package common

import (
	"github.com/DSiSc/txpool/common"
)

type Header struct {
	Version          uint32
	PrevBlockHash    common.Hash
	TransactionsRoot common.Hash
	BlockRoot        common.Hash
	Timestamp        uint32
	Height           uint32
	SigData          [][]byte
	hash             *common.Hash
}
