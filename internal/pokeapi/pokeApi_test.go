package pokeapi

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetLocationAreas(t *testing.T) {
	mockResponseBody, err := os.ReadFile("testData/location-areas-response.json")
	if err != nil {
		t.Fatalf("Failed to read mock response file: %v", err)
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(mockResponseBody)
	}))
	client := NewClient()
	defer server.Close()
	t.Run("Should return first of 20 areas", func(t *testing.T) {
		locations, err := client.GetLocationAreas(server.URL)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if len(locations.Results) != 20 {
			t.Errorf("Expected 1 result, got %d", len(locations.Results))
		}
	})
}
