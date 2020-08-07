package types

import (
	"github.com/TsukiCore/cosmos-sdk/codec"
	"github.com/TsukiCore/cosmos-sdk/codec/types"
	sdk "github.com/TsukiCore/cosmos-sdk/types"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(&MsgClaimValidator{}, "tsukiHub/MsgClaimValidator", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgClaimValidator{},
	)
}

var (
	amino = codec.New()

	// ModuleCdc references the global x/staking module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/staking and
	// defined at the application level.
	ModuleCdc = codec.NewHybridCodec(amino, types.NewInterfaceRegistry())
)

func init() {
	RegisterCodec(amino)
	amino.Seal()
}
