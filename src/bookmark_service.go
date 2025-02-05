package main

import (
	"context"
	"discoveryapp/internal/database"
	"encoding/json"
	"errors"
	"log/slog"

	"github.com/google/uuid"
)

type BookmarkPlaceService interface {
	CreatePlaceData(place_id string, place_name string, country string, city string, catagory string, place_detail *PlaceDetails) error
	GetPlaceDatabaseDetails(place_id string) (*Place, error)
	CreateUserBookmark(user_id uuid.UUID, username string, place_id string, place_name string, place_text string) (*UserBookmarkPlace, error)
	DeleteUserBookmark(user_id uuid.UUID, place_id string) (*UserBookmarkPlace, error)
	CheckPlaceHasBookmarkedByUser(place_id string, user_id uuid.UUID) (bool, error)
}

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
func (svc *PostgresBookmarkService) GetPlaceDatabaseDetails(place_id string) (*Place, error) {

	// Create an empty context
	ctx := context.Background()

	// Check if place inside database
	placeInfo, err := svc.dbQueries.GetPlace(ctx, place_id)
	if err != nil {
		// no place store in db
		slog.Warn("Fail to get place details from provided place id", "error", err)
		return nil, errors.New("fail to get place details from provided place id")
	}

	// Unmarshal the JSONB column into the struct
	var placeDetail PlaceDetails
	err = json.Unmarshal(placeInfo.PlaceDetail, &placeDetail)
	if err != nil {
		slog.Error("Error unmarshal JSONB place detail", "error", err)
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

func (svc *PostgresBookmarkService) CheckPlaceHasBookmarkedByUser(place_id string, user_id uuid.UUID) (bool, error) {
	// Create an empty context
	ctx := context.Background()

	// Create user bookmark into database
	params := database.GetUserBookmarkParams{
		PlaceID: place_id,
		UserID:  user_id,
	}

	_, err := svc.dbQueries.GetUserBookmark(ctx, params)
	if err != nil {
		slog.Warn("Fail to retrieve place data from user bookmark list", "error", err)
		return false, errors.New("fail to retrieve place data from user bookmark list")
	}

	return true, nil
}

func (svc *PostgresBookmarkService) DeleteUserBookmark(user_id uuid.UUID, place_id string) (*UserBookmarkPlace, error) {

	// Create an empty context
	ctx := context.Background()

	// Delete User Bookmark
	params := database.DeleteUserBookmarkParams{
		PlaceID: place_id,
		UserID:  user_id,
	}
	unBookmarkUserRecord, err := svc.dbQueries.DeleteUserBookmark(ctx, params)
	if err != nil {
		slog.Warn("Fail to unbookmark user place", "error", err)
		return nil, errors.New("fail to unbookmark user place")
	}

	// Return user unbookmark place
	unBookmarkPlace := &UserBookmarkPlace{
		UserID:  unBookmarkUserRecord.UserID,
		PlaceID: unBookmarkUserRecord.PlaceID,
	}

	return unBookmarkPlace, nil
}
