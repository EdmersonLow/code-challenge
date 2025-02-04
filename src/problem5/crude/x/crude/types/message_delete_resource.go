package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDeleteResource = "delete_resource"

var _ sdk.Msg = &MsgDeleteResource{}

func NewMsgDeleteResource(creator string, id uint64) *MsgDeleteResource {
	return &MsgDeleteResource{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgDeleteResource) Route() string {
	return RouterKey
}

func (msg *MsgDeleteResource) Type() string {
	return TypeMsgDeleteResource
}

func (msg *MsgDeleteResource) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteResource) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteResource) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
