package crypto

import (
	"crypto/ecdh"
	"crypto/rand"
	"fmt"
)

func KeyGen() (*ecdh.PublicKey, *ecdh.PrivateKey) {
	key := ecdh.P256()
	privKey, err := key.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Println("failed KeyGen ", err)
	}
	pubKey := privKey.PublicKey()
	//fmt.Println("Here is your public Key ", pubKey)

	// fmt.Println("Here is your private Key ", privKey)

	return pubKey, privKey
}

func SharedSecret(pubKey *ecdh.PublicKey, privKey *ecdh.PrivateKey) ([]byte, error) {
	secret, err := privKey.ECDH(pubKey)
	if err != nil {
		fmt.Println("Failed to compute")
		return nil, err
	}
	return secret, nil
}
