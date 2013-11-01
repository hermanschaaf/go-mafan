package mafan

import (
	"fmt"
)

func Split(s string) (tokenized []string) {
	fmt.Println("Start split")
	tokenized = TrieRoot.Split(s)
	return tokenized
}
