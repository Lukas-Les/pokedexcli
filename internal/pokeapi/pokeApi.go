package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Lukas-Les/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient() Client {
	c := pokecache.NewCache(time.Second * 5)
	return Client{
		httpClient: http.Client{},
		cache:      c,
	}
}

func (c *Client) GetLocationAreas(url string) (LocationAreas, error) {
	return handleRequest[LocationAreas](url, c)
}

func (c *Client) GetLocationArea(url string) (LocationArea, error) {
	return handleRequest[LocationArea](url, c)
}

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name)
	return handleRequest[Pokemon](url, c)
}

func fetch(client *http.Client, url string) ([]byte, error) {
	response, err := client.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer response.Body.Close()
	result, err := io.ReadAll(response.Body)
	if err != nil {
		return []byte{}, err
	}
	return result, nil
}

func deserialize[T ApiResponse](data []byte) (T, error) {
	var result T
	err := json.Unmarshal(data, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func handleRequest[T ApiResponse](url string, c *Client) (T, error) {
	var result T
	cached, exists := c.cache.Get(url)
	if exists {
		cachedResult, err := deserialize[T](cached)
		if err == nil {
			return cachedResult, nil
		}
	}
	content, err := fetch(&c.httpClient, url)
	if err != nil {
		return result, err
	}
	result, err = deserialize[T](content)
	if err != nil {
		return result, err
	}
	c.cache.Add(url, content)
	return result, nil
}
