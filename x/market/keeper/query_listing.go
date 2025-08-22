package keeper

import (
	"context"
	"errors"
	
	"cosmossdk.io/collections"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"market/x/market/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) ListListing(ctx context.Context, req *types.QueryAllListingRequest) (*types.QueryAllListingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	listings, pageRes, err := query.CollectionPaginate(
		ctx,
		q.k.Listing,
		req.Pagination,
		func(_ uint64, value types.Listing) (types.Listing, error){
			return value, nil
		},
	)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllListingResponse{Listing: listings, Pagination: pageRes}, nil
}

func (q queryServer) GetListing(ctx context.Context, req *types.QueryGetListingRequest) (*types.QueryGetListingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	listing, err := q.k.Listing.Get(ctx, req.Id)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, sdkerrors.ErrKeyNotFound
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetListingResponse{Listing: listing}, nil
}
