package mafan

import (
	"io/bufio"
	"strings"
)

var TrieRoot *Trie = &Trie{map[string]*Trie{}, "", []string{}}

func init() {
	file, _ := os.Open("data/dict.txt.big")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		fmt.Println(parts[0])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
