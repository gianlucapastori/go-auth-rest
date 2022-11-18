package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gianlucapastori/nausicaa/cfg"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Server struct {
	mux   *mux.Router
	db    *sqlx.DB
	cfg   *cfg.Config
	sugar *zap.SugaredLogger
}

func (s *Server) New(db *sqlx.DB, cfg *cfg.Config, sugar *zap.SugaredLogger) *Server {
	return &Server{mux: mux.NewRouter(), db: db, cfg: cfg, sugar: sugar}
}

func (s *Server) Serve() error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", s.cfg.SERVER.PORT),
		Handler: s.mux,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Panicf("error while listening and serving: %v", err)
			}
		}
	}()

	if err := s.Map(); err != nil {
		return fmt.Errorf("could not map api: %v", err)
	}

	s.sugar.Infof("listening and serving on port %s in %s mode", s.cfg.SERVER.PORT, s.cfg.SERVER.ENV)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	s.sugar.Infof("server exited...")
	return server.Shutdown(ctx)
}
