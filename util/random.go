package util

import "github.com/ethereum/go-ethereum/crypto"

// RandomMessage ...
func RandomMessage(seed string) string {
	// dummy random message generator
	// TODO refactor later
	return crypto.Keccak256Hash([]byte(seed)).Hex()
}
