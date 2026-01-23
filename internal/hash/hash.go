package hash

import "crypto/sha256"

func Get256(salt string, pass string) [32]byte {
	comb := salt + pass
	hash := sha256.Sum256([]byte(comb))
	return hash
}
