package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

// Key is a generated Lamport public or private key, consisting of 256 pairs of 256bit randomly generated bits
type Key [256][2][]byte

// PrivateKey is a 256 pairs of 256bit hex strings representing a Lamport private key
type PrivateKey [256][2]string

// PublicKey is a 256 pairs of 256bit hex strings representing a Lamport public key
type PublicKey [256][2]string

// GeneratePrivateKey will return a private key generated using the Lamport Algorithm
func GeneratePrivateKey() (PrivateKey, error) {
	var privateKey PrivateKey
	for i, section := range privateKey {
		for j := range section {
			b := make([]byte, 32)
			_, err := rand.Read(b)
			// Note that err == nil only if we read len(b) bytes.
			if err != nil {
				return privateKey, err
			}
			privateKey[i][j] = hex.EncodeToString(b)
		}
	}
	return privateKey, nil
}

// GeneratePublicKey will return the corresponding public key for a given Lamport private key
func GeneratePublicKey(privateKey PrivateKey) (PublicKey, error) {
	var publicKey PublicKey
	for i, section := range privateKey {
		for j := range section {

			keyBytes, err := hex.DecodeString(privateKey[i][j])
			if err != nil {
				return publicKey, err
			}
			b := sha256.New()
			publicKey[i][j] = hex.EncodeToString(b.Sum(keyBytes))
		}
	}
	return publicKey, nil
}

// GenerateKeyPairs will generate N keypairs
func GenerateKeyPairs(n int) ([]PrivateKey, []PublicKey, error) {
	publicKeys := make([]PublicKey, n)
	privateKeys := make([]PrivateKey, n)

	var err error
	for i := range publicKeys {
		privateKeys[i], err = GeneratePrivateKey()
		if err != nil {
			return privateKeys, publicKeys, err
		}
		publicKeys[i], err = GeneratePublicKey(privateKeys[i])
		if err != nil {
			return privateKeys, publicKeys, err
		}
	}
	return privateKeys, publicKeys, nil
}
