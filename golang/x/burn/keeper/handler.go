package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/summa-tx/burning-bitcoin/golang/x/burn/types"
)

// NewHandler returns a handler for "nameservice" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case types.MsgBurnProof:
			return handleMsgBurnProof(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized nameservice Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// Handle a message to burn proof
func handleMsgBurnProof(ctx sdk.Context, keeper Keeper, msg types.MsgBurnProof) sdk.Result {
	keeper.setValidated(ctx, msg.Proof)
	addr, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		panic("bad bech32 address in Address")
	}
	keeper.AppendAddr(ctx, addr)

	// TODO: hook in IBC

	return sdk.Result{}
}
