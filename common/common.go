package common

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"github.com/DSiSc/craft/log"
	"github.com/DSiSc/craft/types"
)

// TODO: Hash algorithm will support configurable later
// Sum returns the first 32 bytes of SHA256 of the bz.
func Sum(bz []byte) []byte {
	hash := sha256.Sum256(bz)
	return hash[:types.HashLength]
}

func TxHash(tx *types.Transaction) types.Hash {
	if hash := tx.Hash.Load(); hash != nil {
		return hash.(types.Hash)
	}
	hashData := types.TxData{
		AccountNonce: tx.Data.AccountNonce,
		Price:        tx.Data.Price,
		GasLimit:     tx.Data.GasLimit,
		Recipient:    tx.Data.Recipient,
		Amount:       tx.Data.Amount,
		Payload:      tx.Data.Payload,
		V:            tx.Data.V,
		R:            tx.Data.R,
		S:            tx.Data.S,
	}
	jsonByte, _ := json.Marshal(hashData)
	sumByte := Sum(jsonByte)
	var temp types.Hash
	copy(temp[:], sumByte)
	tx.Hash.Store(temp)
	return temp
}

func HeaderDigest(header *types.Header) (hash types.Hash) {
	var defaultHash types.Hash
	if !bytes.Equal(header.MixDigest[:], defaultHash[:]) {
		log.Info("header hash %v has exits.", header.MixDigest)
		copy(hash[:], header.MixDigest[:])
		return
	}
	newHeader := digestHeader(header)
	jsonByte, _ := json.Marshal(newHeader)
	sumByte := Sum(jsonByte)
	copy(hash[:], sumByte)
	return
}

func digestHeader(header *types.Header) *types.Header {
	return &types.Header{
		ChainID:       header.ChainID,
		PrevBlockHash: header.PrevBlockHash,
		StateRoot:     header.StateRoot,
		TxRoot:        header.TxRoot,
		ReceiptsRoot:  header.ReceiptsRoot,
		Height:        header.Height,
		Timestamp:     header.Timestamp,
		CoinBase:      header.CoinBase,
	}
}

func CopyBytes(b []byte) (copiedBytes []byte) {
	if len(b) == 0 {
		return
	}
	copiedBytes = make([]byte, len(b))
	copy(copiedBytes, b)

	return
}
