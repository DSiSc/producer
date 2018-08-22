package common

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeader_Serialize(t *testing.T) {
	header := Header{}
	header.Height = 321
	header.SigData = make([][]byte, 0)
	buf := bytes.NewBuffer(nil)
	err := header.Serialize(buf)
	assert.Nil(t, err)
}
