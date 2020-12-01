// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// ChannelsEditCreatorRequest represents TL type `channels.editCreator#8f38cd1f`.
type ChannelsEditCreatorRequest struct {
	// Channel field of ChannelsEditCreatorRequest.
	Channel InputChannelClass
	// UserID field of ChannelsEditCreatorRequest.
	UserID InputUserClass
	// Password field of ChannelsEditCreatorRequest.
	Password InputCheckPasswordSRPClass
}

// ChannelsEditCreatorRequestTypeID is TL type id of ChannelsEditCreatorRequest.
const ChannelsEditCreatorRequestTypeID = 0x8f38cd1f

// Encode implements bin.Encoder.
func (e *ChannelsEditCreatorRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode channels.editCreator#8f38cd1f as nil")
	}
	b.PutID(ChannelsEditCreatorRequestTypeID)
	if e.Channel == nil {
		return fmt.Errorf("unable to encode channels.editCreator#8f38cd1f: field channel is nil")
	}
	if err := e.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.editCreator#8f38cd1f: field channel: %w", err)
	}
	if e.UserID == nil {
		return fmt.Errorf("unable to encode channels.editCreator#8f38cd1f: field user_id is nil")
	}
	if err := e.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.editCreator#8f38cd1f: field user_id: %w", err)
	}
	if e.Password == nil {
		return fmt.Errorf("unable to encode channels.editCreator#8f38cd1f: field password is nil")
	}
	if err := e.Password.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.editCreator#8f38cd1f: field password: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *ChannelsEditCreatorRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode channels.editCreator#8f38cd1f to nil")
	}
	if err := b.ConsumeID(ChannelsEditCreatorRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.editCreator#8f38cd1f: %w", err)
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.editCreator#8f38cd1f: field channel: %w", err)
		}
		e.Channel = value
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.editCreator#8f38cd1f: field user_id: %w", err)
		}
		e.UserID = value
	}
	{
		value, err := DecodeInputCheckPasswordSRP(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.editCreator#8f38cd1f: field password: %w", err)
		}
		e.Password = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsEditCreatorRequest.
var (
	_ bin.Encoder = &ChannelsEditCreatorRequest{}
	_ bin.Decoder = &ChannelsEditCreatorRequest{}
)

// ChannelsEditCreator invokes method channels.editCreator#8f38cd1f returning error if any.
func (c *Client) ChannelsEditCreator(ctx context.Context, request *ChannelsEditCreatorRequest) (UpdatesClass, error) {
	var result UpdatesBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
