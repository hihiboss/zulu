package ethereum

import (
	"testing"

	"github.com/DE-labtory/zulu/conf"

	"github.com/DE-labtory/zulu/keychain"
	"github.com/DE-labtory/zulu/types"
	"github.com/stretchr/testify/assert"
)

func TestDeriver_DeriveAccount_eth(t *testing.T) {
	// given
	_ = conf.Configuration{}

	eth := types.Coin{
		Id: "1",
		Blockchain: types.Blockchain{
			Platform: types.Ethereum,
			Network:  types.Ropsten,
		},
		Symbol:   "ETH",
		Decimals: 18,
		Meta:     nil,
	}
	deriver := NewDeriver(eth)
	keyGenerator := keychain.NewKeyGenerator()
	key, _ := keyGenerator.Generate()

	// when
	account, err := deriver.DeriveAccount(key)

	// then
	assert.NoError(t, err)
	assert.NotNil(t, account)
}
