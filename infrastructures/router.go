package infrastructures

import (
	"net/http"

	"github.com/takumi616/go-restapi/adapters/handlers"
)

type Router struct {
	VocabHandler *handlers.VocabHandler
}

func (r *Router) Setup() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /vocabularies", r.VocabHandler.AddNewVocabulary)
	mux.HandleFunc("GET /vocabularies", r.VocabHandler.FetchAllVocabularies)
	mux.HandleFunc("GET /vocabularies/{id}", r.VocabHandler.FetchVocabularyById)
	mux.HandleFunc("PUT /vocabularies/{id}", r.VocabHandler.UpdateVocabularyById)
	mux.HandleFunc("DELETE /vocabularies/{id}", r.VocabHandler.DeleteVocabularyById)

	return mux
}
