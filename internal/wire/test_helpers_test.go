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
	"bytes"
	"encoding/binary"
	"log"
	"testing"

	"github.com/NguyenHien-8/TCQ-Network-Protocol/internal/protocol"
	"github.com/NguyenHien-8/TCQ-Network-Protocol/internal/utils"
	"github.com/NguyenHien-8/TCQ-Network-Protocol/quicvarint"
)

func encodeVarInt(i uint64) []byte {
	return quicvarint.Append(nil, i)
}

func appendVersion(data []byte, v protocol.Version) []byte {
	offset := len(data)
	data = append(data, []byte{0, 0, 0, 0}...)
	binary.BigEndian.PutUint32(data[offset:], uint32(v))
	return data
}

func setupLogTest(t *testing.T, buf *bytes.Buffer) utils.Logger {
	logger := utils.DefaultLogger
	logger.SetLogLevel(utils.LogLevelDebug)
	originalOutput := log.Writer()
	log.SetOutput(buf)
	t.Cleanup(func() { log.SetOutput(originalOutput) })
	return logger
}
