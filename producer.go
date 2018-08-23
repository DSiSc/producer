package producer

import (
	"fmt"
	"github.com/DSiSc/craft/types"
	"github.com/DSiSc/ledger"
	"github.com/DSiSc/producer/common"
	"github.com/DSiSc/txpool"
	"github.com/DSiSc/txpool/common/log"
	"time"
)

type Producer struct {
	txpool *txpool.TxPool
	time   uint64
	ledger *ledger.Ledger
}

func NewProducer(pool *txpool.TxPool, ledger *ledger.Ledger) (*Producer, error) {
	return &Producer{
		txpool: pool,
		ledger: ledger,
	}, nil
}

func (self *Producer) MakeBlock() (*types.Block, error) {
	block, err := self.assembleBlock()
	if nil != err {
		log.Error("Assemble block failed.")
		return nil, fmt.Errorf("Assemble block failed.")
	}
	return block, nil
}

func (self *Producer) assembleBlock() (*types.Block, error) {
	txs := self.txpool.GetTxs()
	txHash := make([]types.Hash, 0, len(txs))
	for _, t := range txs {
		txHash = append(txHash, t.Hash())
	}
	txRoot := common.ComputeMerkleRoot(txHash)
	header := &types.Header{
		TxRoot:    txRoot,
		Timestamp: uint32(time.Now().Unix()),
		Height:    1,
	}

	block := &types.Block{
		Header:       header,
		Transactions: txs,
	}

	return block, nil
}
