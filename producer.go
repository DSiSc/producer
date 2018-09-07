package producer

import (
	"bytes"
	"fmt"
	"github.com/DSiSc/blockchain"
	"github.com/DSiSc/craft/log"
	"github.com/DSiSc/craft/types"
	"github.com/DSiSc/producer/common"
	"github.com/DSiSc/producer/tools"
	"github.com/DSiSc/txpool"
	"github.com/DSiSc/validator/tools/account"
	"github.com/DSiSc/validator/tools/signature"
	"github.com/DSiSc/validator/worker"
	"time"
)

type Producer struct {
	txpool txpool.TxsPool
	time   uint64
	// TODO: we support many workers to promote verification speed  in the future
	workers *worker.Worker
	account *account.Account
}

func NewProducer(pool txpool.TxsPool, Account *account.Account) *Producer {
	return &Producer{
		txpool:  pool,
		account: Account,
	}
}

func (self *Producer) MakeBlock() (*types.Block, error) {
	// Get latest block store and state store
	blockStore, ok := blockchain.NewLatestStateBlockChain()
	if nil != ok {
		log.Error("Get NewLatestStateBlockChain failed.")
		return nil, fmt.Errorf("get NewLatestStateBlockChain failed")
	}
	// make block
	block, err := self.assembleBlock(blockStore)
	if nil != err {
		log.Error("Assemble block failed.")
		return nil, fmt.Errorf("assemble block failed")
	}
	// verify block
	err = self.verifyBlock(block, blockStore)
	if nil != err {
		log.Error("The block verified failed.")
		return nil, err
	}
	// sign
	err = self.signBlock(block)
	if nil != err {
		log.Error("Sign block failed.")
		return nil, fmt.Errorf("signature error:%v", err)
	}
	return block, nil
}

func (self *Producer) assembleBlock(blockStore *blockchain.BlockChain) (*types.Block, error) {
	txs := self.txpool.GetTxs()
	txHash := make([]types.Hash, 0, len(txs))
	for _, t := range txs {
		txHash = append(txHash, common.TxHash(t))
	}
	txRoot := tools.ComputeMerkleRoot(txHash)
	currentBlock := blockStore.GetCurrentBlock()
	header := &types.Header{
		TxRoot:        txRoot,
		Coinbase:      self.account.Address,
		PrevBlockHash: common.BlockHash(currentBlock),
		Timestamp:     uint64(time.Now().Unix()),
		Height:        blockStore.GetCurrentBlockHeight() + 1,
		StateRoot:     blockStore.IntermediateRoot(false),
	}
	block := &types.Block{
		Header:       header,
		Transactions: txs,
	}
	block.HeaderHash = common.HeaderHash(block)
	return block, nil
}

func (self *Producer) verifyBlock(block *types.Block, blockStore *blockchain.BlockChain) error {
	// we support num of works to verify the block
	work := worker.NewWorker(blockStore, block)
	// verify the block
	err := work.VerifyBlock()
	if err != nil {
		log.Error("The block verified failed with err %v.", err)
		return err
	}
	log.Info("The block verified success.")
	return nil
}

func (self *Producer) signBlock(block *types.Block) error {
	sign, err := signature.Sign(self.account, block.HeaderHash[:])
	if nil != err {
		return fmt.Errorf("signature error:%v", err)
	}

	notSigned := true
	for _, value := range block.SigData {
		if bytes.Equal(value, sign) {
			notSigned = false
			log.Warn("Duplicate sign")
		}
	}
	if notSigned {
		block.SigData = append(block.SigData, sign)
	}
	return nil
}
