<div align="center" style="margin-bottom: 15px;">
  <img src="./assets/tcq-protocol-logo2.png" width="700" height="auto">
</div>

# TCQ Network Protocols (Thread Controlled Quick)

![GitHub repo size](https://img.shields.io/github/repo-size/NguyenHien-8/TCQ-Network-Protocol)
[![GitHub](https://img.shields.io/github/license/NguyenHien-8/TCQ-Network-Protocol)](https://github.com/NguyenHien-8/TCQ-Network-Protocol/blob/master/LICENSE)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/NguyenHien-8/TCQ-Network-Protocol)](https://pkg.go.dev/github.com/NguyenHien-8/TCQ-Network-Protocol)
[![Code Coverage](https://codecov.io/gh/NguyenHien-8/Test_Repo/graph/badge.svg?token=EOAAYMCF77)](https://codecov.io/gh/NguyenHien-8/Test_Repo)

TCQ is an implementation of the QUIC protocol but adds several features that allow for bandwidth optimization, processing flow optimization, CPU optimization, and data compression. ([RFC 9000](https://datatracker.ietf.org/doc/html/rfc9000), [RFC 9001](https://datatracker.ietf.org/doc/html/rfc9001), [RFC 9002](https://datatracker.ietf.org/doc/html/rfc9002)) in Go. It has support for HTTP/3 ([RFC 9114](https://datatracker.ietf.org/doc/html/rfc9114)), including QPACK ([RFC 9204](https://datatracker.ietf.org/doc/html/rfc9204)) and HTTP Datagrams ([RFC 9297](https://datatracker.ietf.org/doc/html/rfc9297)).

In addition to these base RFCs, it also implements the following RFCs:

* Unreliable Datagram Extension ([RFC 9221](https://datatracker.ietf.org/doc/html/rfc9221))
* Datagram Packetization Layer Path MTU Discovery (DPLPMTUD, [RFC 8899](https://datatracker.ietf.org/doc/html/rfc8899))
* QUIC Version 2 ([RFC 9369](https://datatracker.ietf.org/doc/html/rfc9369))
* QUIC Event Logging using qlog ([draft-ietf-quic-qlog-main-schema](https://datatracker.ietf.org/doc/draft-ietf-quic-qlog-main-schema/) and [draft-ietf-quic-qlog-quic-events](https://datatracker.ietf.org/doc/draft-ietf-quic-qlog-quic-events/))
* QUIC Stream Resets with Partial Delivery ([draft-ietf-quic-reliable-stream-reset](https://datatracker.ietf.org/doc/html/draft-ietf-quic-reliable-stream-reset-07))

Support for WebTransport over HTTP/3 ([draft-ietf-webtrans-http3](https://datatracker.ietf.org/doc/draft-ietf-webtrans-http3/)) is implemented in [webtransport-go](https://github.com/quic-go/webtransport-go).

## Release Policy

The TCQ protocol always aims to update to the latest versions.

### 1. Tạo một thư mục mới tại (path) và gắn nó với (branch_name)
```bash
git worktree add <path> <branch_name>
```

## Contributing

We are always happy to welcome new contributors. If you have any questions, please feel free to reach out by opening an issue or leaving a comment [issues](https://github.com/NguyenHien-8/TCQ-Network-Protocol/issues).

## License

The code is licensed under the Apache License 2.0. The logo and brand assets are excluded from the Apache License 2.0. See [assets/LICENSE.md](https://github.com/NguyenHien-8/TCQ-Network-Protocol/blob/Document/LICENSE) for the full usage policy and details.
