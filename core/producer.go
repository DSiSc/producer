package core

import (
	"github.com/DSiSc/txpool/common/log"
	"github.com/DSiSc/txpool/core"
	"github.com/producer/common"
	"github.com/producer/config"
	"github.com/producer/core/timer"
	"strconv"
)

type Producer interface {
	MakeBlock() (*common.Block, error)
}

const (
	// config section
	PRODUCER_TIMER = "timer"
	PRODUCER_VOTER = "voter"
	// Structure must matching with defination of config/config.json
	ProducerSymbol     = "producer"
	Policy             = "producer.policy"
	PolicyTimer        = "producer.timer"
	PolicyTimerTime    = "producer.timer.time"
	ProducerConfigPath = "./../config/config.json"
)

func NewProducer(txpool *core.TxPool) (Producer, error) {
	var err error
	var producer Producer
	conf := config.New(ProducerConfigPath)
	policy := conf.GetConfigItem(Policy).(string)
	switch policy {
	case PRODUCER_TIMER:
		log.Info("Get timer policy producer.")
		time, err := strconv.ParseUint(conf.GetConfigItem(PolicyTimerTime).(string), 10, 64)
		if nil != err {
			log.Error("Get time section for timer producer failed.")
			return nil, err
		}
		producer, err = timer.NewTimerProducer(txpool, time)
	default:
		log.Error("Now, we only support timer policy producer.")
	}
	return producer, err
}
