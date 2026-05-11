//  Project: TCQ Network Protocol (Thread Controlled QUIC)
//  Author: Trần Nguyên Hiền (c)
//  Major: Electronic And Communication Engineering
//  Email: trannguyenhien29085@gmail.com
//  Date: 2/3/2026
//  MIT Licence
//
// ----------------------------------------------------------------
//go:build linux

package quic

import (
	"fmt"
)

func init() {
	major, minor := kernelVersion()
	fmt.Printf("Kernel Version: %d.%d\n\n", major, minor)
}
