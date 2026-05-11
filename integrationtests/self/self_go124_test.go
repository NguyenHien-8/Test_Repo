//  Project: TCQ Network Protocol (Thread Controlled QUIC)
//  Author: Trần Nguyên Hiền (c)
//  Major: Electronic And Communication Engineering
//  Email: trannguyenhien29085@gmail.com
//  Date: 2/3/2026
//  MIT Licence
//
// ----------------------------------------------------------------
//go:build !go1.25

package self_test

import "crypto/tls"

func getCurveID(connState tls.ConnectionState) tls.CurveID {
	return 0
}
