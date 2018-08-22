package common

import (
	"crypto/sha256"
	"github.com/DSiSc/txpool/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ComputeMerkleRoot(t *testing.T) {
	var data []common.Hash
	a1 := common.Hash(sha256.Sum256([]byte("a")))
	a2 := common.Hash(sha256.Sum256([]byte("b")))
	a3 := common.Hash(sha256.Sum256([]byte("c")))
	a4 := common.Hash(sha256.Sum256([]byte("d")))
	a5 := common.Hash(sha256.Sum256([]byte("e")))
	data = append(data, a1)
	data = append(data, a2)
	data = append(data, a3)
	data = append(data, a4)
	data = append(data, a5)
	hash := ComputeMerkleRoot(data)
	assert.NotEqual(t, hash, common.Hash{})
}
