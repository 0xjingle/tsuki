# tsuki
Tsuki Hub

## Set permission environment variables

```sh
# permissions
export PermZero=0
export PermSetPermissions=1
export PermClaimValidator=2
export PermClaimCouncilor=3
export PermCreateSetPermissionsProposal=4
export PermVoteSetPermissionProposal=5
export PermUpsertTokenAlias=6
export PermChangeTxFee=7
export PermUpsertTokenRate=8

# transaction_type
export TypeMsgSend      = "send"
export TypeMsgMultiSend = "multisend"
export TypeMsgVerifyInvariant = "verify_invariant"
export TypeMsgSetWithdrawAddress          = "set_withdraw_address"
export TypeMsgWithdrawDelegatorReward     = "withdraw_delegator_reward"
export TypeMsgWithdrawValidatorCommission = "withdraw_validator_commission"
export TypeMsgFundCommunityPool           = "fund_community_pool"
export TypeMsgSubmitEvidence = "submit_evidence"
export TypeMsgDeposit        = "deposit"
export TypeMsgVote           = "vote"
export TypeMsgSubmitProposal = "submit_proposal"
export TypeMsgTransfer = "transfer"
export TypeMsgUnjail = "unjail"
export TypeMsgUndelegate      = "begin_unbonding"
export TypeMsgEditValidator   = "edit_validator"
export TypeMsgCreateValidator = "create_validator"
export TypeMsgDelegate        = "delegate"
export TypeMsgBeginRedelegate = "begin_redelegate"
export MsgTypeProposalSetNetworkProperty = "proposal-set-network-property"
export MsgTypeProposalAssignPermission   = "proposal-assign-permission"
export MsgTypeProposalUpsertDataRegistry = "proposal-upsert-data-registry"
export MsgTypeVoteProposal               = "vote-proposal"
export MsgTypeWhitelistPermissions = "whitelist-permissions"
export MsgTypeBlacklistPermissions = "blacklist-permissions"
export MsgTypeClaimCouncilor       = "claim-councilor"
export MsgTypeSetNetworkProperties = "set-network-properties"
export MsgTypeSetExecutionFee      = "set-execution-fee"
export MsgTypeCreateRole = "create-role"
export MsgTypeAssignRole = "assign-role"
export MsgTypeRemoveRole = "remove-role"
export MsgTypeWhitelistRolePermission       = "whitelist-role-permission"
export MsgTypeBlacklistRolePermission       = "blacklist-role-permission"
export MsgTypeRemoveWhitelistRolePermission = "remove-whitelist-role-permission"
export MsgTypeRemoveBlacklistRolePermission = "remove-blacklist-role-permission"
export MsgTypeClaimValidator = "claim-validator"
export MsgTypeUpsertTokenAlias = "upsert-token-alias"
export MsgTypeUpsertTokenRate  = "upsert-token-rate"

export FuncIDMsgSend   = 1
export FuncIDMultiSend = 2
export FuncIDMsgVerifyInvariant = 3
export FuncIDMsgSetWithdrawAddress          = 4
export FuncIDMsgWithdrawDelegatorReward     = 5
export FuncIDMsgWithdrawValidatorCommission = 6
export FuncIDMsgFundCommunityPool           = 7
export FuncIDMsgSubmitEvidence = 8
export FuncIDMsgDeposit        = 9
export FuncIDMsgVote           = 10
export FuncIDMsgSubmitProposal = 11
export FuncIDMsgTransfer = 12
export FuncIDMsgUnjail = 13
export FuncIDMsgUndelegate      = 14
export FuncIDMsgEditValidator   = 15
export FuncIDMsgCreateValidator = 16
export FuncIDMsgDelegate        = 17
export FuncIDMsgBeginRedelegate = 18
export FuncIDMsgProposalSetNetworkProperty = 19
export FuncIDMsgProposalAssignPermission   = 20
export FuncIDMsgProposalUpsertDataRegistry = 21
export FuncIDMsgVoteProposal               = 22
export FuncIDMsgWhitelistPermissions = 23
export FuncIDMsgBlacklistPermissions = 24
export FuncIDMsgClaimCouncilor       = 25
export FuncIDMsgSetNetworkProperties = 26
export FuncIDMsgSetExecutionFee      = 27
export FuncIDMsgCreateRole = 28
export FuncIDMsgAssignRole = 29
export FuncIDMsgRemoveRole = 30
export FuncIDMsgWhitelistRolePermission       = 31
export FuncIDMsgBlacklistRolePermission       = 32
export FuncIDMsgRemoveWhitelistRolePermission = 33
export FuncIDMsgRemoveBlacklistRolePermission = 34
export FuncIDMsgClaimValidator = 35
export FuncIDMsgUpsertTokenAlias = 36
export FuncIDMsgUpsertTokenRate  = 37
```

