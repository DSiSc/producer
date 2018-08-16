package core

import (
	"github.com/DSiSc/txpool/common/log"
	"github.com/DSiSc/txpool/core"
	"github.com/DSiSc/producer/common"
)

type Producer struct {
	txpool *core.TxPool
}

// define constant infomation
var (
	// config section
	// Structure must matching with defination of config/config.json
	ProducerSymbol     = "producer"
	ProducerPolicyPath = "producer.policy"
)

func NewProducer(pool *core.TxPool) *Producer {
	log.Info("Create a producer.")
	return &Producer{txpool: pool}
}

func (p *Producer) MakeBlock() (*common.Block, error) {
	return nil, nil
}
