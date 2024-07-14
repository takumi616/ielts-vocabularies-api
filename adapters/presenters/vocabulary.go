package presenters

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/takumi616/ielts-vocabularies-api/adapters/dto"
	"github.com/takumi616/ielts-vocabularies-api/adapters/presenters/utils"
	"github.com/takumi616/ielts-vocabularies-api/domains"
)

type VocabPresenter struct {
}

// Write http response data
func (p *VocabPresenter) WriteVocabIdResp(ctx context.Context, vocabID uint, w http.ResponseWriter) {
	utils.CreateHeader(w, http.StatusOK)

	//set response data to dto
	vocabIdDto := &dto.VocabIdDto{
		VocabularyID: vocabID,
	}

	if err := json.NewEncoder(w).Encode(vocabIdDto); err != nil {
		log.Printf("failed to write http response using response writer: %v", err)
	}
}

func (p *VocabPresenter) WriteVocabularyResp(ctx context.Context, vocab domains.Vocabulary, w http.ResponseWriter) {
	utils.CreateHeader(w, http.StatusOK)

	vocabDto := dto.FromDomain(vocab)
	if err := json.NewEncoder(w).Encode(vocabDto); err != nil {
		log.Printf("failed to write http response using response writer: %v", err)
	}
}
