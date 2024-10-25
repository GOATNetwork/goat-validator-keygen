package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/tendermint/tendermint/crypto/secp256k1"
)

func generateTendermintKeys(prvkeyHex string) {
	// Decode hexadecimal private key
	prvkeyBytes, _ := hex.DecodeString(prvkeyHex[2:]) // Remove "0x" prefix

	// Create secp256k1 private key
	privKey := secp256k1.PrivKey(prvkeyBytes)

	// Get public key
	pubKey := privKey.PubKey()

	// Get address
	address := pubKey.Address()

	// Convert public key to compressed format
	compressedPubKey := pubKey.Bytes()

	// Base64 encode the compressed public key
	pubKeyBase64 := base64.StdEncoding.EncodeToString(compressedPubKey)

	// Output results
	fmt.Printf("Address: %X\n", address)
	fmt.Printf("Public Key (Tendermint format): %s\n", pubKeyBase64)

	// Generate output in priv_validator_key.json format
	jsonOutput := fmt.Sprintf(`{
  "address": "%X",
  "pub_key": {
    "type": "tendermint/PubKeySecp256k1",
    "value": "%s"
  },
  "priv_key": {
    "type": "tendermint/PrivKeySecp256k1",
    "value": "%s"
  }
}`, address, pubKeyBase64, base64.StdEncoding.EncodeToString(privKey))

	fmt.Println("\npriv_validator_key.json format:")
	fmt.Println(jsonOutput)
}

func main() {
	prvkeys := strings.Split(os.Args[1], ",")

	for i, prvkey := range prvkeys {
		fmt.Printf("Generating keys for validator %d:\n", i+1)
		generateTendermintKeys(prvkey)
		fmt.Println("--------------------")
	}
}
