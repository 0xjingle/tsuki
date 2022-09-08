package types

import (
	tsukitypes "github.com/TsukiCore/tsuki/types"
	"github.com/TsukiCore/tsuki/x/gov/types"
)

func NewProposalCreateBasket(basket Basket) *ProposalCreateBasket {
	return &ProposalCreateBasket{
		Basket: basket,
	}
}

func (m *ProposalCreateBasket) ProposalType() string {
	return tsukitypes.ProposalTypeResetWholeValidatorRank
}

func (m *ProposalCreateBasket) ProposalPermission() types.PermValue {
	return types.PermCreateBasketProposal
}

func (m *ProposalCreateBasket) VotePermission() types.PermValue {
	return types.PermVoteBasketProposal
}

// ValidateBasic returns basic validation
func (m *ProposalCreateBasket) ValidateBasic() error {
	return nil
}

func NewProposalEditBasket(basket Basket) *ProposalEditBasket {
	return &ProposalEditBasket{
		Basket: basket,
	}
}

func (m *ProposalEditBasket) ProposalType() string {
	return tsukitypes.ProposalTypeResetWholeValidatorRank
}

func (m *ProposalEditBasket) ProposalPermission() types.PermValue {
	return types.PermCreateBasketProposal
}

func (m *ProposalEditBasket) VotePermission() types.PermValue {
	return types.PermVoteBasketProposal
}

// ValidateBasic returns basic validation
func (m *ProposalEditBasket) ValidateBasic() error {
	return nil
}

func NewProposalBasketWithdrawSurplus(basketId uint64) *ProposalBasketWithdrawSurplus {
	return &ProposalBasketWithdrawSurplus{
		BasketId: basketId,
	}
}

func (m *ProposalBasketWithdrawSurplus) ProposalType() string {
	return tsukitypes.ProposalTypeResetWholeValidatorRank
}

func (m *ProposalBasketWithdrawSurplus) ProposalPermission() types.PermValue {
	return types.PermCreateBasketProposal
}

func (m *ProposalBasketWithdrawSurplus) VotePermission() types.PermValue {
	return types.PermVoteBasketProposal
}

// ValidateBasic returns basic validation
func (m *ProposalBasketWithdrawSurplus) ValidateBasic() error {
	return nil
}
