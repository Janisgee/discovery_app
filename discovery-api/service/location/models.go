package location

type LocationDetails struct {
	Image       string `json:"image"`
	Name        string `json:"name"`
	PlaceID     string `json:"place_id"`
	Description string `json:"description"`
	HasBookmark bool   `json:"hasbookmark"`
}

type PlaceDetails struct {
	City          string `json:"city"`
	Country       string `json:"country"`
	ImageURL      string `json:"image_url"`
	Description   string `json:"description"`
	Location      string `json:"location"`
	Opening_hours string `json:"opening_hours"`
	History       string `json:"history"`
	Key_features  string `json:"key_features"`
	Conclusion    string `json:"conclusion"`
}
