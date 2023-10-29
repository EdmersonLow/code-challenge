package keeper

import (
    "encoding/binary"
	"github.com/crude/x/crude/types"
    "github.com/cosmos/cosmos-sdk/store/prefix"
    sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AppendResource(ctx sdk.Context, resource types.Resource) uint64 {
    count := k.GetResourceCount(ctx)
    resource.Idcreator = count
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
    appendedValue := k.cdc.MustMarshal(&resource)
    store.Set(GetResourceIDBytes(resource.Idcreator), appendedValue)
    k.SetResourceCount(ctx, count+1)
    return count
}


func (k Keeper) GetResourceCount(ctx sdk.Context) uint64 {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
    byteKey := types.KeyPrefix(types.PostCountKey)
    bz := store.Get(byteKey)
    if bz == nil {
        return 0
    }
    return binary.BigEndian.Uint64(bz)
}
func GetResourceIDBytes(id uint64) []byte {
    bz := make([]byte, 8)
    binary.BigEndian.PutUint64(bz, id)
    return bz
}

func (k Keeper) SetResourceCount(ctx sdk.Context, count uint64) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
    byteKey := types.KeyPrefix(types.PostCountKey)
    bz := make([]byte, 8)
    binary.BigEndian.PutUint64(bz, count)
    store.Set(byteKey, bz)
}

func (k Keeper) GetResource(ctx sdk.Context, id uint64) (val types.Resource, found bool) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
    b := store.Get(GetResourceIDBytes(id))
    if b == nil {
        return val, false
    }
    k.cdc.MustUnmarshal(b, &val)
    return val, true

}

func (k Keeper) SetResource(ctx sdk.Context, post types.Resource) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
    b := k.cdc.MustMarshal(&post)
    store.Set(GetResourceIDBytes(post.Idcreator), b)
}