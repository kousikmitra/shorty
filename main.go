package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	api "github.com/kousikmitra/shorty/api"
	"github.com/kousikmitra/shorty/services"
	"github.com/kousikmitra/shorty/stores"
	mem "github.com/kousikmitra/shorty/stores/memory"
)

func main() {
	store := chooseStore()
	service := services.NewRedirectService(store)
	handler := api.NewRedirectHandler(service)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/{code}", handler.Get)
	r.Post("/", handler.Post)

	errs := make(chan error, 2)
	go func() {
		fmt.Println("Listening on port :8000")
		errs <- http.ListenAndServe(httpPort(), r)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("Terminated %s", <-errs)
}

func chooseStore() stores.RedirectStore {
	store, err := mem.NewMemoryStore(100)
	if err != nil {
		log.Fatal(err)
	}
	return store
}

func httpPort() string {
	port := "8000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return fmt.Sprintf(":%s", port)
}
