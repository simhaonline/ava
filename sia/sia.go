package sia

import (
	"encoding/hex"

	"golang.org/x/crypto/blake2b"
)

// Validator - Sia address validator
type Validator struct{}

// New - Create a Sia address validator
func New() *Validator {
	return &Validator{}
}

// IsValidAddress - Check a Sia address is valid or not
func (s *Validator) IsValidAddress(address string, isTestnet bool) bool {
	unlockhashWithChecksum, err := hex.DecodeString(address)
	if err != nil || len(unlockhashWithChecksum) != 38 {
		return false
	}
	unlockhash := unlockhashWithChecksum[0:32]
	checksum256 := blake2b.Sum256(unlockhash)

	var validChecksum [6]byte
	copy(validChecksum[:], checksum256[:6])

	var checksum [6]byte
	copy(checksum[:], unlockhashWithChecksum[32:])
	return checksum == validChecksum
}
