package crypto

import (
	"Vanta/client/models"
	"crypto/ecdh"
	"crypto/rand"
	"fmt"
	"log"
)

func keyGen() models.KeyPair {
	key := ecdh.X25519()
	privateKey, err := key.GenerateKey(rand.Reader)
	if err != nil {
		log.Fatal(err)
		//change this later.
	}
	publicKey := privateKey.PublicKey()
	KP := models.KeyPair{PublicKey: publicKey, PrivateKey: privateKey}

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

func CreateIdentityKeys() models.KeyPair {
	return keyGen()
}
func CreateEphemeralKeys() models.KeyPair {
	return keyGen()
}
