package models

import (
	"crypto/ecdh"
)

// DefaultIp const origin string = "http://localhost/"
// const url string = "ws://localhost:1444/ws"
const DefaultIp string = "localhost"

type MessageType int
type KeyType int

const (
	TextMessage MessageType = iota + 1
	Ping
	KeyExchange
)
const (
	IdentityKey KeyType = iota + 1
	EphemeralKey
	PreKeys
)

// Port var ip string
const Port string = "1444"

type Msg struct {
	Username string `json:"username"`
	Text     string `json:"Text"`
}
type DataPayload struct {
	Type     MessageType `json:"type"`
	Username string      `json:"username"`
	Data     interface{} `json:"data"`
}

// Key I decided to have a key struct which will contain the type of key being sent, and then
// I used an interface for the actual key because I wanted flexibility so it could be a slice of byte slices.
type Key struct {
	KeyType KeyType     `json:"keyType"`
	Key     interface{} `json:"key"`
}
type X3dhInit struct {
	IdentityKey []byte   `json:"identityKey"`
	PreKeys     [][]byte `json:"preKeys"`
}
type PrivateKeys struct {
	IdentityKeyPrivate []byte
	PreKeysPrivate     [][]byte
}
type KeyPair struct {
	PublicKey  *ecdh.PublicKey
	PrivateKey *ecdh.PrivateKey
}

func (m MessageType) String() string {
	switch m {
	case TextMessage:
		return "Text Message"
	case Ping:
		return "Ping"
	case KeyExchange:
		return "KeyExchange"

	default:
		return "Unknown"
	}
}
func (k KeyType) String() string {
	switch k {
	case IdentityKey:
		return "IdentityKey"
	case EphemeralKey:
		return "EphemeralKey"
	case PreKeys:
		return "PreKeys"
	default:
		return "Unknown"
	}
}
