package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateResource{}, "crude/CreateResource", nil)
	cdc.RegisterConcrete(&MsgUpdateResource{}, "crude/UpdateResource", nil)
	cdc.RegisterConcrete(&MsgDeleteResource{}, "crude/DeleteResource", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateResource{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateResource{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeleteResource{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
