package main

import (
	"hasher/sm3"
	"hasher/sha3"
	"hasher/blake2b"
)

type Hasher interface {
	digest(data []byte) []byte
}


func main() {
	var hasher Hasher

	hasher = new(sha3.Sha3)
	hasher.digest([]byte(""))

	hasher = new(blake2b.Blake2b)
    hasher.digest([]byte(""))

    hasher = new(sm3.Sm3)
    hasher.digest([]byte(""))
}