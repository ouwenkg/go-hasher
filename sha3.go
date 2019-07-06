package hasher

import (
	"fmt"
	"github.com/ebfe/keccak"
)

type Hasher interface {
	digest(data []byte) []byte
}

type Sha3 struct {
}

func (sha3 Sha3) digest(data []byte) []byte {
	hasher := keccak.New256()
	sum := hasher.Sum(data)
	fmt.Println("digest sha3...", sum)
	return sum
}

// func main() {
// 	var hasher Hasher

// 	hasher = new(Sha3)
// 	hasher.digest([]byte(""))

// 	hasher = new(Blake2b)
//     hasher.digest([]byte(""))

//     hasher = new(Sm3)
//     hasher.digest([]byte(""))
// }
