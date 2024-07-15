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

func (p *VocabPresenter) WriteVocabIdResp(ctx context.Context, vocabID uint, w http.ResponseWriter) {
	//write http response header
	utils.CreateHeader(w, http.StatusOK)

	//set response body
	res := struct {
		VocabularyID uint `json:"vocabulary_id"`
	}{
		VocabularyID: vocabID,
	}

	//write response body
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("failed to write http response using response writer: %v", err)
	}
}

func (p *VocabPresenter) WriteVocabulariesResp(ctx context.Context, vocabs []*domains.Vocabulary, w http.ResponseWriter) {
	//Write http response header
	utils.CreateHeader(w, http.StatusOK)

	var vocabularies []*dto.VocabDto
	for _, vocab := range vocabs {
		vocabularies = append(vocabularies, dto.FromDomain(vocab))
	}

	if err := json.NewEncoder(w).Encode(vocabularies); err != nil {
		log.Printf("failed to write http response using response writer: %v", err)
	}
}

func (p *VocabPresenter) WriteVocabularyResp(ctx context.Context, vocab *domains.Vocabulary, w http.ResponseWriter) {
	//Write http response header
	utils.CreateHeader(w, http.StatusOK)

	//write response body
	vocabDto := dto.FromDomain(vocab)
	if err := json.NewEncoder(w).Encode(vocabDto); err != nil {
		log.Printf("failed to write http response using response writer: %v", err)
	}
}