## Set ChangeTxFee permission
```sh
# command to set PermChangeTxFee permission
tsukid tx customgov permission whitelist-permission --from validator --keyring-backend=test --permission=$PermChangeTxFee --addr=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --chain-id=testing --fees=100ukex --home=$HOME/.tsukid
# good response
"[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"whitelist-permissions\"}]}]}]"
```

## Query permission of an address
```sh
# command
tsukid query customgov permissions $(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid)

# response
blacklist: []
whitelist:
- 4
- 3
```
## Set network properties
```sh

# command with fee set
tsukid tx customgov set-network-properties --from validator --min_tx_fee="2" --max_tx_fee="20000" --keyring-backend=test --chain-id=testing --fees=100ukex --home=$HOME/.tsukid

# no error response
"[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"set-network-properties\"}]}]}]"

# response when not enough permissions to change tx fee
"failed to execute message; message index: 0: PermChangeTxFee: not enough permissions"

# command without fee set
tsukid tx customgov set-network-properties --from validator --min_tx_fee="2" --max_tx_fee="20000" --keyring-backend=test --chain-id=testing --home=$HOME/.tsukid

# response
"fee out of range [1, 10000]: invalid request"

```
## Query network properties
```sh
# command
tsukid query customgov network-properties

# response
properties:
  max_tx_fee: "10000"
  min_tx_fee: "1"
```

## Set Execution Fee
```sh
# command
tsukid tx customgov set-execution-fee --from validator --execution_name="B" --transaction_type="B" --execution_fee=10 --failure_fee=1 --timeout=10 default_parameters=0 --keyring-backend=test --chain-id=testing --fees=10ukex --home=$HOME/.tsukid

# response
"[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"set-execution-fee\"}]}]}]"
```

