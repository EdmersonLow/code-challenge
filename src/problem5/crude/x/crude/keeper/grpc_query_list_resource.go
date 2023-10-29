package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/crude/x/crude/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListResource(goCtx context.Context, req *types.QueryListResourceRequest) (*types.QueryListResourceResponse, error) {
	if req == nil {
        return nil, status.Error(codes.InvalidArgument, "invalid request")
    }

    var resources []types.Resource
    ctx := sdk.UnwrapSDKContext(goCtx)

    store := ctx.KVStore(k.storeKey)
    postStore := prefix.NewStore(store, types.KeyPrefix(types.PostKey))

    pageRes, err := query.Paginate(postStore, req.Pagination, func(key []byte, value []byte) error {
        var resource types.Resource
        if err := k.cdc.Unmarshal(value, &resource); err != nil {
            return err
        }

        resources = append(resources, resource)
        return nil
    })

    if err != nil {
        return nil, status.Error(codes.Internal, err.Error())
    }

    return &types.QueryListResourceResponse{Resource: resources, Pagination: pageRes}, nil
}
