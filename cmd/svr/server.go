package main

import (
	"context"
	"fmt"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func listenAndServe(addr string, handler http.Handler) error {
	log.Infof("Listening on Port: %s", addr)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", addr),
		Handler:      handler,
		WriteTimeout: fifteen,
		ReadTimeout:  fifteen,
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		select {
		case killSignal := <-signals:
			switch killSignal {
			case os.Interrupt:
				log.Infoln(SIGINTMessage)
			case syscall.SIGTERM:
				log.Infoln(SIGTERMMessage)
			default:
			}
		case <-ctx.Done():
			return ctx.Err()
		}

		log.Infoln(shutdownStarted)

		if err := srv.Shutdown(context.Background()); err != nil {
			return err
		}

		log.Infoln(shutdownCompleted)

		return nil
	})

	g.Go(func() error {
		if err := srv.ListenAndServe(); err != nil {
			return err
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func corsHandler() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowCredentials: true,
		AllowedMethods:   allowedMethods,
		AllowedHeaders:   allowedHeaders,
		// Enable Debugging for testing, consider disabling in production
		Debug: false,
	})
}

var (
	allowedOrigins = []string{localhostCRA, localhostVite, localhostVite2, githubPages, githubPages1, githubPages2}
	allowedMethods = []string{"GET", "POST", "OPTIONS", "DELETE", "PUT"}
	allowedHeaders = []string{"Access-Control-Allow-Methods", "Access-Control-Allow-Origin", "X-Requested-With", "Authorization", "Content-Type", "X-Requested-With", "Bearer", "Origin"}
)

const (
	localhostCRA   = "http://localhost:3000"
	localhostVite  = "http://localhost:5173"
	localhostVite2 = "http://localhost:5173/robot-image-ui/"
	githubPages    = "https://calebtracey.github.io/robot-image-ui"
	githubPages1   = "https://calebtracey.github.io"
	githubPages2   = "https://calebtracey.github.io/robot-image-ui/"
)

const (
	SIGINTMessage  = "SIGINT received (Control-C ?)"
	SIGTERMMessage = "SIGTERM received (Deployment shutdown?)"

	shutdownStarted   = "graceful shutdown..."
	shutdownCompleted = "graceful shutdown complete"
	fifteen           = 15 * time.Second
)
