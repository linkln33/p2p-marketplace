package keeper_test

import (
	"context"
	"testing"

    
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"market/x/market/keeper"
	"market/x/market/types"
)

func createNListing(keeper keeper.Keeper, ctx context.Context, n int) []types.Listing {
	items := make([]types.Listing, n)
	for i := range items {
		iu := uint64(i)
		items[i].Id = iu
		items[i].Seller = strconv.Itoa(i)
		items[i].Title = strconv.Itoa(i)
		_ = keeper.Listing.Set(ctx, iu, items[i])
		_ = keeper.ListingSeq.Set(ctx, iu)
	}
	return items
}

func TestListingQuerySingle(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNListing(f.keeper, f.ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetListingRequest
		response *types.QueryGetListingResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetListingRequest{Id: msgs[0].Id},
			response: &types.QueryGetListingResponse{Listing: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetListingRequest{Id: msgs[1].Id},
			response: &types.QueryGetListingResponse{Listing: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetListingRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := qs.GetListing(f.ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
			    require.NoError(t, err)
				require.EqualExportedValues(t, tc.response, response)
			}
		})
	}
}

func TestListingQueryPaginated(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNListing(f.keeper, f.ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllListingRequest {
		return &types.QueryAllListingRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListListing(f.ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Listing), step)
			require.Subset(t, msgs, resp.Listing)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListListing(f.ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Listing), step)
			require.Subset(t, msgs, resp.Listing)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := qs.ListListing(f.ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.EqualExportedValues(t, msgs, resp.Listing)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := qs.ListListing(f.ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
