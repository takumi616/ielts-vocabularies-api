package ports

import (
	"context"
	"net/http"

	"github.com/takumi616/go-restapi/domains"
)

type VocabInputPort interface {
	AddNewVocabulary(ctx context.Context, vocab *domains.Vocabulary, w http.ResponseWriter, err error)
	FetchAllVocabularies(ctx context.Context, w http.ResponseWriter)
	FetchVocabularyById(ctx context.Context, id string, w http.ResponseWriter)
	UpdateVocabularyById(ctx context.Context, id string, vocab *domains.Vocabulary, w http.ResponseWriter, err error)
	DeleteVocabularyById(ctx context.Context, id string, w http.ResponseWriter)
}

type VocabOutputPort interface {
	WriteVocabIdResp(ctx context.Context, vocabID uint, w http.ResponseWriter)
	WriteVocabulariesResp(ctx context.Context, vocab []*domains.Vocabulary, w http.ResponseWriter)
	WriteVocabularyResp(ctx context.Context, vocab *domains.Vocabulary, w http.ResponseWriter)
}
