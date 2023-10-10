package ethutil

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/microverse-dev/tools/packages/airdrop/gateway/onchain/CryptoPokers"
)

func (wallet *Wallet) Transfer(contractAddress string, toAddress string, tokenId int64) (*string, error) {
	chainId := big.NewInt(1) // ethereum mainnet
	auth, err := bind.NewKeyedTransactorWithChainID(wallet.privateKey, chainId)
	if err != nil {
		return nil, err
	}

	address := common.HexToAddress(contractAddress)
	instance, err := CryptoPokers.NewCryptoPokers(address, wallet.Client)
	if err != nil {
		return nil, err
	}

	to := common.HexToAddress(toAddress)
	tx, err := instance.SafeTransferFrom(auth, wallet.Address, to, big.NewInt(tokenId))
	if err != nil {
		return nil, err
	}

	txAddress := tx.Hash().Hex()

	return &txAddress, nil
}
