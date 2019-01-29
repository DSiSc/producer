package producer

import (
	"bytes"
	"fmt"
	"github.com/DSiSc/blockchain"
	"github.com/DSiSc/craft/log"
	"github.com/DSiSc/craft/types"
	"github.com/DSiSc/producer/common"
	"github.com/DSiSc/producer/config"
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
	workers          *worker.Worker
	account          account.Account
	enableSignVerify bool
}

func NewProducer(pool txpool.TxsPool, account account.Account, producerConfig config.ProducerConfig) *Producer {
	return &Producer{
		txpool:           pool,
		account:          account,
		enableSignVerify: producerConfig.EnableSignatureVerify,
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
	block.Header.MixDigest = common.HeaderDigest(block.Header)
	// set state root
	block.Header.StateRoot = blockStore.IntermediateRoot(false)
	// sign
	err = self.signBlock(block)
	if nil != err {
		log.Error("Sign block failed.")
		return nil, fmt.Errorf("signature error: %v", err)
	}
	log.Info("Block %d make success", block.Header.Height)
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
	block := &types.Block{
		Header: &types.Header{
			TxRoot:        txRoot,
			CoinBase:      self.account.Address,
			PrevBlockHash: currentBlock.HeaderHash,
			Timestamp:     uint64(time.Now().Unix()),
			Height:        currentBlock.Header.Height + 1,
		},
		Transactions: txs,
	}
	log.Info("Block %d assemble success with %d txs.", block.Header.Height, len(txs))
	return block, nil
}

func (self *Producer) verifyBlock(block *types.Block, blockStore *blockchain.BlockChain) error {
	// we support num of works to verify the block
	work := worker.NewWorker(blockStore, block, self.enableSignVerify)
	// verify the block
	err := work.VerifyBlock()
	if err != nil {
		log.Error("The block %d verified failed with err %v.", block.Header.Height, err)
		return err
	}
	log.Info("The block %d verified success.", block.Header.Height)
	return nil
}

func (self *Producer) signBlock(block *types.Block) error {
	sign, err := signature.Sign(&self.account, block.Header.MixDigest[:])
	if nil != err {
		log.Error("signature error.")
		return err
	}

	notSigned := true
	for _, value := range block.Header.SigData {
		if bytes.Equal(value, sign) {
			notSigned = false
			log.Warn("Duplicate sign")
		}
	}
	if notSigned {
		block.Header.SigData = append(block.Header.SigData, sign)
	}
	return nil
}
