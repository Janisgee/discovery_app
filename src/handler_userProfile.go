package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func (svr *ApiServer) userProfileHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow Get requests
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Get user detail
	user_id := GetCurrentUserId(r)

	// Get user email from database by userID
	email, err := svr.userSvc.GetUserProfile(*user_id)
	if err != nil {
		http.Error(w, "Failed to get user profile through user id from database.", http.StatusBadRequest)
		return
	}

	// Masking email
	// Split the email at the '@' symbol
	splitEmail := strings.Split(email, "@")
	if len(splitEmail) < 2 {
		fmt.Println("Invalid email")
		http.Error(w, "Invalid email address.", http.StatusBadRequest)
		return
	}

	// Extract the domain part (after @)
	emailHead := splitEmail[0]

	// Mask email
	if len(emailHead) < 3 {
		//if email length less than 3
		emailHead = "xxx"
	} else {
		//if email length more than 3
		emailRune := []rune(emailHead)
		for i := 1; i < len(emailRune)-1; i++ {
			emailRune[i] = 'x'
		}
		emailHead = string(emailRune)
	}

	maskedEmail := emailHead + "@" + splitEmail[1]

	// Console response struct to send back as JSON
	response := map[string]interface{}{
		"user_email": maskedEmail,
	}

	// Set Content-Type to JSON and send a response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200 OK

	// Send JSON response back to client
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to send response", http.StatusInternalServerError)
		return
	}

}

func (svr *ApiServer) userUpdatePwHandler(w http.ResponseWriter, r *http.Request) {
	type request struct {
		CurrentPw string `json:"currentPw"`
		NewPw     string `json:"newPw"`
	}
	// Only allow Post requests
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse incoming request body (JSON) into login detail struct
	var updateRequest request
	err := json.NewDecoder(r.Body).Decode(&updateRequest)
	if err != nil {
		http.Error(w, "Failed to decode user reset password data", http.StatusBadRequest)
		return
	}

	// Get user detail
	user_id := GetCurrentUserId(r)

	// Get database hashedPw and compare with provided pw
	databasePw, err := svr.userSvc.GetUserPwByID(user_id)
	if err != nil {
		http.Error(w, "Failed to get user profile through user id from database.", http.StatusBadRequest)
		return
	}
	// Valified database hashed password and user input hashed password
	err = bcrypt.CompareHashAndPassword([]byte(databasePw), []byte(updateRequest.CurrentPw))
	if err != nil {
		http.Error(w, "Fail to verify user's current password. Please try again", http.StatusBadRequest)
		return
	}

	// Update newPw
	message, err := svr.userSvc.UpdateUserPw(updateRequest.NewPw, *user_id)
	if err != nil {
		http.Error(w, "Fail to update user new password. Please try again", http.StatusBadRequest)
		return
	}

	// Console response struct to send back as JSON
	response := map[string]interface{}{
		"message": message,
	}

	// Set Content-Type to JSON and send a response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200 OK

	// Send JSON response back to client
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to send response", http.StatusInternalServerError)
		return
	}

}

func (svr *ApiServer) userProfilePicChangeHandler(w http.ResponseWriter, r *http.Request) {

	type request struct {
		PublicID  string `json:"public_id"`
		SecureURL string `json:"secure_url"`
	}
	// Only allow Post requests
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Get user detail
	user_id := GetCurrentUserId(r)

	// Parse incoming request body (JSON) into login detail struct
	var imageRequest request
	err := json.NewDecoder(r.Body).Decode(&imageRequest)
	if err != nil {
		http.Error(w, "Failed to decode user reset password data", http.StatusBadRequest)
		return
	}

	//Update user profile picture public ID and secure url
	userUpdatedInfo, err := svr.userSvc.UpdateUserProfileImage(imageRequest.PublicID, imageRequest.SecureURL, user_id)
	if err != nil {
		http.Error(w, "Fail to update user new profile image. Please try again", http.StatusBadRequest)
		return
	}

	// Console response struct to send back as JSON
	response := map[string]interface{}{
		"message":   "User new profile picture has been successfully updated.",
		"publicID":  userUpdatedInfo.ImagePublicID,
		"secureURL": userUpdatedInfo.ImageSecureURL,
	}

	// Set Content-Type to JSON and send a response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200 OK

	// Send JSON response back to client
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to send response", http.StatusInternalServerError)
		return
	}

}

func (svr *ApiServer) userProfilePicDisplayHandler(w http.ResponseWriter, r *http.Request) {

	// Only allow Get requests
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Get user detail
	user_id := GetCurrentUserId(r)

	//Update user profile picture public ID and secure url
	userPictureData, err := svr.userSvc.DisplayUserProfileImage(*user_id)
	if err != nil {
		http.Error(w, "Fail to update user new profile image. Please try again", http.StatusBadRequest)
		return
	}

	// Console response struct to send back as JSON
	response := map[string]interface{}{
		"message":   "User new profile image data has been retrieved",
		"publicID":  userPictureData.ImagePublicID,
		"secureURL": userPictureData.ImageSecureURL,
	}

	// Set Content-Type to JSON and send a response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200 OK

	// Send JSON response back to client
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to send response", http.StatusInternalServerError)
		return
	}

}
