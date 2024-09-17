package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	oidc_copycat "github.com/verkaufer/oidc-copycat"
	"github.com/verkaufer/oidc-copycat/server"
	bitcaskdb "github.com/verkaufer/oidc-copycat/storage/bitcask"
	"go.mills.io/bitcask/v2"
)

// TODO: read from env
const SERVE_PORT = ":8080"

////// TODO /////
/// Follow the cmd pattern from here https://github.com/benbjohnson/wtf/blob/05bc90c940d5f9e2490fc93cf467d9e8aa48ad63/cmd/wtfd/main.go
// We need to inject the services
///

func main() {

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")
	router.Static("/public/assets", "./assets")

	db, err := bitcask.Open("/tmp/db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	store := bitcaskdb.New(db)

	directoryService := oidc_copycat.NewDirectoryService(store)

	server.RegisterRoutes(router, directoryService)

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
