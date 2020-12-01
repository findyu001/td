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

// ChannelsReportSpamRequest represents TL type `channels.reportSpam#fe087810`.
type ChannelsReportSpamRequest struct {
	// Channel field of ChannelsReportSpamRequest.
	Channel InputChannelClass
	// UserID field of ChannelsReportSpamRequest.
	UserID InputUserClass
	// ID field of ChannelsReportSpamRequest.
	ID []int
}

// ChannelsReportSpamRequestTypeID is TL type id of ChannelsReportSpamRequest.
const ChannelsReportSpamRequestTypeID = 0xfe087810

// Encode implements bin.Encoder.
func (r *ChannelsReportSpamRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode channels.reportSpam#fe087810 as nil")
	}
	b.PutID(ChannelsReportSpamRequestTypeID)
	if r.Channel == nil {
		return fmt.Errorf("unable to encode channels.reportSpam#fe087810: field channel is nil")
	}
	if err := r.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.reportSpam#fe087810: field channel: %w", err)
	}
	if r.UserID == nil {
		return fmt.Errorf("unable to encode channels.reportSpam#fe087810: field user_id is nil")
	}
	if err := r.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.reportSpam#fe087810: field user_id: %w", err)
	}
	b.PutVectorHeader(len(r.ID))
	for _, v := range r.ID {
		b.PutInt(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *ChannelsReportSpamRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode channels.reportSpam#fe087810 to nil")
	}
	if err := b.ConsumeID(ChannelsReportSpamRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.reportSpam#fe087810: %w", err)
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.reportSpam#fe087810: field channel: %w", err)
		}
		r.Channel = value
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.reportSpam#fe087810: field user_id: %w", err)
		}
		r.UserID = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode channels.reportSpam#fe087810: field id: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode channels.reportSpam#fe087810: field id: %w", err)
			}
			r.ID = append(r.ID, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsReportSpamRequest.
var (
	_ bin.Encoder = &ChannelsReportSpamRequest{}
	_ bin.Decoder = &ChannelsReportSpamRequest{}
)

// ChannelsReportSpam invokes method channels.reportSpam#fe087810 returning error if any.
func (c *Client) ChannelsReportSpam(ctx context.Context, request *ChannelsReportSpamRequest) (BoolClass, error) {
	var result BoolBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Bool, nil
}
