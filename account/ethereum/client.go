package ethereum

import (
	"context"
	"github.com/DE-labtory/zulu/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type Client interface {
	BalanceAt(address string) (*big.Int, error)
}

type GethClient struct {
	client *ethclient.Client
}

func NewGethClient(network types.Network) *GethClient {
	client, err := ethclient.Dial(Supplier[network].NodeUrl)
	if err != nil {
		panic(err)
	}

	return &GethClient{
		client: client,
	}
}

func (c *GethClient) BalanceAt(address string) (*big.Int, error) {
	account := common.HexToAddress(address)
	balance, err := c.client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return &big.Int{}, err
	}

	return balance, nil
}