## Set execution fee validation test
```sh
# command for setting execution fee
tsukid tx customgov set-execution-fee --from validator --execution_name="set-network-properties" --transaction_type="set-network-properties" --execution_fee=10000 --failure_fee=1000 --timeout=10 default_parameters=0 --keyring-backend=test --chain-id=testing --fees=100ukex --home=$HOME/.tsukid

Here, the value should be looked at is `--execution_name="set-network-properties"`, `--execution_fee=10000` and `--failure_fee=1000`.

# check execution fee validation
tsukid tx customgov set-network-properties --from validator --min_tx_fee="2" --max_tx_fee="20000" --keyring-backend=test --chain-id=testing --fees=100ukex --home=$HOME/.tsukid
# response
"fee is less than failure fee 1000: invalid request"

Here, the value should be looked at is `"fee is less than failure fee 1000: invalid request"`.
In this case, issue is found on ante step and fee is not being paid at all.

# preparation for networks (v1) failure=1000, execution=10000
tsukid tx customgov permission whitelist-permission --from validator --keyring-backend=test --permission=$PermChangeTxFee --addr=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --chain-id=testing --fees=100ukex --home=$HOME/.tsukid --yes
tsukid tx customgov set-execution-fee --from validator --execution_name="set-network-properties" --transaction_type="set-network-properties" --execution_fee=10000 --failure_fee=1000 --timeout=10 default_parameters=0 --keyring-backend=test --chain-id=testing --fees=100ukex --home=$HOME/.tsukid --yes

# preparation for networks (v2) failure=1000, execution=500
tsukid tx customgov permission whitelist-permission --from validator --keyring-backend=test --permission=$PermChangeTxFee --addr=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --chain-id=testing --fees=100ukex --home=$HOME/.tsukid --yes
tsukid tx customgov set-execution-fee --from validator --execution_name="set-network-properties" --transaction_type="set-network-properties" --execution_fee=500 --failure_fee=1000 --timeout=10 default_parameters=0 --keyring-backend=test --chain-id=testing --fees=100ukex --home=$HOME/.tsukid --yes

# init user1 with 100000ukex
tsukid keys add user1 --keyring-backend=test --home=$HOME/.tsukid
tsukid tx bank send validator $(tsukid keys show -a user1 --keyring-backend=test --home=$HOME/.tsukid) 100000ukex --keyring-backend=test --chain-id=testing --fees=100ukex --home=$HOME/.tsukid --yes
tsukid query bank balances $(tsukid keys show -a user1 --keyring-backend=test --home=$HOME/.tsukid) --yes

# try changing set-network-properties with user1 that does not have ChangeTxFee permission
tsukid tx customgov set-network-properties --from user1 --min_tx_fee="2" --max_tx_fee="25000" --keyring-backend=test --chain-id=testing --fees=1000ukex --home=$HOME/.tsukid --yes
# this should fail and balance should be (previousBalance - failureFee)
tsukid query bank balances $(tsukid keys show -a user1 --keyring-backend=test --home=$HOME/.tsukid)

# whitelist user1's permission for ChangeTxFee and try again
tsukid tx customgov permission whitelist-permission --from validator --keyring-backend=test --permission=$PermChangeTxFee --addr=$(tsukid keys show -a user1 --keyring-backend=test --home=$HOME/.tsukid) --chain-id=testing --fees=100ukex --home=$HOME/.tsukid --yes
tsukid tx customgov set-network-properties --from user1 --min_tx_fee="2" --max_tx_fee="25000" --keyring-backend=test --chain-id=testing --fees=1000ukex --home=$HOME/.tsukid --yes
# this should fail and balance should be (previousBalance - successFee)
tsukid query bank balances $(tsukid keys show -a user1 --keyring-backend=test --home=$HOME/.tsukid)
```

## Query execution fee
```sh
tsukid query customgov execution-fee <msg_type>
# command
tsukid query customgov execution-fee "B"
# response
fee:
  default_parameters: "0"
  execution_fee: "10"
  failure_fee: "1"
  name: ABC
  timeout: "10"
  transaction_type: B

# genesis fee configuration test
tsukid query customgov execution-fee "A"
fee:
  default_parameters: "0"
  execution_fee: "10"
  failure_fee: "1"
  name: Claim Validator Seat
  timeout: "10"
  transaction_type: A
```

## Upsert token alias
```sh
# set PermUpsertTokenAlias permission to validator address
tsukid tx customgov permission whitelist-permission --from validator --keyring-backend=test --permission=$PermUpsertTokenAlias --addr=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --chain-id=testing --fees=100ukex --home=$HOME/.tsukid --yes
# run upsert alias
tsukid tx tokens upsert-alias --from validator --keyring-backend=test --expiration=0 --enactment=0 --allowed_vote_types=0,1 --symbol="ETH" --name="Ethereum" --icon="myiconurl" --decimals=6 --denoms="finney" --chain-id=testing --fees=100ukex --home=$HOME/.tsukid  --yes
```
# Query token alias
```sh
# command
tsukid query tokens alias KEX
# response
allowed_vote_types:
- "yes"
- "no"
decimals: 6
denoms:
- ukex
enactment: 0
expiration: 0
icon: myiconurl
name: Tsuki
status: undefined
symbol: KEX
```
```sh
# command
tsukid query tokens alias KE
# response
Error: KE symbol does not exist
```
```sh
# command
tsukid query tokens all-aliases --chain-id=testing --home=$HOME/.tsukid
# response
data:
- allowed_vote_types:
  - "yes"
  - "no"
  decimals: 6
  denoms:
  - ukex
  enactment: 0
  expiration: 0
  icon: myiconurl
  name: Tsuki
  status: undefined
  symbol: KEX

# command
tsukid query tokens aliases-by-denom ukex --chain-id=testing --home=$HOME/.tsukid
# response
data:
  ukex:
    allowed_vote_types:
    - "yes"
    - "no"
    decimals: 6
    denoms:
    - ukex
    enactment: 0
    expiration: 0
    icon: myiconurl
    name: Tsuki
    status: undefined
    symbol: KEX
```

