package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"market/x/market/types"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				ListingList: []types.Listing{{Id: 0}, {Id: 1}}, ListingCount: 2,
				ListingList: []types.Listing{{Id: 0}, {Id: 1}}, ListingCount: 2}, valid: true,
		}, {
			desc: "duplicated listing",
			genState: &types.GenesisState{
				ListingList: []types.Listing{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
				ListingList: []types.Listing{{Id: 0}, {Id: 1}}, ListingCount: 2,
			}, valid: false,
		}, {
			desc: "invalid listing count",
			genState: &types.GenesisState{
				ListingList: []types.Listing{
					{
						Id: 1,
					},
				},
				ListingCount: 0,
				ListingList:  []types.Listing{{Id: 0}, {Id: 1}}, ListingCount: 2,
			}, valid: false,
		}, {
			desc: "duplicated listing",
			genState: &types.GenesisState{
				ListingList: []types.Listing{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		}, {
			desc: "invalid listing count",
			genState: &types.GenesisState{
				ListingList: []types.Listing{
					{
						Id: 1,
					},
				},
				ListingCount: 0,
			},
			valid: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
