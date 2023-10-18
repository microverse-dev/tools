package usecase

import (
	"fmt"
	"time"

	"github.com/microverse-dev/tools/packages/airdrop/di"
	"github.com/microverse-dev/tools/packages/airdrop/lib/csv"
	"github.com/microverse-dev/tools/packages/airdrop/lib/ethutil"
)

type BalanceOfInteractor struct {
	app *di.App
}

func NewBalanceOfInteractor(app *di.App) *BalanceOfInteractor {
	return &BalanceOfInteractor{
		app: app,
	}
}

type Result struct {
	Address   string
	BalanceOf int64
}

func (bodi *BalanceOfInteractor) Start() error {
	// read csv
	targets, err := csv.ReadCsv("list.csv")
	if err != nil {
		return err
	}

	// get owner address
	wallet, err := ethutil.NewWallet(bodi.app.Config.WalletSecretKey, *bodi.app)
	if err != nil {
		return err
	}

	var result []Result

	for _, target := range targets {
		fmt.Println("=====================================")
		fmt.Println("target", target)
		balanceOf, err := wallet.BalanceOf(bodi.app.Config.ContractAddress, target.WalletAddress)
		if err != nil {
			return err
		}

		v := balanceOf.Int64()

		fmt.Println("balanceOf", v)
		result = append(result, Result{
			Address:   target.WalletAddress,
			BalanceOf: v,
		})

		time.Sleep(1 * time.Second)
	}

	var value [][]string
	for _, v := range result {
		value = append(value, []string{v.Address, fmt.Sprintf("%d", v.BalanceOf)})
	}

	err = csv.WriteCsv("balance-of-result.csv", value)

	return err
}
