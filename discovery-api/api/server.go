package api

import (
	"context"
	"discoveryweb/service/bookmark"
	"discoveryweb/service/email"
	"discoveryweb/service/image"
	"discoveryweb/service/location"
	"discoveryweb/service/places"
	"discoveryweb/service/session"
	"discoveryweb/service/user"
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

type ApiServer struct {
	listenPort           uint16
	locationSvc          location.LocationService
	userSvc              user.UserService
	placesService        places.PlacesService
	bookmarkPlaceService bookmark.BookmarkPlaceService
	imgSvc               image.ImageService
	emailSvc             email.EmailService
	sessionSvc           session.SessionService
}

func NewApiServer(listenPort uint16, locationSvc location.LocationService, userSvc user.UserService, placesService places.PlacesService, bookmarkPlaceService bookmark.BookmarkPlaceService, imgSvc image.ImageService, emailSvc email.EmailService, sessionSvc session.SessionService) *ApiServer {
	return &ApiServer{
		listenPort,
		locationSvc,
		userSvc,
		placesService,
		bookmarkPlaceService,
		imgSvc,
		emailSvc,
		sessionSvc,
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

	// Router for get all bookmark new place for user
	router.HandleFunc("/api/getAllBookmark", func(w http.ResponseWriter, r *http.Request) {
		svr.currentUserSessionMiddleware(http.HandlerFunc(svr.userGetAllBookmarkHandler)).ServeHTTP(w, r)
	})

	// Router for get all bookmark by city
	router.HandleFunc("/api/getAllBookmarkByCity", func(w http.ResponseWriter, r *http.Request) {
		svr.currentUserSessionMiddleware(http.HandlerFunc(svr.userGetAllBookmarkByCityHandler)).ServeHTTP(w, r)
	})
	// Router for bookmark new place for user
	router.HandleFunc("/api/bookmark", func(w http.ResponseWriter, r *http.Request) {
		svr.currentUserSessionMiddleware(http.HandlerFunc(svr.userBookmarkHandler)).ServeHTTP(w, r)
	})

	// Router for unbookmark place for user
	router.HandleFunc("/api/unBookmark", func(w http.ResponseWriter, r *http.Request) {
		svr.currentUserSessionMiddleware(http.HandlerFunc(svr.userUnBookmarkHandler)).ServeHTTP(w, r)
	})

	// Router for bookmark new place for user by place name
	router.HandleFunc("/api/bookmarkByPlaceName", func(w http.ResponseWriter, r *http.Request) {
		svr.currentUserSessionMiddleware(http.HandlerFunc(svr.userBookmarkByPlaceNameHandler)).ServeHTTP(w, r)
	})

	// Router for get user email from database
	router.HandleFunc("/api/getUserProfile", func(w http.ResponseWriter, r *http.Request) {
		svr.currentUserSessionMiddleware(http.HandlerFunc(svr.userProfileHandler)).ServeHTTP(w, r)
	})
	// Router for updating user profile picture in database
	router.HandleFunc("/api/updateUserProfileImage", func(w http.ResponseWriter, r *http.Request) {
		svr.currentUserSessionMiddleware(http.HandlerFunc(svr.userProfilePicChangeHandler)).ServeHTTP(w, r)
	})
	// Router for displaying user profile picture in database
	router.HandleFunc("/api/displayUserProfileImage", func(w http.ResponseWriter, r *http.Request) {
		svr.currentUserSessionMiddleware(http.HandlerFunc(svr.userProfilePicDisplayHandler)).ServeHTTP(w, r)
	})
	// Router for getting place image from pixel
	router.HandleFunc("/api/getDisplayPlaceImage", func(w http.ResponseWriter, r *http.Request) {
		svr.currentUserSessionMiddleware(http.HandlerFunc(svr.getPlaceImageURL)).ServeHTTP(w, r)
	})

	// router for receive login details
	router.HandleFunc("/api/login", svr.userLoginHandler)

	// Router for receive signup details
	router.HandleFunc("/api/signup", svr.userSignupHandler)

	// Router for receive forget password email
	router.HandleFunc("/api/forgetPassword", svr.userForgetPasswordHandler)

	// Router for receive new password
	router.HandleFunc("/api/resetPassword", svr.userResetPasswordHandler)

	// Router for get user email from database
	router.HandleFunc("/api/updatePassword", func(w http.ResponseWriter, r *http.Request) {
		svr.currentUserSessionMiddleware(http.HandlerFunc(svr.userUpdatePwHandler)).ServeHTTP(w, r)
	})

	// Use CORS middleware to handle cross-origin requests
	handler := requestTelemetryMiddleware((cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}).Handler(router)))

	server := http.Server{
		Handler:      handler,
		Addr:         fmt.Sprintf(":%d", svr.listenPort),
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	slog.Info("server starting", "PORT", svr.listenPort)

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

		sessionId := sessionIdFromCookie(r)

		if sessionId == nil {
			http.Error(w, "Missing session ID", http.StatusUnauthorized)
			return
		}

		validUserId, err := svr.sessionSvc.CheckAndExtendSession(*sessionId)
		if err != nil {
			http.Error(w, "Session expired or invalid", http.StatusUnauthorized)
			return
		}

		// Increase cookie expiry time
		setSectionCookie(w, sessionId.String())

		// Stores the userId in the request context, making it accessible to downstream handlers.
		newCtx := context.WithValue(r.Context(), currentUserSessionKey, validUserId)
		r = r.WithContext(newCtx)

		next.ServeHTTP(w, r)
	})
}
