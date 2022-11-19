package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sudoku/internal/logger"
	"sudoku/pkg/board"
	"syscall"
	"time"

	"github.com/julienschmidt/httprouter"
)

type App struct {
	L *logger.AppLogger
	R *httprouter.Router
}

func Run() {
	board.Init()
	app := New()
	app.Listen()
}

func New() App {
	app := App{
		L: logger.New(),
		R: httprouter.New(),
	}
	app.Register()
	return app
}

func (app *App) Register() {
	app.R.HandlerFunc(http.MethodPost, "/solve", app.SolveHandler)
}

func (app *App) NewServer() http.Server {
	return http.Server{
		Addr:     ":3000",
		Handler:  app.R,
		ErrorLog: app.L.Err,
	}
}

func (app *App) Listen() {
	srv := app.NewServer()

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			app.L.Error(fmt.Sprintf("%v", err))
		}
	}()

	// graceful shutdown

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	sig := <-sigChan // blocks
	app.L.Info(fmt.Sprintf("Shutting down server %v", sig))

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	srv.Shutdown(ctx)
}