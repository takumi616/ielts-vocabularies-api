package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/takumi616/go-restapi/adapters/dto"
	"github.com/takumi616/go-restapi/domains"
	"github.com/takumi616/go-restapi/usecases/ports"
)

type VocabHandler struct {
	VocabInputPort ports.VocabInputPort
}

func (h *VocabHandler) AddNewVocabulary(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	//Decode request body
	var vocabReq dto.VocabDto
	if err := json.NewDecoder(r.Body).Decode(&vocabReq); err != nil {
		//Write error response.
		err = errors.New(fmt.Sprintf("failed to decode request body. %v", err))
		h.VocabInputPort.AddNewVocabulary(ctx, &domains.Vocabulary{}, w, err)
		return
	}

	//Call interactor's method through inputport interface
	vocabulary := dto.ToDomain(&vocabReq)
	h.VocabInputPort.AddNewVocabulary(ctx, vocabulary, w, nil)
}

func (h *VocabHandler) FetchAllVocabularies(w http.ResponseWriter, r *http.Request) {
	h.VocabInputPort.FetchAllVocabularies(r.Context(), w)
}

func (h *VocabHandler) FetchVocabularyById(w http.ResponseWriter, r *http.Request) {
	//Call interactor's method through inputport interface
	h.VocabInputPort.FetchVocabularyById(r.Context(), r.PathValue("id"), w)
}

func (h *VocabHandler) UpdateVocabularyById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	//Decode request body
	var vocabReq dto.VocabDto
	var id string
	if err := json.NewDecoder(r.Body).Decode(&vocabReq); err != nil {
		//Write erro message
		err := errors.New(fmt.Sprintf("failed to decode reqeust body. %v", err))
		h.VocabInputPort.UpdateVocabularyById(ctx, id, &domains.Vocabulary{}, w, err)
		return
	}

	id = r.PathValue("id")

	//Call interactor's method through inputport interface
	vocabulary := dto.ToDomain(&vocabReq)
	h.VocabInputPort.UpdateVocabularyById(ctx, id, vocabulary, w, nil)
}

func (h *VocabHandler) DeleteVocabularyById(w http.ResponseWriter, r *http.Request) {
	//Call interactor's method through inputport interface
	h.VocabInputPort.DeleteVocabularyById(r.Context(), r.PathValue("id"), w)
}
