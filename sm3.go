package hasher

import (
	"fmt"
	"hasher/sm3"
)

type Sm3 struct {
}

func (s Sm3) digest(data []byte) []byte {
	hasher := sm3.New()
	sum := hasher.Sum(data)
	fmt.Println("digest sm3...", sum)
	return sum
}
