package ports

import (
	"context"
	"net/http"

	"github.com/takumi616/ielts-vocabularies-api/domains"
)

type VocabInputPort interface {
	AddNewVocabulary(ctx context.Context, vocab domains.Vocabulary, w http.ResponseWriter, err error)
}

type VocabOutputPort interface {
	WriteVocabIdResp(ctx context.Context, vocabID uint, w http.ResponseWriter)
}
