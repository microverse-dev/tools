//go:build wireinject

package di

import (
	"context"

	"github.com/google/wire"
)

type App struct {
	Config *Config
}

func NewApp(ctx context.Context) (*App, error) {
	wire.Build(
		providerSet,
		wire.Struct(new(App), "*"),
	)
	return nil, nil
}
