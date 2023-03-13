package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	middleware "req_parallel/pkg/middleware"

	heavyapiload "req_parallel/heavyapiload"
	heavyapiloadhttp "req_parallel/heavyapiload/delivery/http"
	heavyapiloadfile "req_parallel/heavyapiload/repository"
	heavyapiloadusecase "req_parallel/heavyapiload/usecase"

	"github.com/gin-gonic/gin"
)

type App struct {
	httpServer *http.Server

	heavyApiLoadUC heavyapiload.UseCase
}

func NewApp() *App {

	heavyapiloadRepo := heavyapiloadfile.NewHeavyApiLoadRepository()

	return &App{
		heavyApiLoadUC: heavyapiloadusecase.NewHeavyApiLoadUseCase(heavyapiloadRepo),
	}
}

func (a *App) Run(port string) error {
	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	router.Use(middleware.CORSMiddleware())

	heavyapiloadhttp.RegisterHTTPEndpoints(router, a.heavyApiLoadUC)

	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    1000 * time.Second,
		WriteTimeout:   1000 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}
