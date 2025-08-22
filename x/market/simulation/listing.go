package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"market/x/market/keeper"
	"market/x/market/types"
)

func SimulateMsgCreateListing(
	ak types.AuthKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)

		msg := &types.MsgCreateListing{
			Creator: simAccount.Address.String(),
			Seller:  simAccount.Address.String(),
			Title:   "sim listing",
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           txGen,
			Cdc:             nil,
			Msg:             msg,
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func SimulateMsgUpdateListing(
	ak types.AuthKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			simAccount = simtypes.Account{}
			listing    = types.Listing{}
			msg        = &types.MsgUpdateListing{}
			found      = false
		)

		var allListing []types.Listing
		err := k.Listing.Walk(ctx, nil, func(key uint64, value types.Listing) (stop bool, err error) {
			allListing = append(allListing, value)
			return false, nil
		})
		if err != nil {
			panic(err)
		}

		for _, obj := range allListing {
			acc, err := ak.AddressCodec().StringToBytes(obj.Seller)
			if err != nil {
				return simtypes.OperationMsg{}, nil, err
			}

			simAccount, found = simtypes.FindAccount(accs, sdk.AccAddress(acc))
			if found {
				listing = obj
				break
			}
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "listing seller not found"), nil, nil
		}
		msg.Creator = simAccount.Address.String()
		msg.Id = listing.Id

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           txGen,
			Cdc:             nil,
			Msg:             msg,
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func SimulateMsgDeleteListing(
	ak types.AuthKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			simAccount = simtypes.Account{}
			listing = types.Listing{}
			msg = &types.MsgDeleteListing{}
			found = false
		)

		var allListing []types.Listing
		err := k.Listing.Walk(ctx, nil, func(key uint64, value types.Listing) (stop bool, err error) {
			allListing = append(allListing, value)
			return false, nil
		})
		if err != nil {
			panic(err)
		}

		for _, obj := range allListing {
			acc, err := ak.AddressCodec().StringToBytes(obj.Seller)
			if err != nil {
				return simtypes.OperationMsg{}, nil, err
			}

			simAccount, found = simtypes.FindAccount(accs, sdk.AccAddress(acc))
			if found {
				listing = obj
				break
			}
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "listing seller not found"), nil, nil
		}
		msg.Creator = simAccount.Address.String()
		msg.Id = listing.Id

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           txGen,
			Cdc:             nil,
			Msg:             msg,
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}
