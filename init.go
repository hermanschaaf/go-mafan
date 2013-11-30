package mafan

import (
	"bufio"
	"log"
	"os"
)

var TrieRoot *Trie = &Trie{map[string]*Trie{}, "", 0}

func SetupTrie() {
	file, _ := os.Open("data/dict.compressed.spaces")
	defer file.Close()

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
