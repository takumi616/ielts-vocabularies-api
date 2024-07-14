package repositories

import (
	"context"
	"log"
	"strconv"

	"github.com/takumi616/ielts-vocabularies-api/adapters/dto"
	"github.com/takumi616/ielts-vocabularies-api/domains"
)

type VocabRepository struct {
	Persistence VocabPersistence
}

// implementation is in infrastructure database layer
type VocabPersistence interface {
	InsertNewVocabulary(ctx context.Context, vocabDto *dto.VocabDto) (uint, error)
	SelectVocabularyById(ctx context.Context, vocabularyID uint) (*dto.VocabDto, error)
}

func (r *VocabRepository) AddNewVocabulary(ctx context.Context, vocabulary *domains.Vocabulary) (uint, error) {
	vocabDto := dto.FromDomain(vocabulary)

	insertedID, err := r.Persistence.InsertNewVocabulary(ctx, vocabDto)
	if err != nil {
		return 0, err
	} else {
		return insertedID, nil
	}
}

func (r *VocabRepository) FetchVocabularyById(ctx context.Context, id string) (*domains.Vocabulary, error) {
	//Get path value
	vocabularyID, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalf("failed to convert string id into int type: %v", err)
	}

	selected, err := r.Persistence.SelectVocabularyById(ctx, uint(vocabularyID))
	if err != nil {
		return &domains.Vocabulary{}, err
	}

	return dto.ToDomain(selected), nil
}
