package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

// GetImageURL

func (svr *ApiServer) getPlaceImageURL(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Place     string `json:"place_name"`
		SearchFor string `json:"search_for"`
		Country   string `json:"country"` // Only for searching image for cities to use
	}
	// Only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	// Parse incoming request body (JSON) into login detail struct
	var placeInfo request
	err := json.NewDecoder(r.Body).Decode(&placeInfo)
	if err != nil {
		http.Error(w, "Failed to decode place data", http.StatusBadRequest)
		return
	}

	////////////////////////////////////////////////////////////

	if placeInfo.SearchFor == "country" {
		// Check if country have image stored in database
		countryImageData, err := svr.locationSvc.GetCountryImageData(placeInfo.Place)
		if err != nil {
			// [No country image inside data]
			slog.Warn("No country image data found", "place", placeInfo.Place, "error", err)
			// ----- Get country image
			pexelImageData, err := svr.imgSvc.GetImageURL(placeInfo.Place)
			if err != nil {
				slog.Warn("Failed to decode place img url", "error", err)
				http.Error(w, "Failed to decode place img url", http.StatusBadRequest)
				return
			}
			// ----- Store country image into database
			countryImage, err := svr.locationSvc.CreateCountryImageData(placeInfo.Place, pexelImageData.ImageURL)
			if err != nil {
				slog.Error("Unable to get country data from database", "error", err)
				http.Error(w, fmt.Sprintf("Unable to get country data from database: %s", err), http.StatusInternalServerError)
				return
			}

			// Successfully created country image
			slog.Info("Successfully created country image", "place", placeInfo.Place)

			// Create a response struct to send back as JSON
			response := map[string]interface{}{
				"place_name": placeInfo.Place,
				"image_url":  countryImage.CountryImage,
			}

			// Set Content-Type to JSON and send a response
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK) // 200 OK

			// Send JSON response back to client
			err = json.NewEncoder(w).Encode(response)
			if err != nil {
				// Handle error when encoding response
				http.Error(w, "Failed to send response", http.StatusInternalServerError)
				return
			}
			return
		}
		// Create a response struct to send back as JSON
		response := map[string]interface{}{
			"place_name": placeInfo.Place,
			"image_url":  countryImageData.CountryImage,
		}

		// Set Content-Type to JSON and send a response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) // 200 OK

		// Send JSON response back to client
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			// Handle error when encoding response
			http.Error(w, "Failed to send response", http.StatusInternalServerError)
			return
		}
	}

	if placeInfo.SearchFor == "city" {

		// Check if city have image stored in database
		cityImageData, err := svr.locationSvc.GetCityImageData(placeInfo.Place, placeInfo.Country)
		if err != nil {
			// [No city image inside data]
			slog.Warn("No city image data found", "place", placeInfo.Place, "error", err)
			// ----- Get city image
			pexelImageData, err := svr.imgSvc.GetImageURL(placeInfo.Place)
			if err != nil {
				slog.Warn("Failed to decode place img url", "error", err)
				http.Error(w, "Failed to decode place img url", http.StatusBadRequest)
				return
			}
			// ----- Get Country data
			countryData, err := svr.locationSvc.GetCountryImageData(placeInfo.Country)
			if err != nil {
				slog.Warn("Failed to get country image data from database", "error", err)
				// ----- Store country image into database
				countryImage, err := svr.locationSvc.CreateCountryImageData(placeInfo.Place, pexelImageData.ImageURL)
				if err != nil {
					slog.Error("Unable to get country data from database", "error", err)
					http.Error(w, fmt.Sprintf("Unable to get country data from database: %s", err), http.StatusInternalServerError)
					return
				}
				// ----- Store city image into database
				cityImage, err := svr.locationSvc.CreateCityImageData(countryImage.CountryID, countryImage.Country, placeInfo.Place, pexelImageData.ImageURL)
				if err != nil {
					slog.Error("Unable to create city image into database", "error", err)
					http.Error(w, fmt.Sprintf("Unable to create city image into database: %s", err), http.StatusInternalServerError)
					return
				}
				// Successfully created country image
				slog.Info("Successfully created city image", "place", placeInfo.Place)

				// Create a response struct to send back as JSON
				response := map[string]interface{}{
					"place_name": placeInfo.Place,
					"image_url":  cityImage.CityImage,
				}

				// Set Content-Type to JSON and send a response
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK) // 200 OK

				// Send JSON response back to client
				err = json.NewEncoder(w).Encode(response)
				if err != nil {
					// Handle error when encoding response
					http.Error(w, "Failed to send response", http.StatusInternalServerError)
					return
				}
				return
			}
			// ----- Store city image into database
			cityImage, err := svr.locationSvc.CreateCityImageData(countryData.CountryID, countryData.Country, placeInfo.Place, pexelImageData.ImageURL)
			if err != nil {
				slog.Error("Unable to create city image into database", "error", err)
				http.Error(w, fmt.Sprintf("Unable to create city image into database: %s", err), http.StatusInternalServerError)
				return
			}

			// Successfully created country image
			slog.Info("Successfully created city image", "place", placeInfo.Place)

			// Create a response struct to send back as JSON
			response := map[string]interface{}{
				"place_name": placeInfo.Place,
				"image_url":  cityImage.CityImage,
			}

			// Set Content-Type to JSON and send a response
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK) // 200 OK

			// Send JSON response back to client
			err = json.NewEncoder(w).Encode(response)
			if err != nil {
				// Handle error when encoding response
				http.Error(w, "Failed to send response", http.StatusInternalServerError)
				return
			}
			return
		}

		// Create a response struct to send back as JSON
		response := map[string]interface{}{
			"place_name": placeInfo.Place,
			"image_url":  cityImageData.CityImage,
		}

		// Set Content-Type to JSON and send a response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) // 200 OK

		// Send JSON response back to client
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			// Handle error when encoding response
			http.Error(w, "Failed to send response", http.StatusInternalServerError)
			return
		}
	}

}
