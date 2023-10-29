package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/crude/x/crude/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdateResource(goCtx context.Context, msg *types.MsgUpdateResource) (*types.MsgUpdateResourceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	var post = types.Resource{
        Creatorr: msg.Creator,
        Idcreator:      msg.Id,
        Title:   msg.Title,
        Body:    msg.Body,
    }
    val, found := k.GetResource(ctx, msg.Id)
    if !found {
        return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
    }
    if msg.Creator != val.Creatorr {
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
    }
    k.SetResource(ctx, post)

	return &types.MsgUpdateResourceResponse{}, nil
}
