package app

import (
	"api_book/middleware"
	"net/http"
)

func NewServer(authMiddleare *middleware.AuthMiddleware) *http.Server {
	return  &http.Server{
		Addr: "localhost:9090",
		Handler: authMiddleare,
	}
}