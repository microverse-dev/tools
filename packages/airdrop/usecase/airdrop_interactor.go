package usecase

import (
	"context"
	"flag"
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
	// get flag
	var isDryRun = flag.Bool("dryrun", true, "dry run")
	flag.Parse()

	fmt.Printf("dry run: %t\n", *isDryRun)
	if *isDryRun {
		fmt.Println("if you want to send NFT, please add `--dryrun=false` option")
	}

	// read csv
	targets, err := csv.ReadCsv("list.csv")
	if err != nil {
		return err
	}

	// get owner address
	wallet, err := ethutil.NewWallet(adi.app.Config.WalletSecretKey, *adi.app)
	if err != nil {
		return err
	}

	// create context
	ctx := context.Background()

	// call alchemy api
	alchemyClient := alchemy.NewClient(adi.app.Config.RpcUrl)
	res, err := alchemyClient.GetNFTs(ctx, alchemy.GetNFTsRequest{
		Owner:           wallet.Address.Hex(),
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

	var totalAmount int
	for _, target := range targets {
		totalAmount += target.Amount
	}

	fmt.Println("total transaction count:", totalAmount)

	// check stock
	if len(tokenIdList) < totalAmount {
		return fmt.Errorf("stock is not enough")
	}

	var index int
	for _, target := range targets {
		for j := 0; j < target.Amount; j++ {
			tokenId := tokenIdList[index]
			fmt.Println("[", index+1, "]", "sending NFT to", target.WalletAddress, "tokenId is", tokenId, "by", wallet.Address.Hex())

			if !*isDryRun {
				txHash, err := wallet.Transfer(adi.app.Config.ContractAddress, target.WalletAddress, tokenId)
				if err != nil {
					fmt.Println("error", err.Error())
				}

				if txHash != nil {
					fmt.Println("txHash", *txHash)
				}

				// トランザクションが詰まるので、1秒待つ
				time.Sleep(1 * time.Second)
			}

			index++
		}
	}

	return nil
}
