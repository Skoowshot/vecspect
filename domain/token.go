package domain

import (
	"encoding/gob"
	"os"
)

type Vocabulary struct {
	IdToToken map[int]string
	TokenToId map[string]int
	NextID    int
}

func NewVocabulary() *Vocabulary {
	return &Vocabulary{
		IdToToken: make(map[int]string),
		TokenToId: make(map[string]int),
		NextID:    0,
	}
}

func (v *Vocabulary) Add(token string) int {
	if id, ok := v.TokenToId[token]; ok {
		return id
	}

	id := v.NextID

	v.IdToToken[id] = token
	v.TokenToId[token] = id

	v.NextID++

	return id
}

func (v *Vocabulary) Save(filename string) error {
	f, err := os.Create(filename + ".tmp")
	if err != nil {
		return err
	}
	defer f.Close()

	err = gob.NewEncoder(f).Encode(v)
	if err != nil {
		return err
	}

	f.Close()
	return os.Rename(filename+".tmp", filename)
}

func LoadVocabulary(filename string) (*Vocabulary, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	v := NewVocabulary()

	if err := gob.NewDecoder(f).Decode(v); err != nil {
		return nil, err
	}

	return v, err
}
