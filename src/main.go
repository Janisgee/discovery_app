package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type SearchResult struct {
	Image       string `json:"image"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Response struct {
	Places []SearchResult `json:"places"`
}

func main() {

	// Load environment file
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("error loading .env file :%s \n ", err)
	}

	// Connect to web server port: 8080
	m := http.NewServeMux()

	m.HandleFunc("/", handlePage)

	port := os.Getenv("PORT")

	addr := ":" + port
	srv := http.Server{
		Handler:      m,
		Addr:         addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	// This blocks forever, until the server has an unrecoverable error
	fmt.Printf("server started on %s\n", addr)
	err = srv.ListenAndServe()
	if err != nil {
		fmt.Printf("%s\n ", err)
	}

}

func handlePage(w http.ResponseWriter, r *http.Request) {

	// Openai response
	resp, err := handleSearch()
	if err != nil {
		fmt.Printf("error in searching %s \n", err)
	}

	// Print the parsed attractions
	var resultString string
	for i, attraction := range resp.Places {
		fmt.Printf("%d. Name: %s\n   Description: %s\n   Image: %s\n\n",
			i+1, attraction.Name, attraction.Description, attraction.Image)
		resultString += fmt.Sprintf("<p>%d. Name: %s<br>Description: %s<br>Image: <img src='%s' alt='%s'></p>",
			i+1, attraction.Name, attraction.Description, attraction.Image, attraction.Name)
	}
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)
	page := fmt.Sprintf(`<html>
<head></head>
<body>
	<p> Hi there, welcome to Discovery App! </p>
	<p> Generated Attractions in Perth:</p>
	<p>%s</p>
</body>
</html>
`, resultString)
	_, err = w.Write([]byte(page))
	if err != nil {
		fmt.Printf("error in writing web page: %s\n ", err)
	}
}

func handleSearch() (*Response, error) {

	prompt := "Get 3 attractions in Perth, using a field 'places' containing 'image' (a URL to an image),'name' (the attraction name),'description' (a 10-word description)."

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
