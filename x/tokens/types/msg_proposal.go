package types

import (
	tsukitypes "github.com/TsukiCore/tsuki/types"
	"github.com/TsukiCore/tsuki/x/gov/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg       = &MsgProposalUpsertTokenAlias{}
	_ types.Content = &ProposalUpsertTokenAlias{}
	_ sdk.Msg       = &MsgProposalUpsertTokenRates{}
)

func NewMsgProposalUpsertTokenAlias(
	proposer sdk.AccAddress,
	symbol string,
	name string,
	icon string,
	decimals uint32,
	denoms []string,
) *MsgProposalUpsertTokenAlias {
	return &MsgProposalUpsertTokenAlias{
		Proposer: proposer,
		Symbol:   symbol,
		Name:     name,
		Icon:     icon,
		Decimals: decimals,
		Denoms:   denoms,
	}
}

func (m *MsgProposalUpsertTokenAlias) Route() string {
	return ModuleName
}

func (m *MsgProposalUpsertTokenAlias) Type() string {
	return tsukitypes.MsgTypeProposalUpsertTokenAlias
}

func (m *MsgProposalUpsertTokenAlias) ValidateBasic() error {
	return nil
}

func (m *MsgProposalUpsertTokenAlias) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgProposalUpsertTokenAlias) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Proposer}
}

func NewMsgProposalUpsertTokenRates(proposer sdk.AccAddress, denom string, rate sdk.Dec, feePayments bool) *MsgProposalUpsertTokenRates {
	return &MsgProposalUpsertTokenRates{Denom: denom, Rate: rate, FeePayments: feePayments, Proposer: proposer}
}

func (m *MsgProposalUpsertTokenRates) Route() string {
	return ModuleName
}

func (m *MsgProposalUpsertTokenRates) Type() string {
	return tsukitypes.MsgProposalUpsertTokenRatesType
}

func (m *MsgProposalUpsertTokenRates) ValidateBasic() error {
	return nil
}

func (m *MsgProposalUpsertTokenRates) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgProposalUpsertTokenRates) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Proposer}
}
