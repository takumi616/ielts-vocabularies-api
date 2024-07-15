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
	SelectAllVocabularies(ctx context.Context) ([]*dto.VocabDto, error)
	SelectVocabularyById(ctx context.Context, vocabularyID uint) (*dto.VocabDto, error)
	UpdateVocabularyById(ctx context.Context, vocabularyID uint, vocabDto *dto.VocabDto) (uint, error)
	DeleteVocabularyById(ctx context.Context, vocabularyID uint) (uint, error)
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

func (r *VocabRepository) FetchAllVocabularies(ctx context.Context) ([]*domains.Vocabulary, error) {
	selected, err := r.Persistence.SelectAllVocabularies(ctx)
	if err != nil {
		return nil, err
	}

	var vocabularies []*domains.Vocabulary
	for _, vocabDto := range selected {
		vocabularies = append(vocabularies, dto.ToDomain(vocabDto))
	}

	return vocabularies, nil
}

func (r *VocabRepository) FetchVocabularyById(ctx context.Context, id string) (*domains.Vocabulary, error) {
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

func (r *VocabRepository) UpdateVocabularyById(ctx context.Context, id string, vocabulary *domains.Vocabulary) (uint, error) {
	vocabularyID, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalf("failed to convert string id into int type: %v", err)
	}

	updatedID, err := r.Persistence.UpdateVocabularyById(ctx, uint(vocabularyID), dto.FromDomain(vocabulary))
	if err != nil {
		return 0, err
	} else {
		return updatedID, nil
	}
}

func (r *VocabRepository) DeleteVocabularyById(ctx context.Context, id string) (uint, error) {
	vocabularyID, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalf("failed to convert string id into int type: %v", err)
	}

	deletedID, err := r.Persistence.DeleteVocabularyById(ctx, uint(vocabularyID))
	if err != nil {
		return 0, err
	} else {
		return deletedID, nil
	}
}
