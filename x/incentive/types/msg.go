package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgClaimJPYXMintingReward{}

// NewMsgClaimJPYXMintingReward returns a new MsgClaimJPYXMintingReward.
func NewMsgClaimJPYXMintingReward(sender sdk.AccAddress, multiplierName string) MsgClaimJPYXMintingReward {
	return MsgClaimJPYXMintingReward{
		Sender:         sender.Bytes(),
		MultiplierName: multiplierName,
	}
}

// Route return the message type used for routing the message.
func (msg MsgClaimJPYXMintingReward) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgClaimJPYXMintingReward) Type() string { return "claim_jpyx_minting_reward" }

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgClaimJPYXMintingReward) ValidateBasic() error {
	if msg.Sender.AccAddress().Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address cannot be empty")
	}
	return MultiplierName(strings.ToLower(msg.MultiplierName)).IsValid()
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgClaimJPYXMintingReward) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgClaimJPYXMintingReward) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender.AccAddress()}
}