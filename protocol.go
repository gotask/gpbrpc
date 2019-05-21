// protocol.go
package gpbrpc

import (
	"encoding/binary"
	"errors"
)

var (
	ErrInvalidMsgLen = errors.New("message length is invalid.")
)

func EncodeLength(b []byte, i uint32) {
	binary.BigEndian.PutUint32(b, i)
}
func DecodeLength(b []byte) uint32 {
	return binary.BigEndian.Uint32(b)
}
func PackSendProtocol(data []byte) []byte {
	msglen := len(data) + 4
	buf := make([]byte, msglen, msglen)
	EncodeLength(buf, uint32(msglen))
	copy(buf[4:], data)
	return buf
}
