package crypto

import (
	"crypto/ecdh"
	"crypto/rand"
	"fmt"
	"log"
)

type KeyPair struct {
	publicKey  *ecdh.PublicKey
	privateKey *ecdh.PrivateKey
}

func keyGen() KeyPair {
	key := ecdh.X25519()
	privateKey, err := key.GenerateKey(rand.Reader)
	if err != nil {
		log.Fatal(err)
		//change this later.
	}
	publicKey := privateKey.PublicKey()
	KP := KeyPair{publicKey, privateKey}

	return KP
}

func ComputeSecret(pubKey *ecdh.PublicKey, privKey *ecdh.PrivateKey) ([]byte, error) {
	secret, err := privKey.ECDH(pubKey)
	if err != nil {
		fmt.Println("Failed to compute secret!")
		return nil, err
	}
	return secret, nil
}

func CreateIdentityKeys() KeyPair {
	return keyGen()
}
func CreateEphemeralKeys() KeyPair {
	return keyGen()
}

func ComputeKey(DH1, DH2, DH3) []byte {

}
