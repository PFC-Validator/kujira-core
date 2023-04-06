package wasm

import (
	"cosmossdk.io/errors"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	icacontrollerkeeper "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/controller/keeper"
	icacontrollertypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/controller/types"
	icatypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/types"
)

// ProtobufAny is a hack-struct to serialize protobuf Any message into JSON object
// See https://github.com/neutron-org/neutron/blob/main/wasmbinding/bindings/msg.go
type ProtobufAny struct {
	TypeURL string `json:"type_url"`
	Value   []byte `json:"value"`
}

type ICAMsg struct {
	/// Contracts can register a new interchain account.
	RegisterICA *RegisterICA `json:"register,omitempty"`
	/// Contracts can submit transactions to the ICA
	/// associated with the given information.
	SubmitTxs *SubmitTx `json:"submit_txs,omitempty"`
}

// / RegisterICA creates a new interchain account.
// / If the account was created in the past, this will
// / re-establish a dropped connection, or do nothing if
// / the connection is still active.
// / The account is registered using (port, channel, sender, id)
// / as the unique identifier.
type RegisterICA struct {
	ChannelId string `json:"channel"`
	AccountId string `json:"id"`
	Version   string `json:"version"`
}

// / SubmitTx submits transactions to the ICA
// / associated with the given address.
type SubmitTx struct {
	ChannelId string                               `json:"channel"`
	AccountId string                               `json:"id"`
	Tx        icatypes.InterchainAccountPacketData `json:"tx"`
	Memo      string                               `json:"memo"`
	Timeout   uint64                               `json:"timeout"`
}

func register(ctx sdk.Context, contractAddr sdk.AccAddress, register *RegisterICA, ik icacontrollerkeeper.Keeper) ([]sdk.Event, [][]byte, error) {
	_, err := PerformRegisterICA(ik, ctx, contractAddr, register)
	if err != nil {
		return nil, nil, errors.Wrap(err, "perform register ICA")
	}
	// Construct an sdk.Event from the MsgRegisterInterchainAccountResponse.
	// Somewhat hacky way to get the data back to the contract.
	// attrs := []sdk.Attribute{
	// 	sdk.NewAttribute()
	return nil, nil, nil
}

// PerformRegisterICA is used with register to validate the register message and register the ICA.
func PerformRegisterICA(f icacontrollerkeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress, register *RegisterICA) (*icacontrollertypes.MsgRegisterInterchainAccountResponse, error) {
	if register == nil {
		return nil, wasmvmtypes.InvalidRequest{Err: "register ICA null message"}
	}

	msgServer := icacontrollerkeeper.NewMsgServerImpl(&f)

	// format "{owner}-{id}"
	owner := contractAddr.String() + "-" + register.AccountId
	msgRegister := icacontrollertypes.NewMsgRegisterInterchainAccount(register.ChannelId, owner, register.Version)

	if err := msgRegister.ValidateBasic(); err != nil {
		return nil, errors.Wrap(err, "failed validating MsgRegisterInterchainAccount")
	}

	res, err := msgServer.RegisterInterchainAccount(
		sdk.WrapSDKContext(ctx),
		msgRegister,
	)

	if err != nil {
		return nil, errors.Wrap(err, "registering ICA")
	}
	return res, nil
}

func submitTxs(ctx sdk.Context, contractAddr sdk.AccAddress, submitTx *SubmitTx, ik icacontrollerkeeper.Keeper) ([]sdk.Event, [][]byte, error) {
	_, err := PerformSubmitTx(ik, ctx, contractAddr, submitTx)
	if err != nil {
		return nil, nil, errors.Wrap(err, "perform submit txs")
	}
	return nil, nil, nil
}

// PerformSubmitTxs is used with submitTxs to validate the submitTxs message and submit the txs.
func PerformSubmitTx(f icacontrollerkeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress, submitTxs *SubmitTx) (*icacontrollertypes.MsgSendTxResponse, error) {
	if submitTxs == nil {
		return nil, wasmvmtypes.InvalidRequest{Err: "submit txs null message"}
	}

	msgServer := icacontrollerkeeper.NewMsgServerImpl(&f)

	owner := contractAddr.String() + "-" + submitTxs.AccountId
	res, err := msgServer.SendTx(sdk.WrapSDKContext(ctx), icacontrollertypes.NewMsgSendTx(owner, submitTxs.ChannelId, submitTxs.Timeout, submitTx.Tx))
	if err != nil {
		return nil, errors.Wrap(err, "submitting txs")
	}
	return res, nil
}