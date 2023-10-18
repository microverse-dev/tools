//go:build tools
// +build tools

package tools

import (
	_ "github.com/ethereum/go-ethereum/cmd/abigen"
	_ "github.com/google/wire/cmd/wire"
)

//go:generate go build -v -o=./bin/abigen github.com/ethereum/go-ethereum/cmd/abigen
//go:generate go build -v -o=./bin/wire github.com/google/wire/cmd/wire
