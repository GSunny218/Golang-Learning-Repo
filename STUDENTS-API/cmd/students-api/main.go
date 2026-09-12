package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/sunny/students-api/internal/config"
);
func main() {
	//load config
	cfg := config.MustLoad();
	//setup router
	router := http.NewServeMux();
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Students API!"));
	});
	server := http.Server {
		Addr: cfg.Addr,
		Handler: router,
	}
	slog.Info("Server started %s", slog.String("address", cfg.Addr));
	done := make(chan os.Signal, 1);
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM);
	go func() {
		err := server.ListenAndServe();
		if err != nil {
			log.Fatal("Failed to start server");
		}
		fmt.Println("Server started on", cfg.Addr);
	}()
	<- done;
	slog.Info("Shutting down server...");
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second);
	defer cancel();
	err := server.Shutdown(ctx);
	if err != nil {
		slog.Error("Failed to shutdown server", "error", slog.String("error", err.Error()));
	}
	slog.Info("Server shutdown successfully"); // Log a message indicating that the server has been shut down successfully
}