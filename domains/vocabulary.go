package domains

import "context"

type Vocabulary struct {
	Title        string
	Definition   string
	Example      string
	PartOfSpeech string
	IsMemorized  string
}

type VocabRepository interface {
	AddNewVocabulary(ctx context.Context, vocab *Vocabulary) (uint, error)
	FetchAllVocabularies(ctx context.Context) ([]*Vocabulary, error)
	FetchVocabularyById(ctx context.Context, id string) (*Vocabulary, error)
	UpdateVocabularyById(ctx context.Context, id string, vocab *Vocabulary) (uint, error)
	DeleteVocabularyById(ctx context.Context, id string) (uint, error)
}
