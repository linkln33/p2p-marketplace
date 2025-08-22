package market

import (
    "github.com/cosmos/cosmos-sdk/types/module"
    simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
    "github.com/cosmos/cosmos-sdk/x/simulation"
    "math/rand"

    marketsimulation "market/x/market/simulation"
    "market/x/market/types"
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
    accs := make([]string, len(simState.Accounts))
    for i, acc := range simState.Accounts {
        accs[i] = acc.Address.String()
    }
    marketGenesis := types.GenesisState{
        Params:      types.DefaultParams(),
        ListingList: []types.Listing{},
        ListingCount: 0,
    }
    simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&marketGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
    operations := make([]simtypes.WeightedOperation, 0)
    const (
        opWeightMsgCreateListing          = "op_weight_msg_market"
        defaultWeightMsgCreateListing int = 100
    )

    var weightMsgCreateListing int
    simState.AppParams.GetOrGenerate(opWeightMsgCreateListing, &weightMsgCreateListing, nil,
        func(_ *rand.Rand) {
            weightMsgCreateListing = defaultWeightMsgCreateListing
        },
    )
    operations = append(operations, simulation.NewWeightedOperation(
        weightMsgCreateListing,
        marketsimulation.SimulateMsgCreateListing(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
    ))
    const (
        opWeightMsgUpdateListing          = "op_weight_msg_market"
        defaultWeightMsgUpdateListing int = 100
    )

    var weightMsgUpdateListing int
    simState.AppParams.GetOrGenerate(opWeightMsgUpdateListing, &weightMsgUpdateListing, nil,
        func(_ *rand.Rand) {
            weightMsgUpdateListing = defaultWeightMsgUpdateListing
        },
    )
    operations = append(operations, simulation.NewWeightedOperation(
        weightMsgUpdateListing,
        marketsimulation.SimulateMsgUpdateListing(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
    ))
    const (
        opWeightMsgDeleteListing          = "op_weight_msg_market"
        defaultWeightMsgDeleteListing int = 100
    )

    var weightMsgDeleteListing int
    simState.AppParams.GetOrGenerate(opWeightMsgDeleteListing, &weightMsgDeleteListing, nil,
        func(_ *rand.Rand) {
            weightMsgDeleteListing = defaultWeightMsgDeleteListing
        },
    )
    operations = append(operations, simulation.NewWeightedOperation(
        weightMsgDeleteListing,
        marketsimulation.SimulateMsgDeleteListing(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
    ))

    return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
    return []simtypes.WeightedProposalMsg{}
}
