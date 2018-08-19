package common

import (
	"encoding/binary"
	"io"
)

func WriteUint32(writer io.Writer, val uint32) error {
	var p [4]byte
	binary.LittleEndian.PutUint32(p[:], val)
	_, err := writer.Write(p[:])
	return err
}

func WriteUint64(writer io.Writer, val uint64) error {
	var p [8]byte
	binary.LittleEndian.PutUint64(p[:], val)
	_, err := writer.Write(p[:])
	return err
}
