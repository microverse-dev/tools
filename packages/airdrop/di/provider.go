package di

import (
	"context"

	"github.com/google/wire"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	RpcUrl          string `envconfig:"RPC_URL" default:"https://eth-mainnet.g.alchemy.com/v2/7q07HgmNB4AvgtnWTJF0uKgXmHIPpC-e"`
	WalletSecretKey string `envconfig:"WALLET_SECRET_KEY"`
	ContractAddress string `envconfig:"CONTRACT_ADDRESS"`
}

func loadConfig(_ context.Context) (*Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func providerConfig(ctx context.Context) (*Config, error) {
	return loadConfig(ctx)
}

var providerSet = wire.NewSet(providerConfig)
