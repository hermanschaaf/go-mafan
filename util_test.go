package mafan

import (
	"testing"
)

func TestIsHanzi(t *testing.T) {
	cases := map[string]bool{
		"english":       false,
		"_random chars": false,
		"上海十大接吻聖地":      true,
		"上海十大接吻聖地！":     false,
	}
	for c := range cases {
		v := IsHanzi(c)
		if v != cases[c] {
			t.Errorf("Expected IsHanzi(%v) = %v, but got %v", c, cases[c], v)
		}
	}
}

func TestIsChinese(t *testing.T) {
	cases := map[string]bool{
		"english":       false,
		"_random chars": false,
		"上海十大接吻聖地":      true,
		"上海十大接吻聖地！":     true,
	}
	for c := range cases {
		v := IsChinese(c)
		if v != cases[c] {
			t.Errorf("Expected IsChinese(%v) = %v, but got %v", c, cases[c], v)
		}
	}
}
