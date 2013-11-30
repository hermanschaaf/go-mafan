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

func TestRankedSplit(t *testing.T) {
	rankedSplit := RankedSplit("上海十大接吻聖地")
	expected := []string{"上海", "十大", "接吻", "聖地"}
	for i, r := range rankedSplit {
		if r.Word != expected[i] {
			t.Errorf("%s != %s", rankedSplit, expected)
		}
		if r.Rank <= 0 {
			t.Errorf("No rank returned for %s", r.Word)
		}
	}

}

func BenchmarkSetup(b *testing.B) {
	// benchmark how efficient the trie setup is
	for i := 0; i < b.N; i++ {
		SetupTrie()
	}
}

func BenchmarkSplit(b *testing.B) {
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
