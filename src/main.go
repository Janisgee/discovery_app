package main

import (
	"encoding/json"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"log"
	"log/slog"
	"net/http"
	"os"
)

type SearchResult struct {
	Image       string `json:"image"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Response struct {
	Places []SearchResult `json:"places"`
}

type SearchCountry struct {
	Country  string `json:"country"`
	Catagory string `json:"catagory"`
}

func main() {
	// Setup logger
	jsonLogs := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	slog.SetDefault(jsonLogs)

	// Load environment file
	env, err := startupGetEnv()
	if err != nil {
		fmt.Printf("error loading environment config :%s \n ", err)
		os.Exit(1)
	}

	gptClient := openai.NewClient(env.GptKey)
	server := NewApiServer(env, gptClient)

	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}

// handlePage generates the HTML response with the attractions data
func handlePage(w http.ResponseWriter, response *Response) {
	// page := ``

	// `<html>
	// <head><title>Attractions</title></head>
	// <body>
	// <h1>Attractions</h1>
	// <p>Here are 3 attractions:</p>`

	// Add each attraction to the page'
	js_data, err := json.Marshal(response.Places)
	if err != nil {
		http.Error(w, "Failed to serialize response", http.StatusInternalServerError)
		return
	}
	// for _, attraction := range response.Places {
	// 	page += fmt.Sprintf(`
	// 	<div>
	// 		<h2>%s</h2>
	// 		<p>%s</p>
	// 		<img src="%s" alt="%s" style="width:300px; height:auto;">
	// 	</div>
	// 	`, attraction.Name, attraction.Description, attraction.Image, attraction.Name)
	// }

	// End the HTML page
	// page += `</body></html>`

	// Set the content type and write the response
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(js_data)
	if err != nil {
		http.Error(w, "Failed to write HTML response", http.StatusInternalServerError)
	}
}

// handleSearch generates a prompt to fetch attractions for a given location

func handleSearch(location string, catagory string) (*Response, error) {

	prompt := fmt.Sprintf("Get 3 %s in %s, using a field 'places' containing 'image' (a URL to an image), 'name' (the attraction name), and 'description' (a 10-word description).", catagory, location)

	// Get response from openai
	json_content, err := GetOpenAIResponse(prompt)
	if err != nil {
		return nil, fmt.Errorf("error in getting openai resonse with request:%s , error:%w \n ", prompt, err)
	}
	// Parse the JSON response into a Go structure
	var parsedResponse Response

	err = json.Unmarshal([]byte(json_content), &parsedResponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON response: %v\n Response: %s\n\n ", err, json_content)
	}

	return &parsedResponse, nil
}
