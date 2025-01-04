package main

import (
	"context"
	"fmt"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
	"log/slog"
	"math/rand/v2"
	"net/http"
	"time"
)

type RequestId string

const RequestIdKey = RequestId("requestId")

type ApiServer struct {
	env         *EnvConfig
	locationSvc LocationService
}

func NewApiServer(env *EnvConfig, locationSvc LocationService) *ApiServer {
	return &ApiServer{
		env, locationSvc,
	}
}

func (svr *ApiServer) Run() error {
	router := http.NewServeMux()

	// router.HandleFunc("/", handlePage)
	router.HandleFunc("/searchCountry", svr.handleSearchCountry)

	// Use CORS middleware to handle cross-origin requests
	handler := requestTelemetryMiddleware(cors.Default().Handler(router))

	server := http.Server{
		Handler:      handler,
		Addr:         fmt.Sprintf(":%d", svr.env.WebPort),
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	slog.Info("server starting", "PORT", svr.env.WebPort)

	// This blocks forever, until the server has an unrecoverable error
	return server.ListenAndServe()
}

func requestTelemetryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Generate random id for this request
		reqId := rand.Uint32()
		// Copy the original request context and create a new one with added request id
		newCtx := context.WithValue(r.Context(), RequestIdKey, reqId)
		// Recreate request with new context
		r = r.WithContext(newCtx)

		// Wrap the response writer so we can capture status code
		nw := negroni.NewResponseWriter(w)

		// Continue to handle request, afterward we will have a response and status code
		next.ServeHTTP(nw, r)

		// Basic log of request and resulting status
		slog.Info("handled request",
			"method", r.Method,
			"path", r.URL.Path,
			"requestId", reqId,
			"resStatus", nw.Status(),
			"source", r.RemoteAddr,
		)
	})
}
