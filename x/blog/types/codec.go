package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
    cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
    sdk "github.com/cosmos/cosmos-sdk/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
    // this line is used by starport scaffolding # 2
cdc.RegisterConcrete(&MsgCreateClaim{}, "blog/CreateClaim", nil)
cdc.RegisterConcrete(&MsgUpdateClaim{}, "blog/UpdateClaim", nil)
cdc.RegisterConcrete(&MsgDeleteClaim{}, "blog/DeleteClaim", nil)

} 

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
    // this line is used by starport scaffolding # 3
registry.RegisterImplementations((*sdk.Msg)(nil),
	&MsgCreateClaim{},
	&MsgUpdateClaim{},
	&MsgDeleteClaim{},
)
}

var (
	amino = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
