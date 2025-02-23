package bookmark

import (
	"discoveryweb/service"
	"github.com/google/uuid"
)

type Place struct {
	ID          string               `json:"id"`
	PlaceName   string               `json:"place_name"`
	Country     string               `json:"country"`
	City        string               `json:"city"`
	Category    string               `json:"category"`
	PlaceDetail service.PlaceDetails `json:"place_detail"`
}

type UserBookmarkPlace struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Username  string    `json:"username"`
	PlaceID   string    `json:"place_id"`
	PlaceName string    `json:"place_name"`
	Catagory  string    `json:"catagory"`
	PlaceText string    `json:"place_text"`
}
