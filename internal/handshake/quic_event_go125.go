//  Project: TCQ Network Protocol (Thread Controlled QUIC)
//  Author: Trần Nguyên Hiền (c)
//  Major: Electronic And Communication Engineering
//  Email: trannguyenhien29085@gmail.com
//  Date: 2/3/2026
//  MIT Licence
//
// ----------------------------------------------------------------
//go:build go1.25 && !go1.26

package handshake

import "crypto/tls"

const quicErrorEvent tls.QUICEventKind = -1

func extractQUICEventError(tls.QUICEvent) error {
	return nil
}
