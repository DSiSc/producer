package producer

import (
	"bytes"
	"fmt"
	"github.com/DSiSc/blockchain"
	"github.com/DSiSc/craft/types"
	"github.com/DSiSc/producer/common"
	"github.com/DSiSc/producer/tools"
	"github.com/DSiSc/txpool"
	"github.com/DSiSc/txpool/log"
	"github.com/DSiSc/validator/tools/account"
	"github.com/DSiSc/validator/tools/signature"
	"github.com/DSiSc/validator/worker"
	"time"
)

type Producer struct {
	txpool     txpool.TxsPool
	time       uint64
	blockstore *blockchain.BlockChain
	// TODO: we support many workers to promote verification speed  in the future
	workers *worker.Worker
	account *account.Account
}

func NewProducer(pool txpool.TxsPool, Account *account.Account) *Producer {
	blockstore, err := blockchain.NewLatestStateBlockChain()
	if nil != err {
		log.Error("Get latest state block failed.")
		return nil
	}
	return &Producer{
		txpool:     pool,
		blockstore: blockstore,
		account:    Account,
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
	currentBlock := self.blockstore.GetCurrentBlock()
	header := &types.Header{
		TxRoot:        txRoot,
		Coinbase:      self.account.Address,
		PrevBlockHash: common.BlockHash(currentBlock),
		Timestamp:     uint64(time.Now().Unix()),
		Height:        self.blockstore.GetCurrentBlockHeight() + 1,
		StateRoot:     self.blockstore.IntermediateRoot(false),
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
	sign, err := signature.Sign(self.account, hash[:])
	if nil != err {
		return fmt.Errorf("[Signature],Sign error:%s.", err)
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
