package node

import (
	"fmt"
	"net/http"
	"os"

	"example.com/m/src/check"
	"github.com/joho/godotenv"
)

func (c *Chain) GetInfura(use_https bool) string {
	var prefix, suffix string

	godotenv.Load()
	infura_api_key := os.Getenv("INFURA_API_KEY")

	if use_https {
		prefix = "https"
	} else {
		prefix = "wss"
		suffix = "/ws"
	}
	if c.name == "ethereum" {
		if c.isMainnet {
			return fmt.Sprintf("%v://mainnet.infura.io%v/3/%v", prefix, suffix, infura_api_key)
		} else {
			return fmt.Sprintf("%v://%v.infura.io%v/v3/%v", prefix, suffix, TestnetMap["mainnet"], infura_api_key)
		}
	}
	if c.isMainnet {
		return fmt.Sprintf("%v://%v-mainnet.infura.io%v/v3/%v", prefix, c.name, suffix, infura_api_key)
	}
	return fmt.Sprintf("%v://%v-%v.infura.io/%v/v3/%v", prefix, c.name, TestnetMap[c.name], suffix, infura_api_key)
}

func ArbiscanGet(queries []string) []*http.Response {
	godotenv.Load()
	api_key := os.Getenv("ARBISCAN_API_KEY")
	return_value := make([]*http.Response, 0, len(queries)) // Create a slice with capacity len(queries), but length 0
	for _, query := range queries {
		q := fmt.Sprintf("https://api.arbiscan.io/api?%v&apikey=%v", query, api_key)
		resp, err := http.Get(q) //Post could be better
		check.Check(err)

		return_value = append(return_value, resp)
	}
	return return_value
}
