package main


import(
  "crypto/ecdh"
  "fmt"
  "crypto/rand"
)

func KeyGen()(*ecdh.PublicKey, *ecdh.PrivateKey){
  key := ecdh.P256()
  privKey, err := key.GenerateKey(rand.Reader)
  if err != nil{
    fmt.Println("failed KeyGen ", err)
  }
  pubKey := privKey.PublicKey()
  //fmt.Println("Here is your public Key ", pubKey)

 // fmt.Println("Here is your private Key ", privKey)


  return pubKey, privKey
}



