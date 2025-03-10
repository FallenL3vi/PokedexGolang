package pokeapi
import (
	"net/http"
	"time"
	"github.com/FallenL3vi/PokedexGolang/internal/pokecache"
	"math/rand"
)


type Client struct {
	cache *pokecache.Cache
	httpClient http.Client
	catchedPokemons map[string]Pokemon
}

func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	rand.Seed(time.Now().UnixNano())
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		}, catchedPokemons: make(map[string]Pokemon),
	}
}