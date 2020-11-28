// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"

	"github.com/ernado/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// ChannelsDeleteMessagesRequest represents TL type `channels.deleteMessages#84c1fd4e`.
type ChannelsDeleteMessagesRequest struct {
	// Channel field of ChannelsDeleteMessagesRequest.
	Channel InputChannelClass
	// ID field of ChannelsDeleteMessagesRequest.
	ID []int
}

// ChannelsDeleteMessagesRequestTypeID is TL type id of ChannelsDeleteMessagesRequest.
const ChannelsDeleteMessagesRequestTypeID = 0x84c1fd4e

// Encode implements bin.Encoder.
func (d *ChannelsDeleteMessagesRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode channels.deleteMessages#84c1fd4e as nil")
	}
	b.PutID(ChannelsDeleteMessagesRequestTypeID)
	if d.Channel == nil {
		return fmt.Errorf("unable to encode channels.deleteMessages#84c1fd4e: field channel is nil")
	}
	if err := d.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.deleteMessages#84c1fd4e: field channel: %w", err)
	}
	b.PutVectorHeader(len(d.ID))
	for _, v := range d.ID {
		b.PutInt(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *ChannelsDeleteMessagesRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode channels.deleteMessages#84c1fd4e to nil")
	}
	if err := b.ConsumeID(ChannelsDeleteMessagesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.deleteMessages#84c1fd4e: %w", err)
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.deleteMessages#84c1fd4e: field channel: %w", err)
		}
		d.Channel = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode channels.deleteMessages#84c1fd4e: field id: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode channels.deleteMessages#84c1fd4e: field id: %w", err)
			}
			d.ID = append(d.ID, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsDeleteMessagesRequest.
var (
	_ bin.Encoder = &ChannelsDeleteMessagesRequest{}
	_ bin.Decoder = &ChannelsDeleteMessagesRequest{}
)
