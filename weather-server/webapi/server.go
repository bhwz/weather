package webapi

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type ServerConfig struct {
	Debug    bool
	HttpPort string
}

type server struct {
	config ServerConfig
	router *Router
}

func Start(config ServerConfig) {
	srv := &server{
		config: config,
		router: NewRouter(),
	}

	// See: routes.go
	srv.setRoutes()

	httpServer := &http.Server{
		Addr:    ":" + srv.config.HttpPort,
		Handler: http.HandlerFunc(srv.router.ServeHTTP),
	}

	quit := make(chan os.Signal, 1)
	// Listen for interrupt/SIGINT (^C) and SIGTERM (sent by Docker).
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("Starting HTTP server on port " + srv.config.HttpPort)
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Println(err)
		}
	}()

	sig := <-quit

	log.Println("Received ", sig, " signal. Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	err := httpServer.Shutdown(ctx)
	if err != nil {
		log.Println("Received error on http server shutdown: ", err)
	}

	// Database cleanups and such go here...

	log.Println("Exiting.")
	os.Exit(0)
}
