package mafan

import (
	"fmt"
)

func Split(origin string) (result []string) {
	words := TrieRoot.Split(origin)
	for _, w := range words {
		result = append(result, w.Word)
	}
	return result
}

func RankedSplit(s string) (tokenized []WordResult) {
	return TrieRoot.Split(s)
}
