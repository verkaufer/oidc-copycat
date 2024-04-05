package main

import (
	"context"
	"embed"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

// TODO: read from env
const SERVE_PORT = ":8080"

//go:embed assets/* templates/*
var f embed.FS

func main() {

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")
	router.StaticFS("/public", http.FS(f))

	registerRoutes(router)

	srv := &http.Server{
		Addr:    SERVE_PORT,
		Handler: router,
	}

	go func() {
		log.Printf("Starting server on port %s", SERVE_PORT)
		if err := srv.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatalf("listen failed: %s\n", err)
			}
		}
	}()

	<-ctx.Done()
	stop()

	log.Println("Begining graceful shutdown. Press Ctrl+C to force.")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to close: %s", err)
	}

}
