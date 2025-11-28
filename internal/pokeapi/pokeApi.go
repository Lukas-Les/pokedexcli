package pokeapi

import (
	"encoding/json"
	"net/http"
)

type LocationAreas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (l LocationAreas) AsBytes() ([]byte, error) {
	jsonBytes, err := json.Marshal(l)
	if err != nil {
		return []byte{}, err
	}
	return jsonBytes, nil
}

type Client struct {
	httpClient http.Client
}

func NewClient() Client {
	return Client{
		httpClient: http.Client{},
	}
}

func (c Client) GetLocationAreas(url string) (LocationAreas, error) {
	response, err := c.httpClient.Get(url)
	if err != nil {
		return LocationAreas{}, err
	}
	defer response.Body.Close()
	var response_json LocationAreas
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&response_json); err != nil {
		return LocationAreas{}, err
	}
	return response_json, nil
}
