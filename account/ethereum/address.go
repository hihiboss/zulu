package ethereum

import (
	"github.com/DE-labtory/zulu/keychain"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func DeriveAddress(key keychain.Key) string {
	pubkeyBytes := key.PublicKey
	keyHash := crypto.Keccak256(pubkeyBytes[1:])
	addressBytes := common.BytesToAddress(keyHash[12:])
	return addressBytes.String()
}
