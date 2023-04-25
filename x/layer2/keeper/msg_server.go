package keeper

import (
	"context"

	govtypes "github.com/TsukiCore/tsuki/x/gov/types"
	"github.com/TsukiCore/tsuki/x/layer2/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	keeper Keeper
}

// NewMsgServerImpl returns an implementation of the bank MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{
		keeper: keeper,
	}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) CreateDappProposal(goCtx context.Context, msg *types.MsgCreateDappProposal) (*types.MsgCreateDappProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	properties := k.keeper.gk.GetNetworkProperties(ctx)
	addr := sdk.MustAccAddressFromBech32(msg.Sender)

	// permission check PermCreateDappProposalWithoutBond
	isAllowed := k.keeper.CheckIfAllowedPermission(ctx, addr, govtypes.PermCreateDappProposalWithoutBond)
	if !isAllowed {
		minDappBond := properties.MinDappBond
		if msg.Bond.Denom != k.keeper.BondDenom(ctx) {
			return nil, types.ErrInvalidDappBondDenom
		}
		// check 1% of properties.MinDappBond
		if msg.Bond.Amount.Mul(sdk.NewInt(100)).LT(sdk.NewInt(int64(minDappBond)).Mul(sdk.NewInt(1000_000))) {
			return nil, types.ErrLowAmountToCreateDappProposal
		}
	}

	// send initial bond to module account
	if msg.Bond.IsPositive() {
		err := k.keeper.bk.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, sdk.Coins{msg.Bond})
		if err != nil {
			return nil, err
		}
	}

	dapp := k.keeper.GetDapp(ctx, msg.Dapp.Name)
	if dapp.Name != "" {
		return nil, types.ErrDappAlreadyExists
	}

	// create dapp object
	msg.Dapp.TotalBond = msg.Bond
	msg.Dapp.CreationTime = uint64(ctx.BlockTime().Unix())
	msg.Dapp.Status = types.Bootstrap
	k.keeper.SetDapp(ctx, msg.Dapp)
	k.keeper.SetUserDappBond(ctx, types.UserDappBond{
		DappName: msg.Dapp.Name,
		User:     msg.Sender,
		Bond:     msg.Bond,
	})

	return &types.MsgCreateDappProposalResponse{}, nil
}

func (k msgServer) BondDappProposal(goCtx context.Context, msg *types.MsgBondDappProposal) (*types.MsgBondDappProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	dapp := k.keeper.GetDapp(ctx, msg.DappName)
	if dapp.Name == "" {
		return nil, types.ErrDappDoesNotExist
	}

	if k.keeper.BondDenom(ctx) != msg.Bond.Denom {
		return nil, types.ErrInvalidDappBondDenom
	}

	// send initial bond to module account
	addr := sdk.MustAccAddressFromBech32(msg.Sender)
	err := k.keeper.bk.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, sdk.Coins{msg.Bond})
	if err != nil {
		return nil, err
	}

	properties := k.keeper.gk.GetNetworkProperties(ctx)
	if dapp.TotalBond.Amount.GTE(sdk.NewInt(int64(properties.MaxDappBond)).Mul(sdk.NewInt(1000_000))) {
		return nil, types.ErrMaxDappBondReached
	}

	dapp.TotalBond = dapp.TotalBond.Add(msg.Bond)
	k.keeper.SetDapp(ctx, dapp)

	userDappBond := k.keeper.GetUserDappBond(ctx, msg.DappName, msg.Sender)
	if userDappBond.User != "" {
		userDappBond.Bond = userDappBond.Bond.Add(msg.Bond)
	} else {
		userDappBond = types.UserDappBond{
			User:     msg.Sender,
			DappName: msg.DappName,
			Bond:     msg.Bond,
		}
	}
	k.keeper.SetUserDappBond(ctx, userDappBond)

	return &types.MsgBondDappProposalResponse{}, nil
}

