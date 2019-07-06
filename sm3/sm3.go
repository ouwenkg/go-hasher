package sm3

import (
	"fmt"
)

type Sm3 struct {
}

func (s Sm3) digest(data []byte) []byte {
	hasher := New()
	sum := hasher.Sum(data)
	fmt.Println("digest sm3...", sum)
	return sum
}
