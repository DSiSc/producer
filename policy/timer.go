// Implemsnt of producer which make block by fixed time
package policy

import (
	"fmt"
	"github.com/DSiSc/ledger"
	"github.com/DSiSc/producer/common"
	"github.com/DSiSc/txpool"
	types "github.com/DSiSc/txpool/common"
	"github.com/DSiSc/txpool/common/log"
	"time"
)

// symbol to stop producing blocks
var complete chan int

// default version
const Version uint32 = 0

type TimerProducer struct {
	txpool *txpool.TxPool
	time   uint64
	ledger *ledger.Ledger
}

func NewTimerProducer(interval uint64, pool *txpool.TxPool, ledger *ledger.Ledger) (*TimerProducer, error) {
	timerProducer := &TimerProducer{
		txpool: pool,
		time:   interval,
	}
	timerProducer.ledger = ledger
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
	block, err := self.makeBlock()
	if err != nil {
		log.Error("makeBlock failed.")
		return fmt.Errorf("makeBlock failed.")
	}

	err = self.ledger.BlockStore.SaveBlock(block)
	if err != nil {
		log.Error("save block failed.")
		return fmt.Errorf("genBlock DefLedgerPid.RequestFuture Height:%d error:%s", block.Header.Height, err)
	}

	return nil
}

func (self *TimerProducer) makeBlock() (*common.Block, error) {
	txs := self.txpool.GetTxs()
	txHash := make([]types.Hash, 0, len(txs))
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
