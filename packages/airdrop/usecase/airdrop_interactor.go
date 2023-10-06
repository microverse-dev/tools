package usecase

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/microverse-dev/tools/packages/airdrop/di"
	"github.com/microverse-dev/tools/packages/airdrop/gateway/alchemy"
	"github.com/microverse-dev/tools/packages/airdrop/lib/csv"
	"github.com/microverse-dev/tools/packages/airdrop/lib/ethutil"
)

type AirdropInteractor struct {
	app *di.App
}

func NewAirdropInteractor(app *di.App) *AirdropInteractor {
	return &AirdropInteractor{
		app: app,
	}
}

func (adi *AirdropInteractor) Start() error {
	// read csv
	targets, err := csv.ReadCsv()
	if err != nil {
		return err
	}

	// get owner address
	address, err := ethutil.GetWalletAddress(adi.app.Config.WalletSecretKey)
	if err != nil {
		return err
	}

	// create context
	ctx := context.Background()

	// call alchemy api
	alchemyClient := alchemy.NewClient(adi.app.Config.RpcUrl)
	res, err := alchemyClient.GetNFTs(ctx, alchemy.GetNFTsRequest{
		Owner:           *address,
		ContractAddress: adi.app.Config.ContractAddress,
	})
	if err != nil {
		return err
	}

	// parseInt tokenId
	var tokenIdList []int64

	// format owned nfts
	for _, nft := range res.OwnedNfts {
		// get tokenId
		tokenIdHex := nft.Id.TokenId
		tokenId, err := strconv.ParseInt(tokenIdHex, 0, 64)
		if err != nil {
			return err
		}

		tokenIdList = append(tokenIdList, tokenId)
	}

	// check stock
	if len(tokenIdList) < len(targets) {
		return fmt.Errorf("stock is not enough")
	}

	// show targets length and tokenIdList length
	fmt.Println("targets length is", len(targets))
	fmt.Println("tokenIdList length is", len(tokenIdList))

	// chance to cancel
	time.Sleep(3 * time.Second)

	for i, target := range targets {
		tokenId := tokenIdList[i]
		fmt.Println("[", i+1, "]", "sending NFT to", target, "tokenId is", tokenId, "by", *address)

		// トランザクションが詰まるので、1秒待つ
		time.Sleep(1 * time.Second)
	}

	return nil
}