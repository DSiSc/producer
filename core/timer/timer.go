// Implemsnt of producer which make block by fixed time
package timer

import (
	"github.com/DSiSc/producer/common"
	"github.com/DSiSc/producer/config"
	"github.com/DSiSc/producer/ledger/store"
	//"github.com/DSiSc/producer/ledger/store"
	"fmt"
	txpool "github.com/DSiSc/txpool/common"
	"github.com/DSiSc/txpool/core"
	"time"
)

// symbol to stop producing blocks
var complete chan int

// default version
const Version uint32 = 0

type TimerProducer struct {
	txpool     *core.TxPool
	time       uint64
	blockStore *store.BlockStore
	// blockStore *store.BlockStore
	// leder ledger.Ledger
}

func NewTimerProducer(pool *core.TxPool, interval uint64) (*TimerProducer, error) {
	timerProducer := &TimerProducer{
		txpool: pool,
		time:   interval,
	}

	blockStore, err := store.NewBlockStore(config.DBAbsPath(), false)
	if nil != err && nil == blockStore {
		return timerProducer, fmt.Errorf("New blcok store failed.")
	}
	timerProducer.blockStore = blockStore
	return timerProducer, nil
}

func (self *TimerProducer) GetTime() uint64 {
	return self.time
}

func (self *TimerProducer) Start() error {
	timer := time.NewTicker(time.Duration(self.time) * time.Second)
	for {
		defer timer.Stop()
		select {
		case <-timer.C:
			err := self.produceBlock()
			if err != nil {
				return fmt.Errorf("Timer generate block failed.")
			}
		case <-complete:
			return nil
		}
	}
}

func (self *TimerProducer) Stop() error {
	if complete != nil {
		close(complete)
		complete = nil
	}
	return nil
}

func (self *TimerProducer) produceBlock() error {
	_, err := self.makeBlock()
	if err != nil {
		return fmt.Errorf("makeBlock failed.")
	}
	/*
		err = ledger.DefLedger.AddBlock(block)
		if err != nil {
			return fmt.Errorf("genBlock DefLedgerPid.RequestFuture Height:%d error:%s", block.Header.Height, err)
		}
	*/
	return nil
}

func (self *TimerProducer) makeBlock() (*common.Block, error) {
	txs := self.txpool.GetTxs()
	txHash := make([]txpool.Hash, 0, len(txs))
	for _, t := range txs {
		txHash = append(txHash, t.Hash())
	}
	txRoot := common.ComputeMerkleRoot(txHash)
	header := &common.Header{
		Version:          Version,
		TransactionsRoot: txRoot,
		Timestamp:        uint32(time.Now().Unix()),
		Height:           1,
	}

	block := &common.Block{
		Header:       header,
		Transactions: txs,
	}

	return block, nil
}
