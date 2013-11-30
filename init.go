package mafan

import (
	"bufio"
	"log"
	"os"
)

var TrieRoot *Trie = &Trie{map[string]*Trie{}, "", false}

func SetupTrie() {
	file, _ := os.Open("data/dict.compressed.spaces")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := scanner.Text()
		TrieRoot.Insert(parts)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	SetupTrie()
}
