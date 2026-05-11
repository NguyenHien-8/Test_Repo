//	Project: TCQ Network Protocol (Thread Controlled QUIC)
//	Author: Trần Nguyên Hiền (c)
//	Major: Electronic And Communication Engineering
//	Email: trannguyenhien29085@gmail.com
//	Date: 2/3/2026
//	MIT Licence
//
// ----------------------------------------------------------------
package wire

import (
	"github.com/NguyenHien-8/TCQ-Network-Protocol/internal/protocol"
)

// A HandshakeDoneFrame is a HANDSHAKE_DONE frame
type HandshakeDoneFrame struct{}

func (f *HandshakeDoneFrame) Append(b []byte, _ protocol.Version) ([]byte, error) {
	return append(b, byte(FrameTypeHandshakeDone)), nil
}

// Length of a written frame
func (f *HandshakeDoneFrame) Length(_ protocol.Version) protocol.ByteCount {
	return 1
}
