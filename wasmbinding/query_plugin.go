package wasmbinding

import (
	"encoding/json"

	"kujira/wasmbinding/bindings"
	denom "kujira/x/denom/wasm"
	oracle "kujira/x/oracle/wasm"

	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// CustomQuerier dispatches custom CosmWasm bindings queries.
func CustomQuerier(qp *QueryPlugin) func(ctx sdk.Context, request json.RawMessage) ([]byte, error) {
	return func(ctx sdk.Context, request json.RawMessage) ([]byte, error) {
		var contractQuery bindings.CosmosQuery
		if err := json.Unmarshal(request, &contractQuery); err != nil {
			return nil, sdkerrors.Wrap(err, "kujira query")
		}

		if contractQuery.Oracle != nil {
			res, err := oracle.Handle(qp.oraclekeeper, ctx, contractQuery.Oracle)
			if err != nil {
				return nil, err
			}

			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
			}

			return bz, nil
		} else if contractQuery.Bank != nil {
			coin := qp.bankkeeper.GetSupply(ctx, contractQuery.Bank.Supply.Denom)
			res := banktypes.QuerySupplyOfResponse{
				Amount: coin,
			}

			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
			}

			return bz, nil
		} else if contractQuery.Denom != nil {
			res, err := denom.HandleQuery(qp.denomKeeper, ctx, contractQuery.Denom)
			if err != nil {
				return nil, err
			}

			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
			}

			return bz, nil
		} else {
			return nil, wasmvmtypes.UnsupportedRequest{Kind: "unknown Custom variant"}
		}
	}
}
