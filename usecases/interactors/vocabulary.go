package interactors

import (
	"context"
	"net/http"

	"github.com/takumi616/ielts-vocabularies-api/domains"
	"github.com/takumi616/ielts-vocabularies-api/usecases/ports"
)

type VocabInteractor struct {
	Repo            domains.VocabRepository
	VocabOutputPort ports.VocabOutputPort
	ErrOutputPort   ports.ErrOutputPort
}

func (i *VocabInteractor) AddNewVocabulary(ctx context.Context, vocab *domains.Vocabulary, w http.ResponseWriter, err error) {
	//check if error is found in handlers
	if err != nil {
		i.ErrOutputPort.WriteErrResp(ctx, err, w, http.StatusBadRequest)
		return
	}

	//Execute domain logic
	addedID, err := i.Repo.AddNewVocabulary(ctx, vocab)
	if err != nil {
		i.ErrOutputPort.WriteErrResp(ctx, err, w, http.StatusInternalServerError)
		return
	}

	//Write http response
	i.VocabOutputPort.WriteVocabIdResp(ctx, addedID, w)
}

func (i *VocabInteractor) FetchVocabularyById(ctx context.Context, id string, w http.ResponseWriter) {
	//Execute domain logic
	vocab, err := i.Repo.FetchVocabularyById(ctx, id)
	if err != nil {
		i.ErrOutputPort.WriteErrResp(ctx, err, w, http.StatusInternalServerError)
		return
	}

	//Write http response
	i.VocabOutputPort.WriteVocabularyResp(ctx, vocab, w)
}

func (i *VocabInteractor) UpdateVocabularyById(ctx context.Context, id string, vocab *domains.Vocabulary, w http.ResponseWriter, err error) {
	//check if error is found in handlers
	if err != nil {
		i.ErrOutputPort.WriteErrResp(ctx, err, w, http.StatusBadRequest)
		return
	}

	updatedID, err := i.Repo.UpdateVocabularyById(ctx, id, vocab)
	if err != nil {
		i.ErrOutputPort.WriteErrResp(ctx, err, w, http.StatusInternalServerError)
		return
	}

	//Write http response
	i.VocabOutputPort.WriteVocabIdResp(ctx, updatedID, w)
}

func (i *VocabInteractor) DeleteVocabularyById(ctx context.Context, id string, w http.ResponseWriter) {
	deletedID, err := i.Repo.DeleteVocabularyById(ctx, id)
	if err != nil {
		i.ErrOutputPort.WriteErrResp(ctx, err, w, http.StatusInternalServerError)
		return
	}

	i.VocabOutputPort.WriteVocabIdResp(ctx, deletedID, w)
}
