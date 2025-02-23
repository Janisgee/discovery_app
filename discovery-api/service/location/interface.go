package location

type LocationService interface {
	GetDetails(location string, category string) ([]LocationDetails, error)
	GetPlaceDetails(location string) (*PlaceDetails, error)
}
