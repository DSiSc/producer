// Implemsnt of producer which make block by fixed time
package timer

import (
	"github.com/DSiSc/producer/common"
	txpool "github.com/DSiSc/txpool/common"
	"github.com/DSiSc/txpool/common/log"
	"github.com/DSiSc/txpool/core"
	"time"
)

// symbol to stop producing blocks
var complete chan int

// default version
const Version uint32 = 0

type TimerProducer struct {
	txpool *core.TxPool
	time   uint64
}

func NewTimerProducer(pool *core.TxPool, interval uint64) (*TimerProducer, error) {
	timerProducer := &TimerProducer{
		txpool: pool,
		time:   interval,
	}
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
			_, err := self.makeBlock()
			if err != nil {
				log.Error("Timer generate block failed.")
				return err
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
