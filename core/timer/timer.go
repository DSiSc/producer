// Implemsnt of producer which make block by fixed time
package timer

import (
	// "github.com/DSiSc/producer/common"
	"github.com/DSiSc/txpool/common/log"
	"github.com/DSiSc/txpool/core"
	"time"
)

// symbol to stop producing blocks
var complete chan int

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

func (self *TimerProducer) MakeBlock() interface{} {
	timer := time.NewTicker(time.Duration(self.time) * time.Second)
	for {
		defer timer.Stop()
		select {
		case <-timer.C:
			err := self.genBlock()
			if err != nil {
				log.Error("Timer generate block failed.")
				return err
			}
		case <-complete:
			return nil
		}
	}
}

func (self *TimerProducer) genBlock() error {
	log.Info("Timer to generate block.")
	return nil
}

func (self *TimerProducer) StopMakeBlock() {
	if complete != nil {
		close(complete)
		complete = nil
	}
}
