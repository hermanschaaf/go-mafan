package mafan

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	split := strings.Join(Split("上海十大接吻聖地"), " ")
	expected := "上海 十大 接吻 聖地"
	if split != expected {
		t.Errorf("%s != %s", split, expected)
	}
}

func BenchmarkSetup(b *testing.B) {
	// benchmark how efficient the trie setup is
	for i := 0; i < b.N; i++ {
		SetupTrie()
	}
}

func BenchmarkSearch(b *testing.B) {
	// benchmark how efficient searching the trie is
	// for some short strings
	s := []string{
		"上海十大接吻聖地",
		"而能夠操作這種內化影響",
		"這是我見過最潮的髮型了",
	}
	for i := 0; i < b.N; i++ {
		Split(s[i%len(s)])
	}
}
