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

// MessagesGetUnreadMentionsRequest represents TL type `messages.getUnreadMentions#46578472`.
type MessagesGetUnreadMentionsRequest struct {
	// Peer field of MessagesGetUnreadMentionsRequest.
	Peer InputPeerClass
	// OffsetID field of MessagesGetUnreadMentionsRequest.
	OffsetID int
	// AddOffset field of MessagesGetUnreadMentionsRequest.
	AddOffset int
	// Limit field of MessagesGetUnreadMentionsRequest.
	Limit int
	// MaxID field of MessagesGetUnreadMentionsRequest.
	MaxID int
	// MinID field of MessagesGetUnreadMentionsRequest.
	MinID int
}

// MessagesGetUnreadMentionsRequestTypeID is TL type id of MessagesGetUnreadMentionsRequest.
const MessagesGetUnreadMentionsRequestTypeID = 0x46578472

// Encode implements bin.Encoder.
func (g *MessagesGetUnreadMentionsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getUnreadMentions#46578472 as nil")
	}
	b.PutID(MessagesGetUnreadMentionsRequestTypeID)
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getUnreadMentions#46578472: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getUnreadMentions#46578472: field peer: %w", err)
	}
	b.PutInt(g.OffsetID)
	b.PutInt(g.AddOffset)
	b.PutInt(g.Limit)
	b.PutInt(g.MaxID)
	b.PutInt(g.MinID)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetUnreadMentionsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getUnreadMentions#46578472 to nil")
	}
	if err := b.ConsumeID(MessagesGetUnreadMentionsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getUnreadMentions#46578472: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getUnreadMentions#46578472: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getUnreadMentions#46578472: field offset_id: %w", err)
		}
		g.OffsetID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getUnreadMentions#46578472: field add_offset: %w", err)
		}
		g.AddOffset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getUnreadMentions#46578472: field limit: %w", err)
		}
		g.Limit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getUnreadMentions#46578472: field max_id: %w", err)
		}
		g.MaxID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getUnreadMentions#46578472: field min_id: %w", err)
		}
		g.MinID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetUnreadMentionsRequest.
var (
	_ bin.Encoder = &MessagesGetUnreadMentionsRequest{}
	_ bin.Decoder = &MessagesGetUnreadMentionsRequest{}
)

// MessagesGetUnreadMentions invokes method messages.getUnreadMentions#46578472 returning error if any.
func (c *Client) MessagesGetUnreadMentions(ctx context.Context, request *MessagesGetUnreadMentionsRequest) (MessagesMessagesClass, error) {
	var result MessagesMessagesBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Messages, nil
}
