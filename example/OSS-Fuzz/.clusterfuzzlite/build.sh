#!/bin/bash -eu

go version
go env

compile_native_go_fuzzer_v2 github.com/NguyenHien-8/TCQ-Network-Protocol/wire FuzzFrames frame_fuzzer
compile_native_go_fuzzer_v2 github.com/NguyenHien-8/TCQ-Network-Protocol/wire FuzzTransportParameters transportparameter_fuzzer
compile_native_go_fuzzer_v2 github.com/NguyenHien-8/TCQ-Network-Protocol/http3 FuzzFrameParser http3_frame_fuzzer
compile_native_go_fuzzer_v2 github.com/NguyenHien-8/TCQ-Network-Protocol/wire FuzzHeaderParser header_fuzzer
compile_native_go_fuzzer_v2 github.com/NguyenHien-8/TCQ-Network-Protocol/handshake FuzzHandshake handshake_fuzzer
compile_native_go_fuzzer_v2 github.com/NguyenHien-8/TCQ-Network-Protocol/http3 FuzzHeaderParsing http3_header_parsing_fuzzer
