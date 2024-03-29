package ethutil

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/microverse-dev/tools/packages/airdrop/gateway/onchain/ERC721A"
)

func (wallet *Wallet) Transfer(contractAddress string, toAddress string, tokenId int64) (*string, error) {
	chainId := big.NewInt(1) // ethereum mainnet
	auth, err := bind.NewKeyedTransactorWithChainID(wallet.privateKey, chainId)
	if err != nil {
		return nil, err
	}

	address := common.HexToAddress(contractAddress)
	instance, err := ERC721A.NewERC721A(address, wallet.Client)
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

func (wallet *Wallet) BalanceOf(contractAddress string, ownerAddress string) (*big.Int, error) {
	address := common.HexToAddress(contractAddress)
	instance, err := ERC721A.NewERC721A(address, wallet.Client)
	if err != nil {
		return big.NewInt(0), err
	}

	owner := common.HexToAddress(ownerAddress)
	balanceOf, err := instance.BalanceOf(nil, owner)
	if err != nil {
		return big.NewInt(0), err
	}

	return balanceOf, nil
}