func (k msgServer) ReclaimDappBondProposal(goCtx context.Context, msg *types.MsgReclaimDappBondProposal) (*types.MsgReclaimDappBondProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	userDappBond := k.keeper.GetUserDappBond(ctx, msg.DappName, msg.Sender)
	if userDappBond.DappName == "" {
		return nil, types.ErrUserDappBondDoesNotExist
	}
	if userDappBond.Bond.Denom != msg.Bond.Denom {
		return nil, types.ErrInvalidDappBondDenom
	}
	if userDappBond.Bond.Amount.LT(msg.Bond.Amount) {
		return nil, types.ErrNotEnoughUserDappBond
	}

	userDappBond.Bond.Amount = userDappBond.Bond.Amount.Sub(msg.Bond.Amount)
	k.keeper.SetUserDappBond(ctx, userDappBond)

	// send tokens back to user
	addr := sdk.MustAccAddressFromBech32(msg.Sender)
	err := k.keeper.bk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, sdk.Coins{msg.Bond})
	if err != nil {
		return nil, err
	}

	return &types.MsgReclaimDappBondProposalResponse{}, nil
}

func (k msgServer) JoinDappVerifierWithBond(goCtx context.Context, msg *types.MsgJoinDappVerifierWithBond) (*types.MsgJoinDappVerifierWithBondResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	dapp := k.keeper.GetDapp(ctx, msg.DappName)
	operator := k.keeper.GetDappOperator(ctx, msg.DappName, msg.Sender)
	if operator.DappName != "" && operator.Verifier {
		return nil, types.ErrAlreadyADappVerifier
	}

	properties := k.keeper.gk.GetNetworkProperties(ctx)
	verifierBond := properties.DappVerifierBond
	totalSupply := dapp.GetLpTokenSupply()
	dappBondLpToken := dapp.LpToken()
	lpTokenAmount := totalSupply.ToDec().Mul(verifierBond).RoundInt()
	verifierBondCoins := sdk.NewCoins(sdk.NewCoin(dappBondLpToken, lpTokenAmount))
	addr := sdk.MustAccAddressFromBech32(msg.Interx)
	if verifierBondCoins.IsAllPositive() {
		err := k.keeper.bk.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, verifierBondCoins)
		if err != nil {
			return nil, err
		}
	}

	if operator.DappName == "" {
		operator = types.DappOperator{
			DappName:       msg.DappName,
			Interx:         msg.Interx,
			Operator:       msg.Sender,
			Executor:       false,
			Verifier:       true,
			Status:         types.OperatorActive,
			BondedLpAmount: lpTokenAmount,
		}
	} else {
		operator.Verifier = true
		operator.BondedLpAmount = lpTokenAmount
	}
	k.keeper.SetDappOperator(ctx, operator)
	return &types.MsgJoinDappVerifierWithBondResponse{}, nil
}

func (k msgServer) ExitDapp(goCtx context.Context, msg *types.MsgExitDapp) (*types.MsgExitDappResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	operator := k.keeper.GetDappOperator(ctx, msg.DappName, msg.Sender)
	if operator.DappName == "" {
		return nil, types.ErrNotDappOperator
	}

	if operator.Status == types.OperatorJailed {
		return nil, types.ErrOperatorJailed
	}
	if operator.Status == types.OperatorExiting {
		return nil, types.ErrOperatorAlreadyExiting
	}

	operator.Status = types.OperatorExiting
	k.keeper.SetDappOperator(ctx, operator)

	return &types.MsgExitDappResponse{}, nil
}

func (k msgServer) PauseDappTx(goCtx context.Context, msg *types.MsgPauseDappTx) (*types.MsgPauseDappTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	operator := k.keeper.GetDappOperator(ctx, msg.DappName, msg.Sender)
	if operator.DappName == "" {
		return nil, types.ErrNotDappOperator
	}
	if operator.Status != types.OperatorActive {
		return nil, types.ErrDappOperatorNotActive
	}
	operator.Status = types.OperatorPaused
	k.keeper.SetDappOperator(ctx, operator)

	// TODO: if the validator status changes to paused, inactive or jailed then his executor status for ALL dApps should
	// also change to the same paused, inactive or jailed status so that all other executors can be informed that
	// a specific node operator is not available and will miss his execution round.

	return &types.MsgPauseDappTxResponse{}, nil
}