## Upsert token rates
```sh
# set PermUpsertTokenRate permission to validator address
tsukid tx customgov permission whitelist-permission --from validator --keyring-backend=test --permission=$PermUpsertTokenRate --addr=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --chain-id=testing --fees=100ukex --home=$HOME/.tsukid --yes
# run upsert rate
tsukid tx tokens upsert-rate --from validator --keyring-backend=test --denom="mykex" --rate="1.5" --fee_payments=true --chain-id=testing --fees=100ukex --home=$HOME/.tsukid  --yes
```
# Query token rate
```sh
# command
tsukid query tokens rate mykex
# response
denom: mykex
fee_payments: true
rate: "1.500000"
```
```sh
# command
tsukid query tokens rate invalid_denom
# response
Error: invalid_denom denom does not exist
```
```sh
# command
tsukid query tokens all-rates --chain-id=testing --home=$HOME/.tsukid
# response
data:
- denom: ubtc
  fee_payments: true
  rate: "0.000010"
- denom: ukex
  fee_payments: true
  rate: "1.000000"
- denom: xeth
  fee_payments: true
  rate: "0.000100"

# command
tsukid query tokens rates-by-denom ukex --chain-id=testing --home=$HOME/.tsukid
# response
data:
  ukex:
    denom: ukex
    fee_payments: true
    rate: "1.000000"
```
# Fee payment in foreign currency
```sh
# register stake token as 1ukex=100stake
tsukid tx customgov permission whitelist-permission --from validator --keyring-backend=test --permission=$PermUpsertTokenRate --addr=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --chain-id=testing --fees=100ukex --home=$HOME/.tsukid --yes
tsukid tx tokens upsert-rate --from validator --keyring-backend=test --denom="stake" --rate="0.01" --fee_payments=true --chain-id=testing --fees=100ukex --home=$HOME/.tsukid  --yes
tsukid query tokens rate stake
# try to spend stake token as fee
tsukid tx tokens upsert-rate --from validator --keyring-backend=test --denom="valstake" --rate="0.01" --fee_payments=true --chain-id=testing --fees=10000stake --home=$HOME/.tsukid  --yes
# smaller amount of fee in foreign currency
tsukid tx tokens upsert-rate --from validator --keyring-backend=test --denom="valstake" --rate="0.02" --fee_payments=true --chain-id=testing --fees=1000stake --home=$HOME/.tsukid  --yes
# try to spend unregistered token (validatortoken) as fee
tsukid tx tokens upsert-rate --from validator --keyring-backend=test --denom="valstake" --rate="0.03" --fee_payments=true --chain-id=testing --fees=1000validatortoken --home=$HOME/.tsukid  --yes
```

# Fee payment in foreign currency returning failure - execution fee in foreign currency
```sh
# register stake token as 1ukex=100stake
tsukid tx customgov permission whitelist-permission --from validator --keyring-backend=test --permission=$PermUpsertTokenRate --addr=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --chain-id=testing --fees=100ukex --home=$HOME/.tsukid --yes
tsukid tx tokens upsert-rate --from validator --keyring-backend=test --denom="stake" --rate="0.01" --fee_payments=true --chain-id=testing --fees=100ukex --home=$HOME/.tsukid  --yes
tsukid query tokens rate stake

# set execution fee and failure fee for upsert-rate transaction
tsukid tx customgov permission whitelist-permission --from validator --keyring-backend=test --permission=$PermChangeTxFee --addr=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --chain-id=testing --fees=100ukex --home=$HOME/.tsukid --yes

# set execution_fee=1000 failure_fee=5000
tsukid tx customgov set-execution-fee --from validator --execution_name="upsert-token-alias" --transaction_type="upsert-token-alias" --execution_fee=1000 --failure_fee=5000 --timeout=10 default_parameters=0 --keyring-backend=test --chain-id=testing --fees=100ukex --home=$HOME/.tsukid --yes

# set execution_fee=5000 failure_fee=1000
tsukid tx customgov set-execution-fee --from validator --execution_name="upsert-token-alias" --transaction_type="upsert-token-alias" --execution_fee=5000 --failure_fee=1000 --timeout=10 default_parameters=0 --keyring-backend=test --chain-id=testing --fees=100ukex --home=$HOME/.tsukid --yes

# check current balance
tsukid query bank balances $(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid)

# try upsert-token-alias failure in foreign currency
tsukid tx tokens upsert-alias --from validator --keyring-backend=test --expiration=0 --enactment=0 --allowed_vote_types=0,1 --symbol="ETH" --name="Ethereum" --icon="myiconurl" --decimals=6 --denoms="finney" --chain-id=testing --fees=500000stake --home=$HOME/.tsukid  --yes
# set permission for this execution
tsukid tx customgov permission whitelist-permission --from validator --keyring-backend=test --permission=$PermUpsertTokenAlias --addr=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --chain-id=testing --fees=10000stake --home=$HOME/.tsukid --yes
# try upsert-token-alias success in foreign currency
tsukid tx tokens upsert-alias --from validator --keyring-backend=test --expiration=0 --enactment=0 --allowed_vote_types=0,1 --symbol="ETH" --name="Ethereum" --icon="myiconurl" --decimals=6 --denoms="finney" --chain-id=testing --fees=500000stake --home=$HOME/.tsukid  --yes
```

