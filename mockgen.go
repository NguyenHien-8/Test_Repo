//  Project: TCQ Network Protocol (Thread Controlled QUIC)
//  Author: Trần Nguyên Hiền (c)
//  Major: Electronic And Communication Engineering
//  Email: trannguyenhien29085@gmail.com
//  Date: 2/3/2026
//	Apache License 2.0
//
// ----------------------------------------------------------------
//go:build gomock || generate

package quic

//go:generate sh -c "go tool mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/NguyenHien-8/TCQ-Network-Protocol -destination mock_send_conn_test.go github.com/NguyenHien-8/TCQ-Network-Protocol SendConn"
type SendConn = sendConn

//go:generate sh -c "go tool mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/NguyenHien-8/TCQ-Network-Protocol -destination mock_raw_conn_test.go github.com/NguyenHien-8/TCQ-Network-Protocol RawConn"
type RawConn = rawConn

//go:generate sh -c "go tool mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/NguyenHien-8/TCQ-Network-Protocol -destination mock_sender_test.go github.com/NguyenHien-8/TCQ-Network-Protocol Sender"
type Sender = sender

//go:generate sh -c "go tool mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/NguyenHien-8/TCQ-Network-Protocol -destination mock_stream_sender_test.go github.com/NguyenHien-8/TCQ-Network-Protocol StreamSender"
type StreamSender = streamSender

//go:generate sh -c "go tool mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/NguyenHien-8/TCQ-Network-Protocol -destination mock_stream_control_frame_getter_test.go github.com/NguyenHien-8/TCQ-Network-Protocol StreamControlFrameGetter"
type StreamControlFrameGetter = streamControlFrameGetter

//go:generate sh -c "go tool mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/NguyenHien-8/TCQ-Network-Protocol -destination mock_stream_frame_getter_test.go github.com/NguyenHien-8/TCQ-Network-Protocol StreamFrameGetter"
type StreamFrameGetter = streamFrameGetter

//go:generate sh -c "go tool mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/NguyenHien-8/TCQ-Network-Protocol -destination mock_frame_source_test.go github.com/NguyenHien-8/TCQ-Network-Protocol FrameSource"
type FrameSource = frameSource

//go:generate sh -c "go tool mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/NguyenHien-8/TCQ-Network-Protocol -destination mock_ack_frame_source_test.go github.com/NguyenHien-8/TCQ-Network-Protocol AckFrameSource"
type AckFrameSource = ackFrameSource

//go:generate sh -c "go tool mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/NguyenHien-8/TCQ-Network-Protocol -destination mock_sealing_manager_test.go github.com/NguyenHien-8/TCQ-Network-Protocol SealingManager"
type SealingManager = sealingManager

//go:generate sh -c "go tool mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/NguyenHien-8/TCQ-Network-Protocol -destination mock_unpacker_test.go github.com/NguyenHien-8/TCQ-Network-Protocol Unpacker"
type Unpacker = unpacker

//go:generate sh -c "go tool mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/NguyenHien-8/TCQ-Network-Protocol -destination mock_packer_test.go github.com/NguyenHien-8/TCQ-Network-Protocol Packer"
type Packer = packer

//go:generate sh -c "go tool mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/NguyenHien-8/TCQ-Network-Protocol -destination mock_mtu_discoverer_test.go github.com/NguyenHien-8/TCQ-Network-Protocol MTUDiscoverer"
type MTUDiscoverer = mtuDiscoverer

//go:generate sh -c "go tool mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/NguyenHien-8/TCQ-Network-Protocol -destination mock_conn_runner_test.go github.com/NguyenHien-8/TCQ-Network-Protocol ConnRunner"
type ConnRunner = connRunner

//go:generate sh -c "go tool mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/NguyenHien-8/TCQ-Network-Protocol -destination mock_packet_handler_test.go github.com/NguyenHien-8/TCQ-Network-Protocol PacketHandler"
type PacketHandler = packetHandler

//go:generate sh -c "go tool mockgen -typed -package quic -self_package github.com/NguyenHien-8/TCQ-Network-Protocol -self_package github.com/NguyenHien-8/TCQ-Network-Protocol -destination mock_packetconn_test.go net PacketConn"
