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

// UpdatesGetChannelDifferenceRequest represents TL type `updates.getChannelDifference#3173d78`.
type UpdatesGetChannelDifferenceRequest struct {
	// Flags field of UpdatesGetChannelDifferenceRequest.
	Flags bin.Fields
	// Force field of UpdatesGetChannelDifferenceRequest.
	Force bool
	// Channel field of UpdatesGetChannelDifferenceRequest.
	Channel InputChannelClass
	// Filter field of UpdatesGetChannelDifferenceRequest.
	Filter ChannelMessagesFilterClass
	// Pts field of UpdatesGetChannelDifferenceRequest.
	Pts int
	// Limit field of UpdatesGetChannelDifferenceRequest.
	Limit int
}

// UpdatesGetChannelDifferenceRequestTypeID is TL type id of UpdatesGetChannelDifferenceRequest.
const UpdatesGetChannelDifferenceRequestTypeID = 0x3173d78

// Encode implements bin.Encoder.
func (g *UpdatesGetChannelDifferenceRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode updates.getChannelDifference#3173d78 as nil")
	}
	b.PutID(UpdatesGetChannelDifferenceRequestTypeID)
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode updates.getChannelDifference#3173d78: field flags: %w", err)
	}
	if g.Channel == nil {
		return fmt.Errorf("unable to encode updates.getChannelDifference#3173d78: field channel is nil")
	}
	if err := g.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode updates.getChannelDifference#3173d78: field channel: %w", err)
	}
	if g.Filter == nil {
		return fmt.Errorf("unable to encode updates.getChannelDifference#3173d78: field filter is nil")
	}
	if err := g.Filter.Encode(b); err != nil {
		return fmt.Errorf("unable to encode updates.getChannelDifference#3173d78: field filter: %w", err)
	}
	b.PutInt(g.Pts)
	b.PutInt(g.Limit)
	return nil
}

// SetForce sets value of Force conditional field.
func (g *UpdatesGetChannelDifferenceRequest) SetForce(value bool) {
	if value {
		g.Flags.Set(0)
	} else {
		g.Flags.Unset(0)
	}
}

// Decode implements bin.Decoder.
func (g *UpdatesGetChannelDifferenceRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode updates.getChannelDifference#3173d78 to nil")
	}
	if err := b.ConsumeID(UpdatesGetChannelDifferenceRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode updates.getChannelDifference#3173d78: %w", err)
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode updates.getChannelDifference#3173d78: field flags: %w", err)
		}
	}
	g.Force = g.Flags.Has(0)
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode updates.getChannelDifference#3173d78: field channel: %w", err)
		}
		g.Channel = value
	}
	{
		value, err := DecodeChannelMessagesFilter(b)
		if err != nil {
			return fmt.Errorf("unable to decode updates.getChannelDifference#3173d78: field filter: %w", err)
		}
		g.Filter = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.getChannelDifference#3173d78: field pts: %w", err)
		}
		g.Pts = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.getChannelDifference#3173d78: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// Ensuring interfaces in compile-time for UpdatesGetChannelDifferenceRequest.
var (
	_ bin.Encoder = &UpdatesGetChannelDifferenceRequest{}
	_ bin.Decoder = &UpdatesGetChannelDifferenceRequest{}
)
