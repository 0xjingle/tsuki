package types

import "github.com/cosmos/cosmos-sdk/codec"

var (
	amino = codec.NewLegacyAmino()

	// ModuleCdc references the global x/upgrade module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/upgrade and
	// defined at the application level.
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterCodec(amino)
	amino.Seal()
}

// RegisterCodec register codec and metadata
func RegisterCodec(cdc *codec.LegacyAmino) {}
