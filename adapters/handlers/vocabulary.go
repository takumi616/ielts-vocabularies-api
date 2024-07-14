package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/takumi616/ielts-vocabularies-api/adapters/dto"
	"github.com/takumi616/ielts-vocabularies-api/domains"
	"github.com/takumi616/ielts-vocabularies-api/usecases/ports"
)

type VocabHandler struct {
	VocabInputPort ports.VocabInputPort
}

func (h *VocabHandler) AddNewVocabulary(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	//Decode request body
	var vocabReq dto.VocabDto
	if err := json.NewDecoder(r.Body).Decode(&vocabReq); err != nil {
		err = errors.New(fmt.Sprintf("failed to decode request body. %v", err))
		//Write error response.
		h.VocabInputPort.AddNewVocabulary(ctx, domains.Vocabulary{}, w, err)
		return
	}

	//Call services's method through inputport interface
	vocabulary := dto.ToDomain(vocabReq)
	h.VocabInputPort.AddNewVocabulary(ctx, vocabulary, w, nil)
}

func (h *VocabHandler) FetchVocabularyById(w http.ResponseWriter, r *http.Request) {
	h.VocabInputPort.FetchVocabularyById(r.Context(), r.PathValue("id"), w)
}
