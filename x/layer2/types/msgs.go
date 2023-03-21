package types

import (
	"github.com/TsukiCore/tsuki/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgCreateDappProposal{}

func (m *MsgCreateDappProposal) Route() string {
	return ModuleName
}
func (m *MsgCreateDappProposal) Type() string {
	return types.MsgTypeCreateDappProposal
}
func (m *MsgCreateDappProposal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return err
	}

	return nil
}
func (m *MsgCreateDappProposal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}
func (m *MsgCreateDappProposal) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

var _ sdk.Msg = &MsgBondDappProposal{}

func (m *MsgBondDappProposal) Route() string {
	return ModuleName
}
func (m *MsgBondDappProposal) Type() string {
	return types.MsgTypeBondDappProposal
}
func (m *MsgBondDappProposal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return err
	}

	return nil
}
func (m *MsgBondDappProposal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}
func (m *MsgBondDappProposal) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

var _ sdk.Msg = &MsgReclaimDappBondProposal{}

func (m *MsgReclaimDappBondProposal) Route() string {
	return ModuleName
}
func (m *MsgReclaimDappBondProposal) Type() string {
	return types.MsgTypeReclaimDappBondProposal
}
func (m *MsgReclaimDappBondProposal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return err
	}

	return nil
}
func (m *MsgReclaimDappBondProposal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}
func (m *MsgReclaimDappBondProposal) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

var _ sdk.Msg = &MsgExitDapp{}

func (m *MsgExitDapp) Route() string {
	return ModuleName
}
func (m *MsgExitDapp) Type() string {
	return types.MsgTypeExitDapp
}
func (m *MsgExitDapp) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return err
	}

	return nil
}
func (m *MsgExitDapp) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}
func (m *MsgExitDapp) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

var _ sdk.Msg = &MsgRedeemDappPoolTx{}

func (m *MsgRedeemDappPoolTx) Route() string {
	return ModuleName
}
func (m *MsgRedeemDappPoolTx) Type() string {
	return types.MsgTypeRedeemDappPoolTx
}
func (m *MsgRedeemDappPoolTx) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return err
	}

	return nil
}
func (m *MsgRedeemDappPoolTx) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}
func (m *MsgRedeemDappPoolTx) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

var _ sdk.Msg = &MsgSwapDappPoolTx{}

func (m *MsgSwapDappPoolTx) Route() string {
	return ModuleName
}
func (m *MsgSwapDappPoolTx) Type() string {
	return types.MsgTypeSwapDappPoolTx
}
func (m *MsgSwapDappPoolTx) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return err
	}

	return nil
}
func (m *MsgSwapDappPoolTx) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}
func (m *MsgSwapDappPoolTx) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

var _ sdk.Msg = &MsgConvertDappPoolTx{}

func (m *MsgConvertDappPoolTx) Route() string {
	return ModuleName
}
func (m *MsgConvertDappPoolTx) Type() string {
	return types.MsgTypeConvertDappPoolTx
}
func (m *MsgConvertDappPoolTx) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return err
	}

	return nil
}
func (m *MsgConvertDappPoolTx) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}
func (m *MsgConvertDappPoolTx) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

var _ sdk.Msg = &MsgPauseDappTx{}

func (m *MsgPauseDappTx) Route() string {
	return ModuleName
}
func (m *MsgPauseDappTx) Type() string {
	return types.MsgTypePauseDappTx
}
func (m *MsgPauseDappTx) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return err
	}

	return nil
}
func (m *MsgPauseDappTx) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}
func (m *MsgPauseDappTx) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

var _ sdk.Msg = &MsgUnPauseDappTx{}

func (m *MsgUnPauseDappTx) Route() string {
	return ModuleName
}
func (m *MsgUnPauseDappTx) Type() string {
	return types.MsgTypeUnPauseDappTx
}
func (m *MsgUnPauseDappTx) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return err
	}

	return nil
}
func (m *MsgUnPauseDappTx) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}
func (m *MsgUnPauseDappTx) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

var _ sdk.Msg = &MsgReactivateDappTx{}

func (m *MsgReactivateDappTx) Route() string {
	return ModuleName
}
func (m *MsgReactivateDappTx) Type() string {
	return types.MsgTypeReactivateDappTx
}
func (m *MsgReactivateDappTx) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return err
	}

	return nil
}
func (m *MsgReactivateDappTx) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}
func (m *MsgReactivateDappTx) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

var _ sdk.Msg = &MsgExecuteDappTx{}

func (m *MsgExecuteDappTx) Route() string {
	return ModuleName
}
func (m *MsgExecuteDappTx) Type() string {
	return types.MsgTypeExecuteDappTx
}
func (m *MsgExecuteDappTx) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return err
	}

	return nil
}
func (m *MsgExecuteDappTx) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}
func (m *MsgExecuteDappTx) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

var _ sdk.Msg = &MsgDenounceLeaderTx{}

func (m *MsgDenounceLeaderTx) Route() string {
	return ModuleName
}
func (m *MsgDenounceLeaderTx) Type() string {
	return types.MsgTypeDenounceLeaderTx
}
func (m *MsgDenounceLeaderTx) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return err
	}

	return nil
}
func (m *MsgDenounceLeaderTx) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}
func (m *MsgDenounceLeaderTx) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

var _ sdk.Msg = &MsgTransferDappTx{}

func (m *MsgTransferDappTx) Route() string {
	return ModuleName
}
func (m *MsgTransferDappTx) Type() string {
	return types.MsgTypeTransferDappTx
}
func (m *MsgTransferDappTx) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return err
	}

	return nil
}
func (m *MsgTransferDappTx) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}
func (m *MsgTransferDappTx) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

var _ sdk.Msg = &MsgMintCreateFtTx{}

func (m *MsgMintCreateFtTx) Route() string {
	return ModuleName
}
func (m *MsgMintCreateFtTx) Type() string {
	return types.MsgTypeMintCreateFtTx
}
func (m *MsgMintCreateFtTx) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return err
	}

	return nil
}
func (m *MsgMintCreateFtTx) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}
func (m *MsgMintCreateFtTx) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

var _ sdk.Msg = &MsgMintCreateNftTx{}

func (m *MsgMintCreateNftTx) Route() string {
	return ModuleName
}
func (m *MsgMintCreateNftTx) Type() string {
	return types.MsgTypeMintCreateNftTx
}
func (m *MsgMintCreateNftTx) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return err
	}

	return nil
}
func (m *MsgMintCreateNftTx) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}
func (m *MsgMintCreateNftTx) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

var _ sdk.Msg = &MsgMintIssueTx{}

func (m *MsgMintIssueTx) Route() string {
	return ModuleName
}
func (m *MsgMintIssueTx) Type() string {
	return types.MsgTypeMintIssueTx
}
func (m *MsgMintIssueTx) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return err
	}

	return nil
}
func (m *MsgMintIssueTx) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}
func (m *MsgMintIssueTx) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

var _ sdk.Msg = &MsgMintBurnTx{}

func (m *MsgMintBurnTx) Route() string {
	return ModuleName
}
func (m *MsgMintBurnTx) Type() string {
	return types.MsgTypeMintBurnTx
}
func (m *MsgMintBurnTx) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return err
	}

	return nil
}
func (m *MsgMintBurnTx) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}
func (m *MsgMintBurnTx) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}
