package producer

import (
	"fmt"
	"github.com/DSiSc/ledger"
	"github.com/DSiSc/producer/config"
	"github.com/DSiSc/producer/policy"
	"github.com/DSiSc/txpool"
	"github.com/DSiSc/txpool/common/log"
)

type Producer interface {
	Start() error
	Stop() error
}

const (
	// config section
	PRODUCER_TIMER = "timer"
	PRODUCER_VOTER = "voter"
	// Structure must matching with defination of config/config.json
	ProducerSymbol  = "producer"
	Policy          = "producer.policy"
	PolicyTimer     = "producer.timer"
	PolicyTimerTime = "producer.timer.time"
)

func NewProducer(conf *config.ProducerConf, txpool *txpool.TxPool, ledger *ledger.Ledger) (Producer, error) {
	var err error
	var producer Producer
	producerPolicy := conf.PolicyName
	switch producerPolicy {
	case PRODUCER_TIMER:
		log.Info("Producer policy is timer.")
		time := conf.PolicyContext.Timer
		if 0 >= time {
			log.Error("Time section of timer is invalid.")
			return nil, fmt.Errorf("Time section of timer is invalid.")
		}
		producer, err = policy.NewTimerProducer(time, txpool, ledger)
	default:
		log.Error("Now, we only support timer policy producer.")
	}
	return producer, err
}
