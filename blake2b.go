package hasher

import (
	"fmt"
	"golang.org/x/crypto/blake2b"
)

type Blake2b struct {
}

func (b Blake2b) digest(data []byte) []byte {
	hasher, _ := blake2b.New256([]byte("AsceticBear"))
	sum := hasher.Sum(data)
	fmt.Println("digest blake2b...", sum)
	return sum
}
