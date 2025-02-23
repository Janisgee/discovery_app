package bookmark

import (
	"discoveryweb/internal/database"
	"discoveryweb/service"
	"github.com/google/uuid"
)

type BookmarkPlaceService interface {
	CreatePlaceData(place_id string, place_name string, country string, city string, catagory string, place_detail *service.PlaceDetails) error
	GetPlaceDatabaseDetails(place_id string) (*Place, error)
	CreateUserBookmark(user_id uuid.UUID, username string, place_id string, place_name string, catagory string, place_text string) (*UserBookmarkPlace, error)
	DeleteUserBookmark(user_id uuid.UUID, place_id string) (*UserBookmarkPlace, error)
	CheckPlaceHasBookmarkedByUser(place_id string, user_id uuid.UUID) (bool, error)
	GetAllBookmarkedPlace(user_id uuid.UUID) ([]database.GetAllUserBookmarkPlaceIDRow, error)
	GetAllBookmarkedCity(user_id uuid.UUID, city string) ([]database.GetUserBookmarkCityInfoRow, error)
	GetPlaceIDFromDB(placeName string) (string, error)
}
