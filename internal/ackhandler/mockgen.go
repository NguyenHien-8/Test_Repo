//  Project: TCQ Network Protocol (Thread Controlled QUIC)
//  Author: Trần Nguyên Hiền (c)
//  Major: Electronic And Communication Engineering
//  Email: trannguyenhien29085@gmail.com
//  Date: 2/3/2026
//  MIT Licence
//
// ----------------------------------------------------------------
//go:build gomock || generate

package ackhandler

//go:generate sh -c "go tool mockgen -typed -build_flags=\"-tags=gomock\"  -package ackhandler -destination mock_ecn_handler_test.go github.com/NguyenHien-8/TCQ-Network-Protocol/internal/ackhandler ECNHandler"
type ECNHandler = ecnHandler
