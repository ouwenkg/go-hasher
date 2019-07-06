package main

import (
	"hasher/sm3"
	"hasher/sha3"
	"hasher/blake2b"
)

type Hasher interface {
	Digest(data []byte) []byte
}


func main() {
	var hasher Hasher

	hasher = new(sha3.Sha3)
	hasher.Digest([]byte(""))

	hasher = new(blake2b.Blake2b)
    hasher.Digest([]byte(""))

    hasher = new(sm3.Sm3)
    hasher.Digest([]byte(""))
}