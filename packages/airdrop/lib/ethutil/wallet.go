package ethutil

import (
	"crypto/ecdsa"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/microverse-dev/tools/packages/airdrop/di"
)

type Wallet struct {
	Address    common.Address
	Client     *ethclient.Client
	privateKey *ecdsa.PrivateKey
}

func NewWallet(privateHexKey string, app di.App) (*Wallet, error) {
	privateKey, err := crypto.HexToECDSA(privateHexKey)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error casting public key to ECDSA")
	}

	client, err := ethclient.Dial(app.Config.RpcUrl)
	if err != nil {
		return nil, err
	}

	return &Wallet{
		Address:    crypto.PubkeyToAddress(*publicKeyECDSA),
		Client:     client,
		privateKey: privateKey,
	}, nil
}
