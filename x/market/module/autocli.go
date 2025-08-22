package market

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	"market/x/market/types"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: types.Query_serviceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
			RpcMethod: "ListListing",
			Use: "list-listing",
			Short: "List all listing",
		},
		{
			RpcMethod: "GetListing",
			Use: "get-listing [id]",
			Short: "Gets a listing by id",
			Alias: []string{"show-listing"},
			PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
		},
		{
			RpcMethod: "ListListing",
			Use: "list-listing",
			Short: "List all listing",
		},
		{
			RpcMethod: "GetListing",
			Use: "get-listing [id]",
			Short: "Gets a listing by id",
			Alias: []string{"show-listing"},
			PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
		},
		// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service: types.Msg_serviceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:       true, // skipped because authority gated
				},
				{
			RpcMethod: "CreateListing",
			Use: "create-listing [seller] [title] [description] [price] [denom] [status] [buyer] [created-at] [expires-at]",
			Short: "Create listing",
			PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "seller"}, {ProtoField: "title"}, {ProtoField: "description"}, {ProtoField: "price"}, {ProtoField: "denom"}, {ProtoField: "status"}, {ProtoField: "buyer"}, {ProtoField: "created_at"}, {ProtoField: "expires_at"}},
		},
		{
			RpcMethod: "UpdateListing",
			Use: "update-listing [id] [seller] [title] [description] [price] [denom] [status] [buyer] [created-at] [expires-at]",
			Short: "Update listing",
			PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "seller"}, {ProtoField: "title"}, {ProtoField: "description"}, {ProtoField: "price"}, {ProtoField: "denom"}, {ProtoField: "status"}, {ProtoField: "buyer"}, {ProtoField: "created_at"}, {ProtoField: "expires_at"}},
		},
		{
			RpcMethod: "DeleteListing",
			Use: "delete-listing [id]",
			Short: "Delete listing",
			PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
		},
		{
			RpcMethod: "CreateListing",
			Use: "create-listing [seller] [title]",
			Short: "Create listing",
			PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "seller"}, {ProtoField: "title"}},
		},
		{
			RpcMethod: "UpdateListing",
			Use: "update-listing [id] [seller] [title]",
			Short: "Update listing",
			PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "seller"}, {ProtoField: "title"}},
		},
		{
			RpcMethod: "DeleteListing",
			Use: "delete-listing [id]",
			Short: "Delete listing",
			PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
		},
		// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}