package main

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func listenAndServe(addr string, handler http.Handler) error {
	log.Infof("server listening on PORT: %s", addr)

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

//func CorsHandler() *cors.Cors {
//	return cors.New(cors.Options{
//		AllowedOrigins:   allowedOrigins,
//		AllowCredentials: true,
//		AllowedMethods:   allowedMethods,
//		AllowedHeaders:   allowedHeaders,
//		// Enable Debugging for testing, consider disabling in production
//		Debug: false,
//	})
//}

const (
	SIGINTMessage  = "SIGINT received (Control-C ?)"
	SIGTERMMessage = "SIGTERM received (Deployment shutdown?)"

	shutdownStarted   = "graceful shutdown..."
	shutdownCompleted = "graceful shutdown complete"
	fifteen           = 15 * time.Second
)
