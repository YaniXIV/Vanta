# Vanta - End-to-End Encrypted Terminal Chat

A secure, terminal-based chat application implementing the Signal Protocol for end-to-end encryption (E2EE).

## ⚠️ Important Disclaimer

This project is purely experimental and educational. It is:
- Not intended for production use
- Not audited for security
- Far from complete
- Created for learning and experimentation purposes only

**DO NOT** use this for any real-world secure communications. This is a personal project to learn about and experiment with the Signal Protocol implementation.

## Features

- 🔒 Signal Protocol implementation for end-to-end encryption
- 💻 Terminal-based user interface
- 🔑 Double Ratchet algorithm for perfect forward secrecy
- 🌐 WebSocket-based real-time communication
- 🔄 X3DH (Extended Triple Diffie-Hellman) for initial key exchange
- 🔐 Pre-keys for asynchronous messaging

## Technical Details

The application implements the Signal Protocol, which includes:
- X3DH for initial key exchange
- Double Ratchet algorithm for message encryption
- Pre-keys for asynchronous communication
- ECDH (Elliptic Curve Diffie-Hellman) for key generation
- WebSocket for real-time communication
- Go's standard cryptographic libraries

## Project Structure

```
.
├── client/         # Client implementation
│   ├── core/       # Core client functionality
│   └── crypto/     # Cryptographic operations
│       ├── x3dh.go     # X3DH key exchange
│       ├── curve.go    # ECDH operations
│       └── key_manager.go  # Key management
├── server/         # Server implementation
│   └── core/       # Core server functionality
└── Testing/        # Test implementations
```

## Prerequisites

- Go 1.16 or higher
- A terminal emulator

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/Vanta.git
cd Vanta
```

2. Build the server:
```bash
go build -o server ./server
```

3. Build the client:
```bash
go build -o client ./client
```

## Usage

1. Start the server:
```bash
./server
```

2. Start the client:
```bash
./client
```

## Security Features

- **Signal Protocol Components**:
  - **X3DH**: Initial key exchange protocol
  - **Double Ratchet**: Message encryption with perfect forward secrecy
  - **Pre-Keys**: One-time use keys for asynchronous communication
  - **Identity Keys**: Long-term keys for user identification
  - **Ephemeral Keys**: Short-lived keys for forward secrecy

## Development Status

🚧 This project is currently under active development. Features and APIs may change.

### Current Status
- Basic WebSocket communication implemented
- X3DH key exchange implementation in progress
- Double Ratchet algorithm implementation pending
- Terminal UI implementation ongoing
- Many security-critical components are incomplete or missing

### Planned Features
- [ ] Complete Double Ratchet implementation
- [ ] Message persistence
- [ ] Group chat support

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Acknowledgments

- [Signal Protocol Documentation](https://signal.org/docs/)
- [X3DH Protocol Specification](https://signal.org/docs/specifications/x3dh/)
- [Double Ratchet Specification](https://signal.org/docs/specifications/doubleratchet/)
- [Go WebSocket](https://pkg.go.dev/golang.org/x/net/websocket)
- [Go Crypto](https://pkg.go.dev/crypto) 