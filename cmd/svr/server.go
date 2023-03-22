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
	log.Infof("server UP and listening on PORT: %s", addr)

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
		if err := srv.Shutdown(ctx); err != nil {
			return err
		}
		log.Infoln(shutdownCompleted)

		return nil
	})

	g.Go(func() error {
		return srv.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Error(err)
		err = fmt.Errorf("listenAndServe: %w", err)

		return err
	}

	return nil
}

const (
	SIGINTMessage     = "SIGINT received (Control-C ?)"
	SIGTERMMessage    = "SIGTERM received (Deployment shutdown?)"
	shutdownStarted   = "graceful shutdown..."
	shutdownCompleted = "graceful shutdown complete"
	fifteen           = 15 * time.Second
)
