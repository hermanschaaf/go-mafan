package mafan

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

var TrieRoot *Trie = &Trie{map[string]*Trie{}, "", 0}

func SetupTrie() {
	p, err := filepath.Abs("data/dict.compressed.tab")
	if err != nil {
		panic(err)
	}

	file, err := os.Open(p)
	defer file.Close()

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	k := 1
	for scanner.Scan() {
		parts := scanner.Text()
		TrieRoot.Insert(parts, k)
		k += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	SetupTrie()
}
