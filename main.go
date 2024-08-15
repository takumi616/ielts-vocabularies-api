package main

import (
	"context"
	"log"

	"github.com/takumi616/go-restapi/adapters/handlers"
	"github.com/takumi616/go-restapi/adapters/presenters"
	"github.com/takumi616/go-restapi/adapters/repositories"
	"github.com/takumi616/go-restapi/infrastructures"
	"github.com/takumi616/go-restapi/infrastructures/database"
	"github.com/takumi616/go-restapi/usecases/interactors"
)

func run(ctx context.Context) error {
	//Get config
	config, _ := infrastructures.NewConfig()

	//Initialize postgres db with Gorm
	gorm, err := database.Open(ctx, config.PgConfig)
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}

	//Initialize repository
	vocabRepository := &repositories.VocabRepository{Persistence: gorm}

	//Initialize presenter
	vocabPresenter := &presenters.VocabPresenter{}
	errPresenter := &presenters.ErrPresenter{}

	//Initialize interactor with repository and presenter
	vocabInteractor := &interactors.VocabInteractor{Repo: vocabRepository, VocabOutputPort: vocabPresenter, ErrOutputPort: errPresenter}

	//Initialize handler with service
	vocabHandler := &handlers.VocabHandler{VocabInputPort: vocabInteractor}

	router := infrastructures.Router{VocabHandler: vocabHandler}
	mux := router.Setup()

	server := infrastructures.HttpServer{Port: config.AppPort, ServeMux: mux}
	return server.Run(ctx)
}

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatalf("Golang server does not work correctly: %v", err)
	}
}
