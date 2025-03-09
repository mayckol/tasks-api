package server

import (
	"fmt"
	"net/http"
	"tasks-api/configs"
	"time"
)

type Server struct {
	port int
	envs *configs.EnvVars
}

func NewServer(envs *configs.EnvVars, httpHandler http.Handler) *http.Server {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", envs.WebServerPort),
		Handler:      httpHandler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