func (k msgServer) UnPauseDappTx(goCtx context.Context, msg *types.MsgUnPauseDappTx) (*types.MsgUnPauseDappTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	operator := k.keeper.GetDappOperator(ctx, msg.DappName, msg.Sender)
	if operator.DappName == "" {
		return nil, types.ErrNotDappOperator
	}
	if operator.Status != types.OperatorPaused {
		return nil, types.ErrDappOperatorNotPaused
	}
	operator.Status = types.OperatorActive
	k.keeper.SetDappOperator(ctx, operator)

	return &types.MsgUnPauseDappTxResponse{}, nil
}

func (k msgServer) ReactivateDappTx(goCtx context.Context, msg *types.MsgReactivateDappTx) (*types.MsgReactivateDappTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	operator := k.keeper.GetDappOperator(ctx, msg.DappName, msg.Sender)
	if operator.DappName == "" {
		return nil, types.ErrNotDappOperator
	}
	if operator.Status != types.OperatorInactive {
		return nil, types.ErrDappOperatorNotInActive
	}
	operator.Status = types.OperatorActive
	k.keeper.SetDappOperator(ctx, operator)

	return &types.MsgReactivateDappTxResponse{}, nil
}

func (k msgServer) ExecuteDappTx(goCtx context.Context, msg *types.MsgExecuteDappTx) (*types.MsgExecuteDappTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	session := k.keeper.GetDappSession(ctx, msg.DappName)
	if session.DappName == "" {
		return nil, types.ErrNoDappSessionExists
	}

	if session.NextSession.Leader != msg.Sender {
		return nil, types.ErrNotDappSessionLeader
	}
	session.NextSession.Gateway = msg.Gateway
	session.NextSession.Status = types.SessionOngoing

	if session.CurrSession == nil {
		session.CurrSession = session.NextSession
		k.keeper.CreateNewSession(ctx, msg.DappName, session.CurrSession.Leader)
	}

	return &types.MsgExecuteDappTxResponse{}, nil
}

func (k msgServer) TransitionDappTx(goCtx context.Context, msg *types.MsgTransitionDappTx) (*types.MsgTransitionDappTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	session := k.keeper.GetDappSession(ctx, msg.DappName)
	if session.DappName == "" {
		return nil, types.ErrDappSessionDoesNotExist
	}
	session.NextSession.StatusHash = msg.StatusHash
	k.keeper.SetDappSession(ctx, session)

	return &types.MsgTransitionDappTxResponse{}, nil
}

func (k msgServer) DenounceLeaderTx(goCtx context.Context, msg *types.MsgDenounceLeaderTx) (*types.MsgDenounceLeaderTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	operator := k.keeper.GetDappOperator(ctx, msg.DappName, msg.Sender)
	if !operator.Verifier {
		return nil, types.ErrNotDappVerifier
	}

	session := k.keeper.GetDappSession(ctx, msg.DappName)
	if session.CurrSession == nil || session.CurrSession.StatusHash == "" {
		return nil, types.ErrVerificationNotAllowedOnEmptySession
	}
	if session.CurrSession.Leader == msg.Sender {
		return nil, types.ErrLeaderCannotEvaluateSelfSubmission
	}
	// TODO: version check on the msg and dapp
	// TODO: update it to be put on session
	k.keeper.SetDappLeaderDenouncement(ctx, types.DappLeaderDenouncement{
		DappName:     msg.DappName,
		Leader:       msg.Leader,
		Sender:       msg.Sender,
		Denouncement: msg.DenounceText,
	})

	return &types.MsgDenounceLeaderTxResponse{}, nil
}

