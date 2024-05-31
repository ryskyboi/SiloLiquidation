package node

type Chain struct {
	name      string
	chainId   int32
	isMainnet bool
}

var TestnetMap = map[string][]string{
	"arbitrum": {"sepolina"},
	"mainnet":  {"sepolina"},
}

var Arbitrum = Chain{name: "arbitrum", chainId: 42161, isMainnet: true}
var Arbitrum_rpc = "https://arb1.arbitrum.io/rpc"
