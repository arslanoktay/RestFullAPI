package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (app *application) mount() http.Handler {

	r := chi.NewRouter()

	r.Use(middleware.RequestID) // request contextine request id atar
	r.Use(middleware.RealIP)    // ?
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer) // panicleri kurtarması için

	r.Use(middleware.Timeout(50 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healtCheckHandler)
	})

	return r
}

func (app *application) run(mux http.Handler) error {

	srv := http.Server{ // server yaratırken 1 handler ve 1 port gerekli
		Addr:         app.config.addr,
		Handler:      mux,              // Servemux gelen url ile karşılayacak handler eşleştirmeye yarar
		WriteTimeout: time.Second * 30, // error yapısı için
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("server had started at %s.", srv.Addr)

	return srv.ListenAndServe()
}
