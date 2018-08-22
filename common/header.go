package common

import (
	"bytes"
	"crypto/sha256"
	"errors"
	ledger_c "github.com/DSiSc/ledger/common"
	"github.com/DSiSc/txpool/common"
	"io"
)

type Header struct {
	Version          uint32
	PrevBlockHash    common.Hash
	TransactionsRoot common.Hash
	BlockRoot        common.Hash
	Timestamp        uint32
	Height           uint32
	SigData          [][]byte
	hash             *common.Hash
}

//Serialize the blockheader data without program
func (bd *Header) SerializeUnsigned(w io.Writer) error {
	err := ledger_c.WriteUint32(w, bd.Version)
	if err != nil {
		return err
	}
	err = bd.PrevBlockHash.Serialize(w)
	if err != nil {
		return err
	}
	err = bd.TransactionsRoot.Serialize(w)
	if err != nil {
		return err
	}
	err = bd.BlockRoot.Serialize(w)
	if err != nil {
		return err
	}
	err = ledger_c.WriteUint32(w, bd.Timestamp)
	if err != nil {
		return err
	}
	err = ledger_c.WriteUint32(w, bd.Height)
	if err != nil {
		return err
	}

	return nil
}

func (h *Header) Hash() common.Hash {
	if h.hash != nil {
		return *h.hash
	}
	buf := new(bytes.Buffer)
	h.SerializeUnsigned(buf)
	temp := sha256.Sum256(buf.Bytes())
	hash := common.Hash(sha256.Sum256(temp[:]))

	h.hash = &hash
	return hash
}

//Serialize the blockheader
func (h *Header) Serialize(w io.Writer) error {
	h.SerializeUnsigned(w)
	err := ledger_c.WriteVarUint(w, uint64(len(h.SigData)))
	if err != nil {
		return errors.New("serialize sig pubkey length failed")
	}

	for _, sig := range h.SigData {
		err = ledger_c.WriteVarBytes(w, sig)
		if err != nil {
			return err
		}
	}

	return nil
}
