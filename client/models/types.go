package models

import "crypto/ecdh"

// DefaultIp const origin string = "http://localhost/"
// const url string = "ws://localhost:1444/ws"
const DefaultIp string = "localhost"

// Port var ip string
const Port string = "1444"

type Msg struct {
	Username string `json:"username"`
	Text     string `json:"Text"`
}
type X3dhInit struct {
	IdentityKey []byte   `json:"identityKey"`
	Prekeys     [][]byte `json:"preKeys"`
}
type PrivateKeys struct {
	IdentityKeyPrivate []byte
	PreKeysPrivate     [][]byte
}
type KeyPair struct {
	PublicKey  *ecdh.PublicKey
	PrivateKey *ecdh.PrivateKey
}
