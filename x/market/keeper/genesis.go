package keeper

import (
	"context"

	"market/x/market/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx context.Context, genState types.GenesisState) error {
	for _, elem := range genState.ListingList {
		if err := k.Listing.Set(ctx, elem.Id, elem); err != nil {
			return err
		}
	}

	if err := k.ListingSeq.Set(ctx, genState.ListingCount); err != nil {
		return err
	}
	for _, elem := range genState.ListingList {
		if err := k.Listing.Set(ctx, elem.Id, elem); err != nil {
			return err
		}
	}

	if err := k.ListingSeq.Set(ctx, genState.ListingCount); err != nil {
		return err
	}

	return k.Params.Set(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis.
func (k Keeper) ExportGenesis(ctx context.Context) (*types.GenesisState, error) {
	var err error

	genesis := types.DefaultGenesis()
	genesis.Params, err = k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}
	err = k.Listing.Walk(ctx, nil, func(key uint64, elem types.Listing) (bool, error) {
		genesis.ListingList = append(genesis.ListingList, elem)
		return false, nil
	})
	if err != nil {
		return nil, err
	}

	genesis.ListingCount, err = k.ListingSeq.Peek(ctx)
	if err != nil {
		return nil, err
	}
	err = k.Listing.Walk(ctx, nil, func(key uint64, elem types.Listing) (bool, error) {
		genesis.ListingList = append(genesis.ListingList, elem)
		return false, nil
	})
	if err != nil {
		return nil, err
	}

	genesis.ListingCount, err = k.ListingSeq.Peek(ctx)
	if err != nil {
		return nil, err
	}

	return genesis, nil
}
