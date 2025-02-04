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

type contextKey string

const currentUserSessionKey contextKey = "CurrentUserSession"

type UserSession struct {
	userId     *uuid.UUID
	expiryTime time.Time
}

type ApiServer struct {
	env                  *EnvConfig
	locationSvc          LocationService
	userSvc              UserService
	memoryUserSessions   map[string]UserSession
	placesService        PlacesService
	bookmarkPlaceService BookmarkPlaceService
}

func NewApiServer(env *EnvConfig, locationSvc LocationService, userSvc UserService, placesService PlacesService, bookmarkPlaceService BookmarkPlaceService) *ApiServer {
	return &ApiServer{
		env, locationSvc, userSvc, map[string]UserSession{}, placesService, bookmarkPlaceService,
	}
}

func (svr *ApiServer) UnhandledError(e error) {
	slog.Error("Unhandled error serving request", "error", e)
}

func GetCurrentUserId(r *http.Request) *uuid.UUID {
	// Retrieve the value from the request context
	value := r.Context().Value(currentUserSessionKey)

	// Check if the value is of type uuid.UUID
	userId, ok := value.(uuid.UUID)
	if !ok || userId == uuid.Nil {
		// If the value is not a uuid.UUID or is the zero UUID, return nil
		return nil
	}
	// return the userId if it is valid
	return &userId
}

func (svr *ApiServer) Run() error {
	router := http.NewServeMux()

	// router for cities search autocomplete
	router.HandleFunc("/api/place/autocomplete", svr.autocompleteCitiesSearch)

	// router for search country place
	router.HandleFunc("/searchCountry", func(w http.ResponseWriter, r *http.Request) {
		svr.currentUserSessionMiddleware(http.HandlerFunc(svr.gptSearchCountry)).ServeHTTP(w, r)
	})

	// router for search place details
	router.HandleFunc("/searchPlace", func(w http.ResponseWriter, r *http.Request) {
		svr.currentUserSessionMiddleware(http.HandlerFunc(svr.gptSearchPlaceDetails)).ServeHTTP(w, r)
	})
	// Router for receive logout request
	router.HandleFunc("/api/logout", func(w http.ResponseWriter, r *http.Request) {
		svr.currentUserSessionMiddleware(http.HandlerFunc(svr.userLogoutHandler)).ServeHTTP(w, r)
	})

	// Router for bookmark new place for user
	router.HandleFunc("/api/bookmark", func(w http.ResponseWriter, r *http.Request) {
		svr.currentUserSessionMiddleware(http.HandlerFunc(svr.userBookmarkHandler)).ServeHTTP(w, r)
	})

	// Router for unbookmark place for user
	router.HandleFunc("/api/unBookmark", func(w http.ResponseWriter, r *http.Request) {
		svr.currentUserSessionMiddleware(http.HandlerFunc(svr.userUnBookmarkHandler)).ServeHTTP(w, r)
	})

	// router for receive login details
	router.HandleFunc("/api/login", svr.userLoginHandler)

	// Router for receive signup details
	router.HandleFunc("/api/signup", svr.userSignupHandler)

	// Router for receive forget password email
	router.HandleFunc("/api/forgetPassword", svr.userForgetPasswordHandler)

	// Router for receive new password
	router.HandleFunc("/api/resetPassword", svr.userResetPasswordHandler)

	// Use CORS middleware to handle cross-origin requests
	handler := requestTelemetryMiddleware((cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}).Handler(router)))

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

func (svr *ApiServer) currentUserSessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Log all cookies in the request for debugging purposes
		for _, cookie := range r.Cookies() {
			slog.Info("Request Cookie", "name", cookie.Name, "value", cookie.Value)
		}
		// Fetch the session ID from the request cookie
		sessionId, err := r.Cookie("DA_SESSION_ID")
		if err != nil {
			slog.Warn("Failed to get request cookie of field: 'DA_SESSION_ID'")
			http.Error(w, "Missing session ID", http.StatusUnauthorized)
			return
		}
		if sessionId == nil {
			// If there's no session ID, continue to the next handler
			next.ServeHTTP(w, r)
			return
		}

		currentUserSession, exists := svr.memoryUserSessions[sessionId.Value]
		if !exists {
			// If no user session exists for this session ID, continue to the next handler
			next.ServeHTTP(w, r)
			return
		}

		// if the session has not expired (Add userId from the UserSession)
		if currentUserSession.expiryTime.After(time.Now()) {

			// Stores the userId in the request context, making it accessible to downstream handlers.
			newCtx := context.WithValue(r.Context(), currentUserSessionKey, *currentUserSession.userId)
			r = r.WithContext(newCtx)

			//Update the expiry time in memory to extend the session
			newExpiryTime := time.Now().Add(600 * time.Second)
			currentUserSession.expiryTime = newExpiryTime
			svr.memoryUserSessions[sessionId.Value] = currentUserSession

		}

		next.ServeHTTP(w, r)
	})
}
