package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"tasks-api/configs"
	"tasks-api/internal/auth/jwtpkg"
	"tasks-api/internal/infra/database"
	"tasks-api/internal/infra/database/queries"
	"tasks-api/internal/infra/messaging/rabbitmqpkg"
	"tasks-api/internal/infra/notify"
	"tasks-api/internal/infra/repository"
	"tasks-api/internal/infra/web"
	"tasks-api/internal/infra/web/server"
	"tasks-api/internal/validation"

	"time"
)

// @title Swagger Tasks API
// @version 1.0
// @description This is an api for managing tasks
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	envs := configs.LoadEnv()

	messagingConnection := rabbitmqpkg.NewConnection(envs).Dial()
	defer messagingConnection.Close()

	messaging := rabbitmqpkg.NewMessaging(messagingConnection)

	db := database.New(envs)

	defer db.Close()

	qrs := queries.New(db)

	uRepo := repository.NewUserRepository(qrs)

	v := validation.NewWrapper()

	jwtService := jwtpkg.NewJWTService(envs.JwtSecret)

	notifyService := notify.NewSimpleNotifier(messaging)

	httpHandler := server.StartHttpHandler(&server.HandlersContainer{
		UserHandler:       *web.NewUserHandler(envs, uRepo, jwtService, v),
		TechnicianHandler: *web.NewTechnicianHandler(envs, repository.NewTechnicianRepository(qrs), v, notifyService),
		ManagerHandler:    *web.NewManagerHandler(envs, repository.NewManagerRepository(qrs), v),
	}, envs.WebServerPort)

	s := server.NewServer(envs, httpHandler)

	_ = chi.Walk(httpHandler, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf("[%s]: '%s' has %d middlewares\n", method, route, len(middlewares))
		return nil
	})
	// Create a done channel to signal when the shutdown is complete
	done := make(chan bool, 1)

	// Run graceful shutdown in a separate goroutine
	go gracefulShutdown(s, messagingConnection, done)

	err := s.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(fmt.Sprintf("http server error: %s", err))
	}

	// Wait for the graceful shutdown to complete
	<-done
	log.Println("Graceful shutdown complete.")
}

func gracefulShutdown(apiServer *http.Server, rabbitConn *amqp.Connection, done chan bool) {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := apiServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	if err := rabbitConn.Close(); err != nil {
		log.Printf("RabbitMQ connection forced to close with error: %v", err)
	}

	log.Println("Server exiting")

	// Notify the main goroutine that the shutdown is complete
	done <- true
}
