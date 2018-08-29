package producer

import (
	"fmt"
	"github.com/DSiSc/blockchain"
	"github.com/DSiSc/craft/types"
	"github.com/DSiSc/txpool"
	"github.com/DSiSc/txpool/log"
	"time"
)

type Producer struct {
	txpool     txpool.TxsPool
	time       uint64
	blockstore *blockchain.BlockChain
}

func NewProducer(pool txpool.TxsPool, blockchain *blockchain.BlockChain) (*Producer, error) {
	return &Producer{
		txpool:     pool,
		blockstore: blockchain,
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
	txRoot := types.ComputeMerkleRoot(txHash)
	header := &types.Header{
		TxRoot:    txRoot,
		Timestamp: uint64(time.Now().Unix()),
		Height:    1,
	}

	block := &types.Block{
		Header:       header,
		Transactions: txs,
	}

	return block, nil
}
