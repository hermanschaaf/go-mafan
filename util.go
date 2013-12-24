package mafan

import "unicode"

func IsHanzi(s string) bool {
	isChinese := true
	runeForm := []rune(s)
	for _, r := range runeForm {
		isChinese = isChinese && unicode.IsOneOf([]*unicode.RangeTable{
			unicode.Unified_Ideograph,
		}, r)
		if !isChinese {
			return isChinese
		}
	}
	return isChinese
}

func IsChinese(s string) bool {
	isChinese := true
	runeForm := []rune(s)
	for _, r := range runeForm {
		isChinese = isChinese && unicode.IsOneOf([]*unicode.RangeTable{
			unicode.Unified_Ideograph,
			unicode.Terminal_Punctuation,
			unicode.Hyphen,
			unicode.Diacritic,
			unicode.White_Space,
		}, r)
		if !isChinese {
			return isChinese
		}
	}
	return isChinese
}
