package main

import (
	"context"
	"fmt"
	"log/slog"
	"math/rand/v2"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

type RequestId string

const RequestIdKey = RequestId("requestId")

type UserSession struct {
	userId     *uuid.UUID
	expiryTime time.Time
}

type ApiServer struct {
	env                *EnvConfig
	locationSvc        LocationService
	userSvc            UserService
	memoryUserSessions map[string]UserSession
}

func NewApiServer(env *EnvConfig, locationSvc LocationService, userSvc UserService) *ApiServer {
	return &ApiServer{
		env, locationSvc, userSvc, map[string]UserSession{},
	}
}

func (svr *ApiServer) UnhandledError(e error) {
	slog.Error("Unhandled error serving request", "error", e)
}

func (svr *ApiServer) Run() error {
	router := http.NewServeMux()

	// router for search country place
	router.HandleFunc("/searchCountry", svr.gptSearchCountry)

	// router for search place details
	router.HandleFunc("/searchPlace", svr.gptSearchPlaceDetails)

	// router for receive login details
	router.HandleFunc("/api/login", svr.userLoginHandler)

	// Router for receive signup details
	router.HandleFunc("/api/signup", svr.userSignupHandler)

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
		/* #nosec */
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
