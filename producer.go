package producer

import (
	"fmt"
	"github.com/DSiSc/blockchain"
	"github.com/DSiSc/craft/types"
	"github.com/DSiSc/producer/common"
	"github.com/DSiSc/producer/tools"
	"github.com/DSiSc/producer/tools/account"
	"github.com/DSiSc/producer/tools/signature"
	"github.com/DSiSc/txpool"
	"github.com/DSiSc/txpool/log"
	"github.com/DSiSc/validator/worker"
	"time"
)

type Producer struct {
	txpool     txpool.TxsPool
	time       uint64
	blockstore *blockchain.BlockChain
	// TODO: we support many workers to promote verification speed  in the future
	workers *worker.Worker
	Account *account.Account
}

func NewProducer(pool txpool.TxsPool, blockchain *blockchain.BlockChain) *Producer {
	return &Producer{
		txpool:     pool,
		blockstore: blockchain,
	}
}

func (self *Producer) MakeBlock() (*types.Block, error) {
	// make block
	block, err := self.assembleBlock()
	if nil != err {
		log.Error("Assemble block failed.")
		return nil, fmt.Errorf("Assemble block failed.")
	}
	// verify block
	err = self.verifyBlock(block)
	if nil != err {
		log.Error("The block verified failed.")
		return nil, err
	}
	// sign
	err = self.signBlock(block)
	if nil != err {
		log.Error("Sign block failed.")
		return nil, fmt.Errorf("[Signature],Sign error:%s.", err)
	}
	return block, nil
}

func (self *Producer) assembleBlock() (*types.Block, error) {
	txs := self.txpool.GetTxs()
	txHash := make([]types.Hash, 0, len(txs))
	for _, t := range txs {
		txHash = append(txHash, common.TxHash(t))
	}
	txRoot := tools.ComputeMerkleRoot(txHash)
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

func (self *Producer) verifyBlock(block *types.Block) error {
	// we support num of works to verify the block
	work := worker.NewWorker(self.blockstore, block)
	// verify the block
	err := work.VerifyBlock()
	if err != nil {
		log.Error("The block verified failed.")
		return err
	}
	log.Info("The block verified success.")
	return nil
}

func (self *Producer) signBlock(block *types.Block) error {
	hash := common.BlockHash(block)
	sig, err1 := signature.Sign(self.Account, hash[:])
	if nil != err1 {
		return fmt.Errorf("[Signature],Sign error:%s.", err1)
	}
	block.Header.SigData = append(block.Header.SigData, sig)
	return nil
}
