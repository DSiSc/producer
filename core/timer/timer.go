// Implemsnt of producer which make block by fixed time
package timer

import (
	"github.com/DSiSc/txpool/core"
	"github.com/producer/common"
)

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

func (self *TimerProducer) MakeBlock() (*common.Block, error) {
	var temp common.Block
	return &temp, nil
}

func (self *TimerProducer) GetTime() uint64 {
	return self.time
}