func (k msgServer) ApproveDappTransitionTx(goCtx context.Context, msg *types.MsgApproveDappTransitionTx) (*types.MsgApproveDappTransitionTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	operator := k.keeper.GetDappOperator(ctx, msg.DappName, msg.Sender)
	if !operator.Verifier {
		return nil, types.ErrNotDappVerifier
	}

	session := k.keeper.GetDappSession(ctx, msg.DappName)
	if session.CurrSession == nil || session.CurrSession.StatusHash == "" {
		return nil, types.ErrVerificationNotAllowedOnEmptySession
	}
	if session.CurrSession.Leader == msg.Sender {
		return nil, types.ErrLeaderCannotEvaluateSelfSubmission
	}
	// TODO: version check on the msg and dapp
	k.keeper.SetDappSessionApproval(ctx, types.DappSessionApproval{
		DappName:   msg.DappName,
		Approver:   msg.Sender,
		IsApproved: true,
	})

	// TODO: check after full implementation
	// The current session status can change to `accepted` if and only if 2/3 of executors who are NOT a leader send
	// `approve-dapp-transition-tx` and no verifiers submitted the evidence of faults requesting the application to be halted,
	// additionally the total number of approvals must be no less then `verifiers_min`.
	// It might happen that the application will only have a single executor,
	// meaning that there is always an insufficient number of verifiers to approve the transition.
	// In such cases where only one executor of the dApp exists, the approval of **NO LESS THAN** 2/3 of ALL active verifiers is required
	// for the session state to change into `accepted` (the `verifiers_min` rule also applies).

	// if more than 2/3 verify, convert to accepted
	verifiers := k.keeper.GetDappVerifiers(ctx, msg.DappName)
	approvals := k.keeper.GetDappSessionApprovals(ctx, msg.DappName)
	if len(verifiers)*2/3 <= len(approvals) {
		session.CurrSession.Status = types.SessionAccepted
		k.keeper.SetDappSession(ctx, session)
		k.keeper.CreateNewSession(ctx, msg.DappName, session.NextSession.Leader)

		isApprover := make(map[string]bool)
		for _, approval := range approvals {
			isApprover[approval.Approver] = true
		}

		// dapp operator rank management
		executor := k.keeper.GetDappOperator(ctx, session.DappName, session.CurrSession.Leader)
		k.keeper.HandleSessionParticipation(ctx, executor, true)
		for _, verifier := range verifiers {
			k.keeper.HandleSessionParticipation(ctx, verifier, isApprover[verifier.Operator])
		}
	}

	return &types.MsgApproveDappTransitionTxResponse{}, nil
}

func (k msgServer) RejectDappTransitionTx(goCtx context.Context, msg *types.MsgRejectDappTransitionTx) (*types.MsgRejectDappTransitionTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	operator := k.keeper.GetDappOperator(ctx, msg.DappName, msg.Sender)
	if !operator.Verifier {
		return nil, types.ErrNotDappVerifier
	}

	session := k.keeper.GetDappSession(ctx, msg.DappName)
	if session.CurrSession.Leader == msg.Sender {
		return nil, types.ErrLeaderCannotEvaluateSelfSubmission
	}
	// TODO: version check on the msg and dapp

	if session.CurrSession == nil || session.CurrSession.StatusHash == "" {
		return nil, types.ErrVerificationNotAllowedOnEmptySession
	}

	// halt the session
	session.CurrSession.Status = types.SessionHalted
	k.keeper.SetDappSession(ctx, session)

	k.keeper.SetDappSessionApproval(ctx, types.DappSessionApproval{
		DappName:   msg.DappName,
		Approver:   msg.Sender,
		IsApproved: true,
	})

	return &types.MsgRejectDappTransitionTxResponse{}, nil
}

func (k msgServer) TransferDappTx(goCtx context.Context, msg *types.MsgTransferDappTx) (*types.MsgTransferDappTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	helper := k.keeper.GetBridgeRegistrarHelper(ctx)
	nextXid := helper.NextXam
	for _, xam := range msg.Requests {
		// TODO: check accuracy of source in case of user
		// TODO: transfer of token in case of user
		// TODO: check accuracy of source in case of dapp
		// TODO: transfer of token in case of dapp
		k.keeper.SetXAM(ctx, types.XAM{
			Xid: nextXid,
			Req: xam,
			Irc: 0,
			Src: 0,
			Drc: 0,
			Irm: 0,
			Srm: 0,
			Drm: 0,
		})
		nextXid += 1
	}
	helper.NextXam = nextXid
	k.keeper.SetBridgeRegistrarHelper(ctx, helper)

	return &types.MsgTransferDappTxResponse{}, nil
}

