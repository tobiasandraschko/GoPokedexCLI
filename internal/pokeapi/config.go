package pokeapi

import (
	"net/http"
	"time"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

func SpawnClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}