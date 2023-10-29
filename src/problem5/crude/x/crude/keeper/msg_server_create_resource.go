package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/crude/x/crude/types"
)

func (k msgServer) CreateResource(goCtx context.Context, msg *types.MsgCreateResource) (*types.MsgCreateResourceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	var resource = types.Resource{
        Creatorr: msg.Creator,
        Title:   msg.Title,
        Body:    msg.Body,
    }
    id := k.AppendResource(
        ctx,
        resource,
    )
    return &types.MsgCreateResourceResponse{
        Id: id,
    }, nil
}
