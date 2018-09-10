package producer

import (
	"fmt"
	"github.com/DSiSc/blockchain"
	"github.com/DSiSc/blockchain/config"
	"github.com/DSiSc/craft/types"
	"github.com/DSiSc/monkey"
	"github.com/DSiSc/txpool"
	"github.com/DSiSc/validator/tools"
	account2 "github.com/DSiSc/validator/tools/account"
	"github.com/DSiSc/validator/tools/signature"
	"github.com/DSiSc/validator/worker"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

var mockAddress = types.Address{
	0x33, 0x3c, 0x33, 0x10, 0x82, 0x4b, 0x7c, 0x68, 0x51, 0x33,
	0xf2, 0xbe, 0xdb, 0x2c, 0xa4, 0xb8, 0xb4, 0xdf, 0x63, 0x3d,
}

type eventCenter struct {
}

// subscriber subscribe specified eventType with eventFunc
func (*eventCenter) Subscribe(eventType types.EventType, eventFunc types.EventFunc) types.Subscriber {
	return nil
}

// subscriber unsubscribe specified eventType
func (*eventCenter) UnSubscribe(eventType types.EventType, subscriber types.Subscriber) (err error) {
	return nil
}

// notify subscriber of eventType
func (*eventCenter) Notify(eventType types.EventType, value interface{}) (err error) {
	return nil
}

// notify specified eventFunc
func (*eventCenter) NotifySubscriber(eventFunc types.EventFunc, value interface{}) {

}

// notify subscriber traversing all events
func (*eventCenter) NotifyAll() (errs []error) {
	return nil
}

// unsubscrible all event
func (*eventCenter) UnSubscribeAll() {
}

func TestNewProducer(t *testing.T) {
	assert := assert.New(t)
	txpool := txpool.NewTxPool(txpool.DefaultTxPoolConfig)
	account := &account2.Account{
		Address: tools.HexToAddress("333c3310824b7c685133f2bedb2ca4b8b4df633d"),
	}
	MockProducer := NewProducer(txpool, account)
	assert.NotNil(MockProducer)
	assert.Equal(mockAddress, account.Address)
}

func TestProducer_assembleBlock(t *testing.T) {
	assert := assert.New(t)
	txpool := txpool.NewTxPool(txpool.DefaultTxPoolConfig)
	account := &account2.Account{
		Address: tools.HexToAddress("333c3310824b7c685133f2bedb2ca4b8b4df633d"),
	}
	MockProducer := NewProducer(txpool, account)
	types.GlobalEventCenter = &eventCenter{}
	conf := config.BlockChainConfig{
		PluginName:    blockchain.PLUGIN_MEMDB,
		StateDataPath: "./state",
		BlockDataPath: "./state",
	}
	err := blockchain.InitBlockChain(conf)
	assert.Nil(err)
	blockChain, err := blockchain.NewLatestStateBlockChain()
	assert.NotNil(blockChain)
	assert.Nil(err)
	block, err1 := MockProducer.assembleBlock(blockChain)
	assert.Nil(err1)
	assert.NotNil(block)
	assert.Equal(uint64(1), block.Header.Height)
}

func TestProducer_MakeBlock(t *testing.T) {
	assert := assert.New(t)
	txpool := txpool.NewTxPool(txpool.DefaultTxPoolConfig)
	account := &account2.Account{
		Address: tools.HexToAddress("333c3310824b7c685133f2bedb2ca4b8b4df633d"),
	}
	MockProducer := NewProducer(txpool, account)
	block, err := MockProducer.MakeBlock()
	assert.Nil(err)
	assert.NotNil(block)
	assert.Equal(uint64(1), block.Header.Height)

	monkey.Patch(blockchain.NewLatestStateBlockChain, func() (*blockchain.BlockChain, error) {
		return nil, fmt.Errorf("mock error")
	})
	block, err = MockProducer.MakeBlock()
	assert.NotNil(err)
	assert.Nil(block)
}

func Test_verifyBlock(t *testing.T) {
	assert := assert.New(t)
	MockProducer := NewProducer(nil, nil)
	var d *worker.Worker
	block := &types.Block{
		Header: &types.Header{
			Height: 0,
		},
	}
	monkey.PatchInstanceMethod(reflect.TypeOf(d), "VerifyBlock", func(*worker.Worker) error {
		return fmt.Errorf("mock verify failed")
	})
	err := MockProducer.verifyBlock(block, nil)
	assert.NotNil(err)

	monkey.PatchInstanceMethod(reflect.TypeOf(d), "VerifyBlock", func(*worker.Worker) error {
		return nil
	})
	err = MockProducer.verifyBlock(block, nil)
	assert.Nil(err)
}

func Test_signBlock(t *testing.T) {
	assert := assert.New(t)
	MockProducer := NewProducer(nil, nil)
	block := &types.Block{
		SigData: [][]byte{
			{0x1, 0x2, 0x3},
		},
	}
	// test sign error
	monkey.Patch(signature.Sign, func(signer signature.Signer, data []byte) ([]byte, error) {
		except := []byte{0x1, 0x2, 0x4}
		return except, fmt.Errorf("mock sign error")
	})
	err := MockProducer.signBlock(block)
	assert.Equal(err, fmt.Errorf("signature error:mock sign error"))

	// test new sign
	monkey.Patch(signature.Sign, func(signer signature.Signer, data []byte) ([]byte, error) {
		except := []byte{0x1, 0x2, 0x4}
		return except, nil
	})
	err = MockProducer.signBlock(block)
	assert.Nil(err)
	assert.Equal(2, len(block.SigData))
	// test duplicate sign
	monkey.Patch(signature.Sign, func(signer signature.Signer, data []byte) ([]byte, error) {
		except := []byte{0x1, 0x2, 0x3}
		return except, nil
	})
	err = MockProducer.signBlock(block)
	assert.Nil(err)
	assert.Equal(2, len(block.SigData))
}
