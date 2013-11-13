package mafan

import ()

func Split(s string) (tokenized []string) {
	tokenized = TrieRoot.Split(s)
	return tokenized
}
