package pokeapi

import (
	"net/http"
	"time"

	"github.com/tobiasandraschko/pokedexcli/internal/pokecache"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
	reapInterval = 10 * time.Second
)

func SpawnClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.CreateNewCache(reapInterval),
	}
}