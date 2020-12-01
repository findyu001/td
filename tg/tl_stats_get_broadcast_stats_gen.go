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

// StatsGetBroadcastStatsRequest represents TL type `stats.getBroadcastStats#ab42441a`.
type StatsGetBroadcastStatsRequest struct {
	// Flags field of StatsGetBroadcastStatsRequest.
	Flags bin.Fields
	// Dark field of StatsGetBroadcastStatsRequest.
	Dark bool
	// Channel field of StatsGetBroadcastStatsRequest.
	Channel InputChannelClass
}

// StatsGetBroadcastStatsRequestTypeID is TL type id of StatsGetBroadcastStatsRequest.
const StatsGetBroadcastStatsRequestTypeID = 0xab42441a

// Encode implements bin.Encoder.
func (g *StatsGetBroadcastStatsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode stats.getBroadcastStats#ab42441a as nil")
	}
	b.PutID(StatsGetBroadcastStatsRequestTypeID)
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.getBroadcastStats#ab42441a: field flags: %w", err)
	}
	if g.Channel == nil {
		return fmt.Errorf("unable to encode stats.getBroadcastStats#ab42441a: field channel is nil")
	}
	if err := g.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.getBroadcastStats#ab42441a: field channel: %w", err)
	}
	return nil
}

// SetDark sets value of Dark conditional field.
func (g *StatsGetBroadcastStatsRequest) SetDark(value bool) {
	if value {
		g.Flags.Set(0)
	} else {
		g.Flags.Unset(0)
	}
}

// Decode implements bin.Decoder.
func (g *StatsGetBroadcastStatsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode stats.getBroadcastStats#ab42441a to nil")
	}
	if err := b.ConsumeID(StatsGetBroadcastStatsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stats.getBroadcastStats#ab42441a: %w", err)
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stats.getBroadcastStats#ab42441a: field flags: %w", err)
		}
	}
	g.Dark = g.Flags.Has(0)
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.getBroadcastStats#ab42441a: field channel: %w", err)
		}
		g.Channel = value
	}
	return nil
}

// Ensuring interfaces in compile-time for StatsGetBroadcastStatsRequest.
var (
	_ bin.Encoder = &StatsGetBroadcastStatsRequest{}
	_ bin.Decoder = &StatsGetBroadcastStatsRequest{}
)

// StatsGetBroadcastStats invokes method stats.getBroadcastStats#ab42441a returning error if any.
func (c *Client) StatsGetBroadcastStats(ctx context.Context, request *StatsGetBroadcastStatsRequest) (*StatsBroadcastStats, error) {
	var result StatsBroadcastStats
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