func (k msgServer) RedeemDappPoolTx(goCtx context.Context, msg *types.MsgRedeemDappPoolTx) (*types.MsgRedeemDappPoolTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	addr := sdk.MustAccAddressFromBech32(msg.Sender)
	dapp := k.keeper.GetDapp(ctx, msg.DappName)
	if dapp.Name != "" {
		return nil, types.ErrDappDoesNotExist
	}
	lpTokenPrice := k.keeper.LpTokenPrice(ctx, dapp)
	withoutSlippage := msg.LpToken.Amount.ToDec().Mul(lpTokenPrice)

	receiveCoin, err := k.keeper.RedeemDappPoolTx(ctx, addr, dapp, dapp.PoolFee, msg.LpToken)
	if err != nil {
		return nil, err
	}

	properties := k.keeper.gk.GetNetworkProperties(ctx)
	maxSlippage := msg.Slippage
	if maxSlippage.IsZero() {
		maxSlippage = properties.DappPoolSlippageDefault
	}
	slippage := sdk.OneDec().Sub(receiveCoin.Amount.ToDec().Quo(withoutSlippage))
	if slippage.GT(maxSlippage) {
		return nil, types.ErrOperationExceedsSlippage
	}

	return &types.MsgRedeemDappPoolTxResponse{}, nil
}

func (k msgServer) SwapDappPoolTx(goCtx context.Context, msg *types.MsgSwapDappPoolTx) (*types.MsgSwapDappPoolTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	addr := sdk.MustAccAddressFromBech32(msg.Sender)
	dapp := k.keeper.GetDapp(ctx, msg.DappName)
	if dapp.Name != "" {
		return nil, types.ErrDappDoesNotExist
	}

	lpTokenPrice := k.keeper.LpTokenPrice(ctx, dapp)
	if lpTokenPrice.IsZero() {
		return nil, types.ErrOperationExceedsSlippage
	}
	withoutSlippage := msg.Token.Amount.ToDec().Quo(lpTokenPrice)

	receiveCoin, err := k.keeper.SwapDappPoolTx(ctx, addr, dapp, dapp.PoolFee, msg.Token)
	if err != nil {
		return nil, err
	}

	properties := k.keeper.gk.GetNetworkProperties(ctx)
	maxSlippage := msg.Slippage
	if maxSlippage.IsZero() {
		maxSlippage = properties.DappPoolSlippageDefault
	}
	slippage := sdk.OneDec().Sub(receiveCoin.Amount.ToDec().Quo(withoutSlippage))
	if slippage.GT(maxSlippage) {
		return nil, types.ErrOperationExceedsSlippage
	}

	return &types.MsgSwapDappPoolTxResponse{}, nil
}

