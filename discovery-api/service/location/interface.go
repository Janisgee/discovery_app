package location

import "github.com/google/uuid"

type LocationService interface {
	GetDetails(location string, category string) ([]LocationDetails, error)
	GetPlaceDetails(location string) (*PlaceDetails, error)
	CheckCountryCityInData(country string, city string) (string, error)
	GetCountryImageData(country string) (*CountryImage, error)
	CreateCityImageData(countryID uuid.UUID, country string, city string, cityImage string) (*CityImage, error)
	CreateCountryImageData(country string, countryImage string) (*CountryImage, error)
}
