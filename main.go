package main

import (
	"fmt"
	"strings"
)

const (
	vowels = "aeiou"
)

type Config struct {
	VowelsSfx     string
	ConsonantsSfx string
	IgnoreE       bool
}

// todo добавить возможность перевода цельного куска текста.
func main() {
	config := &Config{
		VowelsSfx:     "hay",
		ConsonantsSfx: "ay",
		IgnoreE:       true,
	}
	//var word = "yeild"
	var text = "This is text, to test translator!"
	fmt.Println(Translate(config, text))
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
				if
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

func Translate(cfg *Config, words string) []string {
	var result []string
	var splited = strings.Split(words, " ")
	for _, word := range splited {
		word = strings.TrimSpace(word)
		if cfg.IgnoreE &&
			word[len(word)-1:] == "e" &&
			word[len(word)-2:len(word)-1] != "e" {

			word = word[:len(word)-1]

		}
		if !isFirstVowel(word) {
			cons := separateCons(word)
			word = word[len(cons):] + cons + cfg.ConsonantsSfx

		} else {
			word = word + cfg.VowelsSfx
		}

		result = append(result, word)
	}
	return result
}
