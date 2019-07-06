package sha3

import (
	"fmt"
	"github.com/ebfe/keccak"
)

type Sha3 struct {
}

func (s Sha3) digest(data []byte) []byte {
	hasher := keccak.New256()
	sum := hasher.Sum(data)
	fmt.Println("digest sha3...", sum)
	return sum
}

