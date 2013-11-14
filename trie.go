package mafan

type Trie struct {
	children map[string]*Trie
	letter   string
	end      bool
}

func newTrie() *Trie {
	/*
	   Build a trie for efficient retrieval of entries
	*/
	var root *Trie = &Trie{map[string]*Trie{}, "", false}
	return root
}

func (t *Trie) Insert(letters string) {
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
			t.children[letter_str] = &Trie{map[string]*Trie{}, "", false}
			t = t.children[letter_str]
		}

		if l == len(letters_rune)-1 {
			// last letter, save ending and exit
			t.end = true
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
			found = t.end
		}
	}
	return found
}

func (t *Trie) Split(origin string) (result []string) {
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
		path := ""
		depth := 0
		for i := 0; i+l < len(origin_rune); i++ {
			letter := string(origin_rune[l+i])
			path += letter
			if t.children[letter] == nil {
				// not found
				break
			} else {
				if t.children[letter].end == true {
					found_value = path
					depth = i
				}
				t = t.children[letter]
			}
		}
		if found_value != "" {
			result = append(result, found_value)
			l += depth
		} else {
			result = append(result, string(origin_rune[l:l+1]))
		}
	}
	return result
}
