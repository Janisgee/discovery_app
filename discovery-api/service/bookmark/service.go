package bookmark

import (
	"context"
	"discoveryweb/internal/database"
	"discoveryweb/service"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"log/slog"
)

type postgresBookmarkService struct {
	dbQueries *database.Queries
}

func NewBookmarkPlaceService(dbQueries *database.Queries) BookmarkPlaceService {
	return &postgresBookmarkService{
		dbQueries,
	}
}

func (svc *postgresBookmarkService) GetAllBookmarkedCity(user_id uuid.UUID, city string) ([]database.GetUserBookmarkCityInfoRow, error) {
	// Create an empty context
	ctx := context.Background()

	// Get all user bookmarked place ID

	params := database.GetUserBookmarkCityInfoParams{
		UserID: user_id,
		City:   city,
	}

	allCityPlaceList, err := svc.dbQueries.GetUserBookmarkCityInfo(ctx, params)
	if err != nil {
		slog.Warn("Fail to get place ID from user bookmarked database", "error", err)
		return nil, errors.New("fail to get place ID from user bookmarked database")
	}

	return allCityPlaceList, nil

}

func (svc *postgresBookmarkService) GetAllBookmarkedPlace(user_id uuid.UUID) ([]database.GetAllUserBookmarkPlaceIDRow, error) {
	// Create an empty context
	ctx := context.Background()

	// Get all user bookmarked place ID

	allPlaceIDList, err := svc.dbQueries.GetAllUserBookmarkPlaceID(ctx, user_id)
	if err != nil {
		slog.Warn("Fail to get place ID from user bookmarked database", "error", err)
		return nil, errors.New("fail to get place ID from user bookmarked database")
	}

	return allPlaceIDList, nil

}

func (svc *postgresBookmarkService) CreatePlaceData(place_id string, place_name string, country string, city string, catagory string, place_details *service.PlaceDetails) error {
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

func (svc *postgresBookmarkService) GetPlaceDatabaseDetails(place_id string) (*Place, error) {

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
	var placeDetail service.PlaceDetails
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

func (svc *postgresBookmarkService) GetPlaceIDFromDB(placeName string) (string, error) {
	// Create an empty context
	ctx := context.Background()

	place_id, err := svc.dbQueries.GetPlaceIDFromDB(ctx, placeName)
	if err != nil {
		slog.Warn("Fail to get place id from bookmark database", "error", err)
		return "", errors.New("fail to get place id from bookmark database")
	}
	return place_id, nil
}

func (svc *postgresBookmarkService) CreateUserBookmark(user_id uuid.UUID, username string, place_id string, place_name string, catagory string, place_text string) (*UserBookmarkPlace, error) {
	// Create an empty context
	ctx := context.Background()

	// Create user bookmark into database
	params := database.CreateUserBookmarkParams{
		UserID:    user_id,
		Username:  username,
		PlaceID:   place_id,
		PlaceName: place_name,
		Catagory:  catagory,
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

func (svc *postgresBookmarkService) CheckPlaceHasBookmarkedByUser(place_id string, user_id uuid.UUID) (bool, error) {
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

func (svc *postgresBookmarkService) DeleteUserBookmark(user_id uuid.UUID, place_id string) (*UserBookmarkPlace, error) {

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
