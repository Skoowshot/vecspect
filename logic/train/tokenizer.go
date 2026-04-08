package train

import (
	"regexp"
	"strings"
)

var splitPattern = regexp.MustCompile(`(<[A-Z0-9_]+>|[a-zа-яё0-9]+('[a-z]+)?|\n|[[:punct:]])`)

type Tokenizer struct{}

func NewTokenizer() *Tokenizer {
	return &Tokenizer{}
}

func (t *Tokenizer) PrepareString(text string) string {
	lowered := strings.ToLower(text)
	return lowered
}

func (t *Tokenizer) Tokenize(text string) []string {
	prepared := t.PrepareString(text)
	split := splitPattern.FindAllString(prepared, -1)

	if split == nil {
		return []string{}
	}

	return split
}
