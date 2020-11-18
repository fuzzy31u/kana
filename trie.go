package kana

// Trie is a trie data structure
type Trie struct {
	children map[string]*Trie
	letter   string
	values   []string
}

// Build a trie for efficient retrieval of entries
func newTrie() *Trie {
	return &Trie{map[string]*Trie{}, "", []string{}}
}

// Insert a value into the trie
func (t *Trie) insert(letters, value string) {
	lettersRune := []rune(letters)

	// loop through letters in argument word
	for l, letter := range lettersRune {

		letterStr := string(letter)

		// if letter in children
		if t.children[letterStr] != nil {
			t = t.children[letterStr]
		} else {
			// not found, so add letter to children
			t.children[letterStr] = &Trie{map[string]*Trie{}, "", []string{}}
			t = t.children[letterStr]
		}

		if l == len(lettersRune)-1 {
			// last letter, save value and exit
			t.values = append(t.values, value)
			break
		}
	}
}

// Convert a given string to the corresponding values
// in the trie. This performed in a greedy fashion,
// replacing the longest valid string it can find at any
// given point.
func (t *Trie) convert(origin string) (result string) {
	root := t
	originRune := []rune(origin)
	result = ""

	for l := 0; l < len(originRune); l++ {
		t = root
		foundVal := ""
		depth := 0
		for i := 0; i+l < len(originRune); i++ {
			letter := string(originRune[l+i])
			if t.children[letter] == nil {
				// not found
				break
			}
			if len(t.children[letter].values) > 0 {
				foundVal = t.children[letter].values[0]
				depth = i
			}
			t = t.children[letter]
		}
		if foundVal != "" {
			result += foundVal
			l += depth
		} else {
			result += string(originRune[l : l+1])
		}
	}
	return result
}

func (t *Trie) convertMulti(origin string) (results []string) {
	root := t
	originRune := []rune(origin)
	//result := ""
	var words []string
	//words:=make([][]string,0,10)

	for l := 0; l < len(originRune); l++ {
		t = root
		//foundVal := ""
		var foundVals []string

		depth := 0
		for i := 0; i+l < len(originRune); i++ {
			letter := string(originRune[l+i])
			if t.children[letter] == nil {
				// not found
				break
			}
			childrenPattern := t.children[letter].values
			if len(childrenPattern) > 0 {
				for ii, val := range childrenPattern {
					if len(childrenPattern)> len(words){
						if l==0{
							words = append(words, "")
						}else {
							//words = append(words, words[l-1])
							words = append(words, words[0])
						}
					}

					previousChars := words[ii]
					previousChars+=val
					words[ii]=previousChars

					if len(childrenPattern)< len(words){
						// plus
						diff := len(words) - len(childrenPattern)
						for iii := diff; iii < len(words); iii++ {
							previousChars := words[iii]
							previousChars+=val
							words[iii]=previousChars
						}
					}

				}
				//foundVals=t.children[letter].values
				depth = i
			}
			t = t.children[letter]
		}
		if len(foundVals)>0 {
			l += depth
		} else {
			//result += string(originRune[l : l+1])
		}
	}
	return words
}
