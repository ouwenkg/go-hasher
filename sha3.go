package main

import (
	"fmt"
    "github.com/ebfe/keccak"
    "golang.org/x/crypto/blake2b"
    "hasher/sm3"
)

type Hasher interface {
	digest(data []byte) []byte
}

type Sha3 struct {
}

type Blake2b struct {
}

type Sm3 struct {
}

func (sha3 Sha3) digest(data []byte) []byte {
    hasher := keccak.New256()
    sum := hasher.Sum(data)
    fmt.Println("digest sha3...", sum)
	return sum
}

func (b Blake2b) digest(data []byte) []byte {
    hasher, _ := blake2b.New256([]byte("AsceticBear"))
    sum := hasher.Sum(data)
	fmt.Println("digest blake2b...", sum)
	return sum
}

func (s Sm3) digest(data []byte) []byte {
    hasher := sm3.New()
    sum := hasher.Sum(data)
    fmt.Println("digest sm3...", sum)
    return sum
}

func main() {
	var hasher Hasher

	hasher = new(Sha3)
	hasher.digest([]byte(""))

	hasher = new(Blake2b)
    hasher.digest([]byte(""))
    
    hasher = new(Sm3)
    hasher.digest([]byte(""))
}
