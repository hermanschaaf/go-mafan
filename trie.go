package mafan

type Trie struct {
	children map[string]*Trie
	letter   string
	rank     int // rank <= 0 indicates not a full word
}

func newTrie() *Trie {
	/*
	   Build a trie for efficient retrieval of entries
	*/
	var root *Trie = &Trie{map[string]*Trie{}, "", 0}
	return root
}

func (t *Trie) Insert(letters string, rank int) {
	/*
		Insert a value into the trie
	*/

	letters_rune := []rune(letters)

	// loop through letters in argument word
	for l, letter := range letters_rune {

		letter_str := string(letter)

		// if letter in children
		if t.children[letter_str] != nil {
			t = t.children[letter_str]
		} else {
			// not found, so add letter to children
			t.children[letter_str] = &Trie{map[string]*Trie{}, "", 0}
			t = t.children[letter_str]
		}

		if l == len(letters_rune)-1 {
			// last letter, save ending and exit
			t.rank = rank
			break
		}
	}
}

func (t *Trie) Search(srch string) (found bool) {
	/*
		Search for a string in the Trie.

		Returns the corresponding array of strings if found,
		or an empty array otherwise.
	*/
	found = false
	srch_rune := []rune(srch)

	for l, letter := range srch_rune {
		letter_string := string(letter)
		if t.children[letter_string] != nil {
			t = t.children[letter_string]
		} else {
			found = false
			return found
		}
		if l == len(srch_rune)-1 {
			found = t.rank >= 1
		}
	}
	return found
}

type WordResult struct {
	Word string
	Rank int
}

func (t *Trie) Split(origin string) (result []WordResult) {
	/*
		Convert a given string to the corresponding values
		in the trie. This performed in a greedy fashion,
		replacing the longest valid string it can find at any
		given point.
	*/
	root := t
	origin_rune := []rune(origin)

	for l := 0; l < len(origin_rune); l++ {
		t = root
		found_value := ""
		found_rank := 0
		path := ""
		depth := 0
		skipped := false
		for i := 0; i+l < len(origin_rune); i++ {
			letter := string(origin_rune[l+i])
			path += letter
			if (i == 0 || skipped) && !IsHanzi(letter) {
				depth = i
				skipped = true
				found_value = path
				continue
			} else if skipped {
				break
			}
			if t.children[letter] == nil {
				// not found
				break
			} else {
				if t.children[letter].rank >= 0 {
					found_value = path
					depth = i
					found_rank = t.children[letter].rank
				}
				t = t.children[letter]
			}
		}
		if found_value != "" {
			result = append(result, WordResult{found_value, found_rank})
			l += depth
		} else {
			result = append(result, WordResult{string(origin_rune[l : l+1]), -1})
		}
	}
	return result
}
