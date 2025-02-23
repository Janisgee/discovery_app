package places

type CityResult struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Region string `json:"region"`
}

type placeFindResponse struct {
	Candidates []struct {
		PlaceID string `json:"place_id"`
	} `json:"candidates"`
	Status string `json:"status"`
}
