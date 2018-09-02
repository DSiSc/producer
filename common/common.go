package common

import (
	"crypto/sha256"
	"encoding/json"
	"github.com/DSiSc/craft/types"
)

// TODO: Hash algorithm will support configurable later
// Sum returns the first 20 bytes of SHA256 of the bz.
func Sum(bz []byte) []byte {
	hash := sha256.Sum256(bz)
	return hash[:types.HashLength]
}

func TxHash(tx *types.Transaction) (hash types.Hash) {
	jsonByte, _ := json.Marshal(tx)
	sumByte := Sum(jsonByte)
	copy(hash[:], sumByte)
	return
}

func BlockHash(block *types.Block) (hash types.Hash) {
	jsonByte, _ := json.Marshal(block)
	sumByte := Sum(jsonByte)
	copy(hash[:], sumByte)
	return
}

func CopyBytes(b []byte) (copiedBytes []byte) {
	if b == nil {
		return nil
	}
	copiedBytes = make([]byte, len(b))
	copy(copiedBytes, b)

	return
}
