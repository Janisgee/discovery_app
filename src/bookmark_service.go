package main

import (
	"context"
	"discoveryapp/internal/database"
	"encoding/json"
	"errors"
	"log"
	"log/slog"

	"github.com/google/uuid"
)

type BookmarkPlaceService interface {
	CreatePlaceData(place_id string, place_name string, country string, city string, catagory string, place_detail *PlaceDetails) error
	GetBookmarkPlaceDetails(place_id string) (*Place, error)
	CreateUserBookmark(user_id uuid.UUID, username string, place_id string, place_name string, place_text string) (*UserBookmarkPlace, error)
}

// type PlaceDetail struct {
// 	Description  string `json:"description"`
// 	Address      string `json:"address"`
// 	OpeningHours string `json:"history"`
// 	KeyFeatures  string `json:"key_feature"`
// }

// Place struct to hold input
type Place struct {
	ID          string       `json:"id"`
	PlaceName   string       `json:"place_name"`
	Country     string       `json:"country"`
	City        string       `json:"city"`
	Category    string       `json:"category"`
	PlaceDetail PlaceDetails `json:"place_detail"`
}

// userBookmark struct
type UserBookmarkPlace struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Username  string    `json:"username"`
	PlaceID   string    `json:"place_id"`
	PlaceName string    `json:"place_name"`
	PlaceText string    `json:"place_text"`
}

type PostgresBookmarkService struct {
	dbQueries *database.Queries
}

func (svc *PostgresBookmarkService) CreatePlaceData(place_id string, place_name string, country string, city string, catagory string, place_details *PlaceDetails) error {
	// Create an empty context
	ctx := context.Background()

	// Turn the place details into bytes for the database
	details, err := json.Marshal(place_details)
	if err != nil {
		slog.Warn("Fail to marshal place detail", "error", err)
		return errors.New("fail to marshal place detail")
	}

	// Check if place inside database
	params := database.CreatePlaceParams{
		ID:          place_id,
		PlaceName:   place_name,
		Country:     country,
		City:        city,
		Category:    catagory,
		PlaceDetail: details,
	}

	_, err = svc.dbQueries.CreatePlace(ctx, params)
	if err != nil {
		slog.Warn("Fail to get place details from provided place information", "error", err)
		return errors.New("fail to get place details from provided place information")
	}
	return nil
}

// ////////////////////////////////////////////////////////////////////
func (svc *PostgresBookmarkService) GetBookmarkPlaceDetails(place_id string) (*Place, error) {

	// Create an empty context
	ctx := context.Background()

	// Check if place inside database
	placeInfo, err := svc.dbQueries.GetPlace(ctx, place_id)
	if err != nil {
		slog.Warn("Fail to get place details from provided place id", "error", err)
		return nil, errors.New("fail to get place details from provided place id")
	}

	// Unmarshal the JSONB column into the struct
	var placeDetail PlaceDetails
	var placeDetailBytes []byte
	err = json.Unmarshal(placeDetailBytes, &placeDetail)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("error unmarshal JSONB place detail")
	}

	// Return place Info
	placeInformation := &Place{
		ID:          placeInfo.ID,
		PlaceName:   placeInfo.PlaceName,
		Country:     placeInfo.Country,
		City:        placeInfo.City,
		Category:    placeInfo.Category,
		PlaceDetail: placeDetail,
	}

	return placeInformation, nil
}

func (svc *PostgresBookmarkService) CreateUserBookmark(user_id uuid.UUID, username string, place_id string, place_name string, place_text string) (*UserBookmarkPlace, error) {
	// Create an empty context
	ctx := context.Background()

	// Create user bookmark into database
	params := database.CreateUserBookmarkParams{
		UserID:    user_id,
		Username:  username,
		PlaceID:   place_id,
		PlaceName: place_name,
		PlaceText: place_text,
	}

	userBookmarkPlace, err := svc.dbQueries.CreateUserBookmark(ctx, params)
	if err != nil {
		slog.Warn("Fail to create user bookmark", "error", err)
		return nil, errors.New("fail to create user bookmark")
	}

	// Return user bookmark place
	bookmarkPlace := &UserBookmarkPlace{
		UserID:  userBookmarkPlace.UserID,
		PlaceID: userBookmarkPlace.PlaceID,
	}

	return bookmarkPlace, nil
}
