package core

import (
	"github.com/DSiSc/producer/config"
	"github.com/DSiSc/producer/core/timer"
	"github.com/DSiSc/txpool/common/log"
	"github.com/DSiSc/txpool/core"
	"strconv"
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

func NewProducer(txpool *core.TxPool) (Producer, error) {
	var err error
	var producer Producer
	conf := config.New(config.ConfigAbsPath())
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

/*
func main(){
	st := config.ConfigAbsPath()
	NewProducer(nil)
	fmt.Printf("ddd %s.\n", st)
}
*/
