package main

import (
	"context"
	"log"
	"log/slog"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/tomanagle/url-shortener/internal/handlers"
	"github.com/tomanagle/url-shortener/store/dbstore"
)

func generateSlug() string {

	const charSet = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, 6)

	for i := range result {
		result[i] = charSet[rand.Intn(len(charSet))]
	}

	return string(result)
}

func main() {

	r := chi.NewRouter()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	serverCtx, serverStopCtx := context.WithCancel(context.Background())
	killSig := make(chan os.Signal, 1)
	signal.Notify(killSig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	svr := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	shortURLStore := dbstore.NewShortURLStore(dbstore.NewShortURLStoreParams{
		Logger: logger,
	})

	go func() {
		sig := <-killSig

		logger.Info("got kill signal - shutting down", slog.String("signal", sig.String()))

		shutdownCtx, cancel := context.WithTimeout(serverCtx, 5*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("shutdown deadline exceeded")
			}
		}()

		err := svr.Shutdown(shutdownCtx)

		if err != nil {
			log.Fatal(err)
		}

		serverStopCtx()
		logger.Info("server shutting down")
		cancel()
	}()

	go func() {
		err := svr.ListenAndServe()

		if err != nil {
			log.Fatal(err)
		}
	}()

	r.Get("/healthcheck", handlers.NewHealthHandler().ServeHTTP)

	r.Post("/shorturl", handlers.NewCreateShortURLHandler(handlers.CreateShortURLHandlerParams{
		ShortURLStore: shortURLStore,
		GenerateSlug:  generateSlug,
	}).ServeHTTP)

	r.Get("/{slug}", handlers.NewGetShortURLHandler(handlers.GetShortURLHandlerParams{
		ShortURLStore: shortURLStore,
	}).ServeHTTP)

	logger.Info("read to work")

	<-serverCtx.Done()
}
