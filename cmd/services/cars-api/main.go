package main

import (
	"net/http"
	"os"
	"path/filepath"

	"compile-and-run-sandbox/internal/files"
	"compile-and-run-sandbox/internal/parser"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"

	enTranslations "github.com/go-playground/validator/v10/translations/en"

	"compile-and-run-sandbox/internal/queue"
	"compile-and-run-sandbox/internal/repository"
	"compile-and-run-sandbox/internal/routing"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func getTranslator() ut.Translator {
	english := en.New()
	uni := ut.New(english, english)
	translator, _ := uni.GetTranslator("en")

	return translator
}

func main() {
	log.Info().Msg("starting cars-api")
	args := parser.ParseDefaultConfigurationArguments()

	producer, err := queue.NewNsqProducer(&queue.NsqParams{
		NsqLookupAddress: args.NsqAddress,
		NsqLookupPort:    args.NsqPort,
	})

	if err != nil {
		log.Fatal().Err(err)
	}

	repo, err := repository.NewRepository(args.DatabaseConn)
	localFileHandler := files.NewLocalFileHandler(filepath.Join(os.TempDir(), "executions"))

	if err != nil {
		log.Fatal().Err(err)
	}

	r := mux.NewRouter()

	validate := validator.New()
	translator := getTranslator()

	// register the validator with the translator to get clean readable generated
	// error messages from validation actions. This massively simplifies returning
	// error messages in the future.
	_ = enTranslations.RegisterDefaultTranslations(validate, translator)

	r.Handle("/", handlers.
		LoggingHandler(os.Stdout, routing.CompilerHandler{
			FileHandler: localFileHandler,
			Repo:        repo,
			Publisher:   producer,
			Translator:  translator,
			Validator:   validate,
		})).
		Methods("POST")

	r.Handle("/{id}", handlers.
		LoggingHandler(os.Stdout, routing.CompilerInfoHandler{
			FileHandler: localFileHandler,
			Repo:        repo,
		})).Methods("GET")

	log.Info().Msg("listening on :8080")
	log.Fatal().Err(http.ListenAndServe(":8080", handlers.CompressHandler(r)))
}
