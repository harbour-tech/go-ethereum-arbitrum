package eth

import (
	"context"

	"https://github.com/harbour-tech/go-ethereum-arbitrum/core"
	"github.com/harbour-tech/go-ethereum-arbitrum/go-ethereum-arbitrum/core/state"
	"github.com/harbour-tech/go-ethereum-arbitrum/go-ethereum-arbitrum/core/types"
	"github.com/harbour-tech/go-ethereum-arbitrum/go-ethereum-arbitrum/core/vm"
	"github.com/harbour-tech/go-ethereum-arbitrum/go-ethereum-arbitrum/eth/tracers"
	"github.com/harbour-tech/go-ethereum-arbitrum/go-ethereum-arbitrum/ethdb"
)

func NewArbEthereum(
	blockchain *core.BlockChain,
	chainDb ethdb.Database,
) *Ethereum {
	return &Ethereum{
		blockchain: blockchain,
		chainDb:    chainDb,
	}
}

func (eth *Ethereum) StateAtTransaction(ctx context.Context, block *types.Block, txIndex int, reexec uint64) (*core.Message, vm.BlockContext, *state.StateDB, tracers.StateReleaseFunc, error) {
	return eth.stateAtTransaction(ctx, block, txIndex, reexec)
}
