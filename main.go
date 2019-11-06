package main

import (
	"crypto/rand"
	"fmt"
	"math/big"

	bls "github.com/nmarley/bls-signatures/go-bindings"
)

var RFieldModulus, _ = new(big.Int).SetString("52435875175126190479447740508185965837690552500527637822603658699938581184513", 10)

func main() {
	r, _ := rand.Int(rand.Reader, RFieldModulus)

	// bls.PrivateKeyFromBytes is probably a better call as it wraps
	// RFieldModulus w/the 2nd argument set to true but this is quicker to
	// flesh out an example
	sk := bls.PrivateKeyFromBN(r)

	fmt.Printf("Got a random SecretKey: %x\n", sk.Serialize())

	pk := sk.PublicKey()
	fmt.Printf("Corresponding PublicKey: %x\n", pk.Serialize())

	msg := []byte("Hello, world")
	fmt.Printf("Message: %s\n", string(msg))

	sig := sk.Sign(msg)
	fmt.Printf("Signature: %x\n", sig.Serialize())

	if sig.Verify() {
		fmt.Println("Sig did verify")
	} else {
		fmt.Println("Sig did NOT verify")
	}

	keyBytes := []byte{
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	}
	// this should blow up
	pk2, err := bls.PrivateKeyFromBytes(keyBytes, false)
	if err != nil {
		fmt.Println("THIS IS EXPECTED err:", err.Error())
	} else {
		fmt.Printf("THIS SHOULD NOT HAPPEN. pk2: %x\n", pk2.Serialize())
	}
}
