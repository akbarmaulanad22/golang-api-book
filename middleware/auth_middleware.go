package middleware

import (
	"api_book/helper"
	"api_book/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	if apiKey := r.Header.Get("x-api-key"); apiKey == "KEPOBEJIR" {
		middleware.Handler.ServeHTTP(w, r)
	} else {
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code: http.StatusUnauthorized,
			Message: "UNAUTHORIZED",
			Data: "API KEY not match",
		}

		helper.WriteToResponse(w, webResponse)
	}
}
