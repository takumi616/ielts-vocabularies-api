package domains

import "context"

type Vocabulary struct {
	VocabularyID uint
	Title        string
	Definition   string
	Example      string
	PartOfSpeech string
	IsMemorized  string
}

type VocabRepository interface {
	AddNewVocabulary(ctx context.Context, vocab Vocabulary) (uint, error)
}
