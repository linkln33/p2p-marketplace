package keeper

import (
    "context"
    "errors"
    "fmt"

    "cosmossdk.io/collections"
    errorsmod "cosmossdk.io/errors"
    sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

    "market/x/market/types"
)


func (k msgServer) CreateListing(ctx context.Context,  msg *types.MsgCreateListing) (*types.MsgCreateListingResponse, error) {
    if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
        return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
    }

    nextId, err := k.ListingSeq.Next(ctx)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "failed to get next id")
	}

    var listing = types.Listing{
        Id:         nextId,
        Seller:     msg.Seller,
        Title:      msg.Title,
        Description: msg.Description,
        Price:      msg.Price,
        Denom:      msg.Denom,
        Status:     msg.Status,
        Buyer:      msg.Buyer,
        CreatedAt:  msg.CreatedAt,
        ExpiresAt:  msg.ExpiresAt,
    }

    if err = k.Listing.Set(
        ctx,
        nextId,
        listing,
    ); err != nil {
        return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to set listing")
    }

	return &types.MsgCreateListingResponse{
	    Id: nextId,
	}, nil
}

func (k msgServer) UpdateListing(ctx context.Context,  msg *types.MsgUpdateListing) (*types.MsgUpdateListingResponse, error) {
    if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
        return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
    }

    var listing = types.Listing{
        Id:         msg.Id,
        Seller:     msg.Seller,
        Title:      msg.Title,
        Description: msg.Description,
        Price:      msg.Price,
        Denom:      msg.Denom,
        Status:     msg.Status,
        Buyer:      msg.Buyer,
        CreatedAt:  msg.CreatedAt,
        ExpiresAt:  msg.ExpiresAt,
    }

    // Checks that the element exists
    val, err := k.Listing.Get(ctx, msg.Id)
    if err != nil {
        if errors.Is(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
		}

        return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to get listing")
    }

    // Checks if the msg creator is the same as the current owner
    if msg.Creator != val.Seller {
        return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
    }

	if err := k.Listing.Set(ctx, msg.Id, listing); err != nil {
        return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to update listing")
    }

	return &types.MsgUpdateListingResponse{}, nil
}

func (k msgServer) DeleteListing(ctx context.Context,  msg *types.MsgDeleteListing) (*types.MsgDeleteListingResponse, error) {
    if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
        return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
    }

    // Checks that the element exists
    val, err := k.Listing.Get(ctx, msg.Id)
    if err != nil {
        if errors.Is(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
		}

        return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to get listing")
    }

    // Checks if the msg creator is the same as the current owner
    if msg.Creator != val.Seller {
        return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
    }

	if err := k.Listing.Remove(ctx, msg.Id); err != nil {
        return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to delete listing")
    }

	return &types.MsgDeleteListingResponse{}, nil
}
