package pig

import (
	"strings"
	"unicode"
)

const (
	vowels      = "aeiou"
	punctuation = ",./&!?:;\\|/!@#$%^&*()-_+"
)

type translator struct {
	VowelsSfx     string
	ConsonantsSfx string
	IgnoreE       bool
}

func New(vowelSfx, consonantsSfx string, ignoreE bool) *translator {
	return &translator{
		VowelsSfx:     vowelSfx,
		ConsonantsSfx: consonantsSfx,
		IgnoreE:       ignoreE,
	}
}

func isFirstVowel(s string) bool {
	if strings.Contains(vowels, string(s[0])) {
		return true
	}
	return false
}

// take all consonants from words begin
func separateCons(s string) string {
	var cons []string
	if isFirstVowel(s) {
		return s
	}
	if !isFirstVowel(s) {
		for i, k := range s {
			if !strings.ContainsRune(vowels, k) {
				// 121 == 'y'
				if i > 1 && k == 121 {
					break
				}
				cons = append(cons, string(k))
			} else {
				break
			}
		}
	}
	return strings.Join(cons, "")
}

func (cfg *translator) Translate(textfields ...string) []string {
	var resultContainer []string
	var stagingContainer []string
	var cons string
	var punc string
	var isFirstTitle bool

	for _, words := range textfields {
		var splitted = strings.Fields(words)

		for _, word := range splitted {

			if unicode.IsUpper(rune(word[0])) {
				isFirstTitle = true
			}

			for _, letter := range word {
				if strings.ContainsAny(punctuation, string(letter)) {
					punc = punc + string(letter)
					word = word[:len(word)-len(punc)]
				}
			}

			word = strings.TrimSpace(word)
			if cfg.IgnoreE &&
				word[len(word)-1:] == "e" &&
				word[len(word)-2:len(word)-1] != "e" {

				word = word[:len(word)-1]

			}

			if !isFirstVowel(word) {
				cons = separateCons(word)
				word = word[len(cons):] + cons + cfg.ConsonantsSfx
			} else {
				word = word+ cfg.VowelsSfx
			}

			if punc != "" {
				word = word + punc
				punc = punc[0:0]
			}
			word = strings.ToLower(word)
			if isFirstTitle {
				word = strings.Title(word)
				isFirstTitle = false
			}
			stagingContainer = append(stagingContainer, word)
		}

		resultContainer = append(resultContainer, strings.Join(stagingContainer, " "))
		stagingContainer = stagingContainer[0:0]
	}

	return resultContainer
}