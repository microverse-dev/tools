package main

import (
	"context"
	"log"

	"github.com/microverse-dev/tools/packages/airdrop/di"
	"github.com/microverse-dev/tools/packages/airdrop/usecase"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {
	// init wire
	app, err := di.NewApp(ctx)
	if err != nil {
		return err
	}

	// airdrop
	airdrop := usecase.NewAirdropInteractor(app)
	err = airdrop.Start()
	if err != nil {
		return err
	}

	return nil
}
