package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/KadirbekSharau/apprentice-backend/configs"
	"github.com/KadirbekSharau/apprentice-backend/internal/auth"
	authHttp "github.com/KadirbekSharau/apprentice-backend/internal/auth/delivery/http"
	authRepository "github.com/KadirbekSharau/apprentice-backend/internal/auth/repository"
	authUsecase "github.com/KadirbekSharau/apprentice-backend/internal/auth/usecase"

	"github.com/KadirbekSharau/apprentice-backend/internal/employer"
	employerHttp "github.com/KadirbekSharau/apprentice-backend/internal/employer/delivery/http"
	employerRepository "github.com/KadirbekSharau/apprentice-backend/internal/employer/repository"
	employerUsecase "github.com/KadirbekSharau/apprentice-backend/internal/employer/usecase"

	jobpost "github.com/KadirbekSharau/apprentice-backend/internal/job_post"
	jobpostHttp "github.com/KadirbekSharau/apprentice-backend/internal/job_post/delivery/http"
	jobpostRepository "github.com/KadirbekSharau/apprentice-backend/internal/job_post/repository"
	jobpostUsecase "github.com/KadirbekSharau/apprentice-backend/internal/job_post/usecase"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type App struct {
	httpServer *http.Server
	authUC     auth.UseCase
	employerUC employer.UseCase
	jobpostUC  jobpost.UseCase
}

func NewApp() *App {
	db, err := configs.NewPostgresDB()
	if err != nil {
		log.Fatal("Can't connect to the database. Error: %s", err.Error())
	}

	authRepository := authRepository.NewUserRepository(db)
	employerRepository := employerRepository.NewEmployerRepository(db)
	jobPostRepository := jobpostRepository.NewJobPostRepository(db)

	return &App{
		authUC: authUsecase.NewAuthUseCase(
			authRepository,
			os.Getenv("HASH_SALT"),
			[]byte(os.Getenv("SIGNING_KEY")),
			viper.GetDuration("auth.token_ttl"),
		),
		employerUC: employerUsecase.NewEmployerUseCase(
			employerRepository,
		),
		jobpostUC: jobpostUsecase.NewJobPostUseCase(
			jobPostRepository,
		),
	}
}

func (a *App) Run(port string) error {
	// Init Gin Handler
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	// Registering API endpoints
	authHttp.RegisterHTTPEndpoints(router, a.authUC)
	employerHttp.RegisterHTTPEndpoints(router, a.employerUC)
	jobpostHttp.RegisterHTTPEndpoints(router, a.jobpostUC)

	// HTTP Server
	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Server run
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
