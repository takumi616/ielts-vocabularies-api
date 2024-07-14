package repositories

import (
	"context"

	"github.com/takumi616/ielts-vocabularies-api/adapters/dto"
	"github.com/takumi616/ielts-vocabularies-api/domains"
)

type VocabRepository struct {
	Persistence VocabPersistence
}

type VocabPersistence interface {
	InsertNewVocabulary(ctx context.Context, vocabDto dto.VocabDto) (dto.VocabIdDto, error)
}

func (r *VocabRepository) AddNewVocabulary(ctx context.Context, vocabulary domains.Vocabulary) (uint, error) {
	vocabDto := dto.FromDomain(vocabulary)

	vocabIdDto, err := r.Persistence.InsertNewVocabulary(ctx, vocabDto)
	if err != nil {
		return 0, err
	}

	vocabularyID := vocabIdDto.VocabularyID
	return vocabularyID, nil
}
