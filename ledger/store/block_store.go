package store

import (
	"fmt"
	"github.com/DSiSc/producer/config"
	"github.com/DSiSc/producer/ledger/store/leveldbstore"
	"github.com/DSiSc/txpool/common/log"
)

const (
	// DB plugin
	PLUGIN_LEVELDB = "leveldb"
	// Defaule setting of blcok
	BLOCK_CONFIG    = "block"
	DB_STORE_PLUGIN = "block.store"
)

type DBStore interface {
	Put(key []byte, value []byte) error
	Get(key []byte) ([]byte, error)
	Delete(key []byte) error
}

// Block store save the data of block & transaction
type BlockStore struct {
	cache       *BlockCache // The cache of block, if have.
	enableCache bool        //Is enable lru cache
	dbDir       string      // The path of store file
	store       DBStore     // Block store handler
}

// NewBlockStore return the block store instance
func NewBlockStore(dbDir string, enableCache bool) (*BlockStore, error) {
	var cache *BlockCache
	var store DBStore
	var err error
	if enableCache {
		cache, err = NewBlockCache()
		if err != nil {
			return nil, fmt.Errorf("NewBlockCache error %s", err)
		}
	}

	conf := config.New(config.ConfigAbsPath())
	plugin := conf.GetConfigItem(DB_STORE_PLUGIN).(string)
	switch plugin {
	case PLUGIN_LEVELDB:
		store, err = leveldbstore.NewLevelDBStore(dbDir)
		if err != nil {
			return nil, err
		}
	default:
		log.Error("Not support plugin.")
		return nil, fmt.Errorf("Not support plugin type %s", plugin)
	}

	blockStore := &BlockStore{
		dbDir:       dbDir,
		enableCache: enableCache,
		store:       store,
		cache:       cache,
	}
	return blockStore, nil
}
