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

}
