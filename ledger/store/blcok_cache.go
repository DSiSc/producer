package store

import (
	"fmt"
	"github.com/DSiSc/producer/common"
	"github.com/hashicorp/golang-lru"
)

const (
	BLOCK_CAHE_SIZE        = 10    //Block cache size
	TRANSACTION_CACHE_SIZE = 10000 //Transaction cache size
)

// BlockCache store block cache and transaction hash
type BlockCache struct {
	blockCache       *lru.ARCCache
	transactionCache *lru.ARCCache
}

// NewBlockCache return BlockCache instance
func NewBlockCache() (*BlockCache, error) {
	blockCache, err := lru.NewARC(BLOCK_CAHE_SIZE)
	if err != nil {
		return nil, fmt.Errorf("NewARC block error %s", err)
	}
	transactionCache, err := lru.NewARC(TRANSACTION_CACHE_SIZE)
	if err != nil {
		return nil, fmt.Errorf("NewARC header error %s", err)
	}
	return &BlockCache{
		blockCache:       blockCache,
		transactionCache: transactionCache,
	}, nil
}

//AddBlock to cache
func (this *BlockCache) AddBlock(block *common.Block) {
	blockHash := block.Hash()
	this.blockCache.Add(string(blockHash.ToArray()), block)
}
