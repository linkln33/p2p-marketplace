package keeper_test

import (
	"testing"

	"market/x/market/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ListingList:  []types.Listing{{Id: 0}, {Id: 1}},
		ListingCount: 2,
		ListingList:  []types.Listing{{Id: 0}, {Id: 1}},
		ListingCount: 2,
	}
	f := initFixture(t)
	err := f.keeper.InitGenesis(f.ctx, genesisState)
	require.NoError(t, err)
	got, err := f.keeper.ExportGenesis(f.ctx)
	require.NoError(t, err)
	require.NotNil(t, got)

	require.EqualExportedValues(t, genesisState.Params, got.Params)
	require.EqualExportedValues(t, genesisState.ListingList, got.ListingList)
	require.Equal(t, genesisState.ListingCount, got.ListingCount)
	require.EqualExportedValues(t, genesisState.ListingList, got.ListingList)
	require.Equal(t, genesisState.ListingCount, got.ListingCount)

}