func (k msgServer) ConvertDappPoolTx(goCtx context.Context, msg *types.MsgConvertDappPoolTx) (*types.MsgConvertDappPoolTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	addr := sdk.MustAccAddressFromBech32(msg.Sender)
	dapp1 := k.keeper.GetDapp(ctx, msg.DappName)
	if dapp1.Name != "" {
		return nil, types.ErrDappDoesNotExist
	}

	lpTokenPrice1 := k.keeper.LpTokenPrice(ctx, dapp1)
	if lpTokenPrice1.IsZero() {
		return nil, types.ErrOperationExceedsSlippage
	}
	dapp2 := k.keeper.GetDapp(ctx, msg.TargetDappName)
	if dapp2.Name != "" {
		return nil, types.ErrDappDoesNotExist
	}

	lpTokenPrice2 := k.keeper.LpTokenPrice(ctx, dapp2)
	if lpTokenPrice2.IsZero() {
		return nil, types.ErrOperationExceedsSlippage
	}

	withoutSlippage := msg.LpToken.Amount.ToDec().Mul(lpTokenPrice1).Quo(lpTokenPrice2)

	receiveCoin, err := k.keeper.ConvertDappPoolTx(ctx, addr, dapp1, dapp2, msg.LpToken)
	if err != nil {
		return nil, err
	}

	properties := k.keeper.gk.GetNetworkProperties(ctx)
	maxSlippage := msg.Slippage
	if maxSlippage.IsZero() {
		maxSlippage = properties.DappPoolSlippageDefault
	}
	slippage := sdk.OneDec().Sub(receiveCoin.Amount.ToDec().Quo(withoutSlippage))
	if slippage.GT(maxSlippage) {
		return nil, types.ErrOperationExceedsSlippage
	}

	return &types.MsgConvertDappPoolTxResponse{}, nil
}

func (k msgServer) MintCreateFtTx(goCtx context.Context, msg *types.MsgMintCreateFtTx) (*types.MsgMintCreateFtTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	properties := k.keeper.gk.GetNetworkProperties(ctx)
	fee := sdk.NewInt64Coin(k.keeper.BondDenom(ctx), int64(properties.MintingFtFee))
	sender := sdk.MustAccAddressFromBech32(msg.Sender)
	err := k.keeper.bk.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.Coins{fee})
	if err != nil {
		return nil, err
	}

	err = k.keeper.bk.BurnCoins(ctx, types.ModuleName, sdk.Coins{fee})
	if err != nil {
		return nil, err
	}

	denom := "ku/" + msg.DenomSuffix

	info := k.keeper.GetTokenInfo(ctx, denom)
	if info.Denom != "" {
		return nil, types.ErrTokenAlreadyRegistered
	}

	k.keeper.SetTokenInfo(ctx, types.TokenInfo{
		TokenType:   "adr20",
		Denom:       denom,
		Name:        msg.Name,
		Symbol:      msg.Symbol,
		Icon:        msg.Icon,
		Description: msg.Description,
		Website:     msg.Website,
		Social:      msg.Social,
		Decimals:    msg.Decimals,
		Cap:         msg.Cap,
		Supply:      msg.Supply,
		Holders:     msg.Holders,
		Fee:         msg.Fee,
		Owner:       msg.Owner,
		Metadata:    "",
		Hash:        "",
	})

	return &types.MsgMintCreateFtTxResponse{}, nil
}

func (k msgServer) MintCreateNftTx(goCtx context.Context, msg *types.MsgMintCreateNftTx) (*types.MsgMintCreateNftTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	properties := k.keeper.gk.GetNetworkProperties(ctx)
	fee := sdk.NewInt64Coin(k.keeper.BondDenom(ctx), int64(properties.MintingFtFee))
	sender := sdk.MustAccAddressFromBech32(msg.Sender)
	err := k.keeper.bk.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.Coins{fee})
	if err != nil {
		return nil, err
	}

	err = k.keeper.bk.BurnCoins(ctx, types.ModuleName, sdk.Coins{fee})
	if err != nil {
		return nil, err
	}

	denom := "ku/" + msg.DenomSuffix
	info := k.keeper.GetTokenInfo(ctx, denom)
	if info.Denom != "" {
		return nil, types.ErrTokenAlreadyRegistered
	}

	k.keeper.SetTokenInfo(ctx, types.TokenInfo{
		TokenType:   "adr43",
		Denom:       denom,
		Name:        msg.Name,
		Symbol:      msg.Symbol,
		Icon:        msg.Icon,
		Description: msg.Description,
		Website:     msg.Website,
		Social:      msg.Social,
		Decimals:    msg.Decimals,
		Cap:         msg.Cap,
		Supply:      msg.Supply,
		Holders:     msg.Holders,
		Fee:         msg.Fee,
		Owner:       msg.Owner,
		Metadata:    msg.Metadata,
		Hash:        msg.Hash,
	})

	return &types.MsgMintCreateNftTxResponse{}, nil
}

