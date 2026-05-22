package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/shaurya21343/go-url-shortner/internal/config"
	"github.com/shaurya21343/go-url-shortner/internal/database"
	httphandelers "github.com/shaurya21343/go-url-shortner/internal/http-handelers"
	_ "modernc.org/sqlite"
)

func main() {
	cfg, err := config.NewConfig()

	if err != nil {
		panic("error in config")
	}
	database.Connect(cfg)

	router := http.NewServeMux()

	router.HandleFunc("POST /api/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	})
	router.HandleFunc("POST /api/createShortUrl", httphandelers.CreateShortUrl)

	done := make(chan os.Signal, 1)

	server := http.Server{
		Addr:    cfg.HTTPServer.Address,
		Handler: router,
	}

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Starting server on", cfg.HTTPServer.Address)
	go func() {

		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic("error in seting up the server")
		}

	}()

	<-done

	slog.Info("server shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = server.Shutdown(ctx)

	if err != nil {
		slog.Error("failed to shutdown server", "error", err)
	}
	slog.Info("server stopped")

}
