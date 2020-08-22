package ethereum

import (
	"github.com/DE-labtory/zulu/types"
	"github.com/spf13/viper"
)

type Params struct {
	NodeUrl string
}

var (
	MainnetParams = Params{
		NodeUrl: "https://mainnet.infura.io",
	}
	RopstenParams = Params{
		NodeUrl: viper.GetString("endpoint.ethereum.ropsten"),
	}
)

var Supplier = map[types.Network]Params{
	types.Mainnet: MainnetParams,
	types.Ropsten: RopstenParams,
}

func init() {}