func (k msgServer) MintIssueTx(goCtx context.Context, msg *types.MsgMintIssueTx) (*types.MsgMintIssueTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sender := sdk.MustAccAddressFromBech32(msg.Sender)
	tokenInfo := k.keeper.GetTokenInfo(ctx, msg.Denom)
	if tokenInfo.Denom == "" {
		return nil, types.ErrTokenNotRegistered
	}

	if msg.Sender != tokenInfo.Owner {
		fee := msg.Amount.Mul(tokenInfo.Fee).Quo(Pow10(tokenInfo.Decimals))
		feeCoins := sdk.Coins{sdk.NewCoin(k.keeper.BondDenom(ctx), fee)}
		if fee.IsPositive() {
			if tokenInfo.Owner == "" {
				err := k.keeper.bk.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, feeCoins)
				if err != nil {
					return nil, err
				}
			} else {
				owner := sdk.MustAccAddressFromBech32(tokenInfo.Owner)
				err := k.keeper.bk.SendCoins(ctx, sender, owner, feeCoins)
				if err != nil {
					return nil, err
				}
			}
		} else {
			return nil, types.ErrNotAbleToMintCoinsWithoutFee
		}
	}

	mintCoin := sdk.NewCoin(msg.Denom, msg.Amount)
	err := k.keeper.bk.MintCoins(ctx, types.ModuleName, sdk.Coins{mintCoin})
	if err != nil {
		return nil, err
	}

	err = k.keeper.bk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, sdk.Coins{mintCoin})
	if err != nil {
		return nil, err
	}

	tokenInfo.Supply = tokenInfo.Supply.Add(msg.Amount)
	if tokenInfo.Supply.GT(tokenInfo.Cap) {
		return nil, types.ErrCannotExceedTokenCap
	}
	k.keeper.SetTokenInfo(ctx, tokenInfo)

	return &types.MsgMintIssueTxResponse{}, nil
}

func (k msgServer) MintBurnTx(goCtx context.Context, msg *types.MsgMintBurnTx) (*types.MsgMintBurnTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sender := sdk.MustAccAddressFromBech32(msg.Sender)
	tokenInfo := k.keeper.GetTokenInfo(ctx, msg.Denom)
	if tokenInfo.Denom == "" {
		return nil, types.ErrTokenNotRegistered
	}

	burnCoin := sdk.NewCoin(msg.Denom, msg.Amount)
	err := k.keeper.bk.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.Coins{burnCoin})
	if err != nil {
		return nil, err
	}

	err = k.keeper.bk.BurnCoins(ctx, types.ModuleName, sdk.Coins{burnCoin})
	if err != nil {
		return nil, err
	}

	tokenInfo.Supply = tokenInfo.Supply.Sub(msg.Amount)
	k.keeper.SetTokenInfo(ctx, tokenInfo)

	return &types.MsgMintBurnTxResponse{}, nil
}

// TODO:
// ### c**) Team & Investors Incentives**
// Here are a few examples of ways in which “issuance” and “pool” configuration parameters can be used:
// - Fair Launch - no extra tokens issued and all LP coins are immediately unlocked (`pool.drip` set to 0).
// - User Assisted Launch - LP Spending Pool is configured to slowly distribute LP tokens, the `issuance.premint` is set to
// a small reasonable amount while the `issuance.postmint` is not used.
// This enables small teams that need to hire a few developers to establish a token treasury
// and sell their stake to users that are locked in the LP.
// - Investor Assisted Launch - LP Spending Pool is configured to slowly distribute LP tokens
// while premint and postmint enable the creation of treasury and sale of SAFT agreements for large-scale projects.
// The `issuance.time` parameter can be used to clearly define the time when investor tokens will be issued during the “postmint” event
// while the `issuance.deposit` address can be set up by the team as a Spending Pool to easily distribute tokens to their
// rightful owners as well as configure an **optional** “drip” if needed to not scare the LP token holders with an immediate increase of
// the token supply.
