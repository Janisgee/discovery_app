package places

type PlacesService interface {
	AutocompleteCities(search string, lang string) ([]CityResult, error)
	GetPlaceID(location string) (string, error)
}
