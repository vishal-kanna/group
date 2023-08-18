package group

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (p *Proposal) GetMsgs() ([]sdk.Msg, error) {
	msgs := make([]sdk.Msg, len(p.Messages))
	for i, msgAny := range p.Messages {
		msg, ok := msgAny.GetCachedValue().(sdk.Msg)
		if !ok {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "messages contains %T which is not a sdk.MsgRequest", msgAny)
		}
		msgs[i] = msg
	}
	return msgs, nil
}

func (p *Proposal) SetMsgs(msgs []sdk.Msg) error {
	anys := make([]*types.Any, len(msgs))
	for i, msg := range msgs {
		any, err := types.NewAnyWithValue(msg)
		if err != nil {
			return err
		}
		anys[i] = any
	}
	p.Messages = anys
	return nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (p Proposal) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	err := p.UnpackInterfaces(unpacker)
	if err != nil {
		return err
	}
	return nil
}
