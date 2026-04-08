package domain

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
