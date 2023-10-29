package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/crude/x/crude/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowResource(goCtx context.Context, req *types.QueryShowResourceRequest) (*types.QueryShowResourceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	resource, found := k.GetResource(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryShowResourceResponse{Resource: resource}, nil
}

// $ blogd tx blog create-post hello world --from alice
// blogd q blog show-post 0