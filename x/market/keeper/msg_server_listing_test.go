package keeper_test

import (
	"testing"

    sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

    "market/x/market/keeper"
    "market/x/market/types"
)

func TestListingMsgServerCreate(t *testing.T) {
	f := initFixture(t)
	srv := keeper.NewMsgServerImpl(f.keeper)

	creator, err := f.addressCodec.BytesToString([]byte("signerAddr__________________"))
	require.NoError(t, err)

	for i := 0; i < 5; i++ {
		resp, err := srv.CreateListing(f.ctx, &types.MsgCreateListing{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestListingMsgServerUpdate(t *testing.T) {
	f := initFixture(t)
	srv := keeper.NewMsgServerImpl(f.keeper)

	creator, err := f.addressCodec.BytesToString([]byte("signerAddr__________________"))
	require.NoError(t, err)

	unauthorizedAddr, err := f.addressCodec.BytesToString([]byte("unauthorizedAddr___________"))
	require.NoError(t, err)

	_, err = srv.CreateListing(f.ctx, &types.MsgCreateListing{Creator: creator})
	require.NoError(t, err)

	tests := []struct {
		desc    string
		request *types.MsgUpdateListing
		err     error
	}{
		{
			desc:    "invalid address",
			request: &types.MsgUpdateListing{Creator: "invalid"},
			err:     sdkerrors.ErrInvalidAddress,
		},
		{
			desc:    "unauthorized",
			request: &types.MsgUpdateListing{Creator: unauthorizedAddr},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "key not found",
			request: &types.MsgUpdateListing{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc:    "completed",
			request: &types.MsgUpdateListing{Creator: creator},
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			_, err = srv.UpdateListing(f.ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestListingMsgServerDelete(t *testing.T) {
	f := initFixture(t)
	srv := keeper.NewMsgServerImpl(f.keeper)

	creator, err := f.addressCodec.BytesToString([]byte("signerAddr__________________"))
	require.NoError(t, err)

	unauthorizedAddr, err := f.addressCodec.BytesToString([]byte("unauthorizedAddr___________"))
	require.NoError(t, err)

	_, err = srv.CreateListing(f.ctx, &types.MsgCreateListing{Creator: creator})
	require.NoError(t, err)

	tests := []struct {
		desc    string
		request *types.MsgDeleteListing
		err     error
	}{
		{
			desc:    "invalid address",
			request: &types.MsgDeleteListing{Creator: "invalid"},
			err:     sdkerrors.ErrInvalidAddress,
		},
		{
			desc:    "unauthorized",
			request: &types.MsgDeleteListing{Creator: unauthorizedAddr},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "key not found",
			request: &types.MsgDeleteListing{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc:    "completed",
			request: &types.MsgDeleteListing{Creator: creator},
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			_, err = srv.DeleteListing(f.ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
