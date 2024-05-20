package node

import "fmt"

type Chain struct {
	name      string
	chainId   int32
	isMainnet bool
}

var TestnetMap = map[string][]string{
	"arbitrum": {"sepolina"},
	"mainnet":  {"sepolina"},
}

const infura_api_key = "165de63556f04a0188b11c51733d1ab0" // TODO: Read from file

func (c *Chain) GetInfura(use_https bool) string {
	var prefix, suffix string
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

var Arbitrum = Chain{name: "arbitrum", chainId: 42161, isMainnet: true}
