package domain

type CoOccurrenceMatrix map[int]map[int]int

func NewCoOccurrenceMatrix() CoOccurrenceMatrix {
	return make(CoOccurrenceMatrix)
}

func (m CoOccurrenceMatrix) AddOccurrence(id1, id2 int) {
	if m[id1] == nil {
		m[id1] = make(map[int]int)
	}
	m[id1][id2]++

	if m[id2] == nil {
		m[id2] = make(map[int]int)
	}
	m[id2][id1]++
}