# Query validator account
```sh
# query validator account
tsukid query validator --addr  $(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid)
```

# Custom governance module commands

```sh
tsukid tx customgov councilor claim-seat --from validator --keyring-backend=test --home=$HOME/.tsukid

tsukid tx customgov permission blacklist-permission
tsukid tx customgov permission whitelist-permission

tsukid tx customgov proposal assign-permission
tsukid tx customgov proposal vote

tsukid tx customgov role blacklist-permission
tsukid tx customgov role create
tsukid tx customgov role remove
tsukid tx customgov role remove-blacklist-permission
tsukid tx customgov role remove-whitelist-permission
tsukid tx customgov role whitelist-permission

# querying for voters of a specific proposal
tsukid query customgov voters 1
# querying for votes of a specific proposal
tsukid query customgov votes 1
# querying for a vote of a specific propsal/voter pair
tsukid query customgov vote 1 $(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid)
```

# Commands for adding more validators

```sh
# tsukid keys add val2 --keyring-backend=test --home=$HOME/.tsukid
# tsukid tx bank send validator $(tsukid keys show -a val2 --keyring-backend=test --home=$HOME/.tsukid) 100000ukex --keyring-backend=test --chain-id=testing --fees=100ukex --home=$HOME/.tsukid --yes

tsukid tx customgov permission whitelist-permission --from validator --keyring-backend=test --permission=$PermCreateSetPermissionsProposal --addr=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --chain-id=testing --fees=100ukex --home=$HOME/.tsukid --yes
tsukid tx customgov permission whitelist-permission --from validator --keyring-backend=test --permission=$PermVoteSetPermissionProposal --addr=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --chain-id=testing --fees=100ukex --home=$HOME/.tsukid --yes

tsukid tx customgov proposal assign-permission $PermClaimValidator --addr=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes

tsukid query customgov proposals
tsukid query customgov proposal 1

tsukid tx customgov proposal vote 1 1 --from validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes 

tsukid tx claim-validator-seat --from validator --keyring-backend=test --home=$HOME/.tsukid --validator-key=tsukivaloper1ntk7n5y38en5dvnhvmruwagmkemq76x8s4pnwu --moniker="validator" --chain-id=testing --fees=100ukex --yes

# tsukid tx claim-validator-seat --from val2 --keyring-backend=test --home=$HOME/.tsukid --pubkey=tsukivalconspub1zcjduepqdllep3v5wv04hmu987rv46ax7fml65j3dh5tf237ayn5p59jyamq04048n --validator-key=tsukivaloper1ewgq8gtsefakhal687t8hnsw5zl4y8eksup39w --moniker="val2" --chain-id=testing --fees=100ukex --yes
# tsukid tx claim-validator-seat --from val2 --keyring-backend=test --home=$HOME/.tsukid --validator-key=tsukivaloper1ewgq8gtsefakhal687t8hnsw5zl4y8eksup39w --moniker="val2" --chain-id=testing --fees=100ukex --yes
```

# Tx for set network property proposal

```
tsukid tx customgov proposal set-network-property MIN_TX_FEE 101 --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes
```
---
`dev` branch
