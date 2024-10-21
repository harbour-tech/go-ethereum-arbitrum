package arbitrum

import (
	"context"

	"https://github.com/harbour-tech/go-ethereum-arbitrum/arbitrum_types"
	"github.com/harbour-tech/go-ethereum-arbitrum/go-ethereum-arbitrum/core"
	"github.com/harbour-tech/go-ethereum-arbitrum/go-ethereum-arbitrum/core/types"
)

type ArbInterface interface {
	PublishTransaction(ctx context.Context, tx *types.Transaction, options *arbitrum_types.ConditionalOptions) error
	BlockChain() *core.BlockChain
	ArbNode() interface{}
}
