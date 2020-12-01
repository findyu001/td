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

// MessagesGetMessageEditDataRequest represents TL type `messages.getMessageEditData#fda68d36`.
type MessagesGetMessageEditDataRequest struct {
	// Peer field of MessagesGetMessageEditDataRequest.
	Peer InputPeerClass
	// ID field of MessagesGetMessageEditDataRequest.
	ID int
}

// MessagesGetMessageEditDataRequestTypeID is TL type id of MessagesGetMessageEditDataRequest.
const MessagesGetMessageEditDataRequestTypeID = 0xfda68d36

// Encode implements bin.Encoder.
func (g *MessagesGetMessageEditDataRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getMessageEditData#fda68d36 as nil")
	}
	b.PutID(MessagesGetMessageEditDataRequestTypeID)
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getMessageEditData#fda68d36: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getMessageEditData#fda68d36: field peer: %w", err)
	}
	b.PutInt(g.ID)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetMessageEditDataRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getMessageEditData#fda68d36 to nil")
	}
	if err := b.ConsumeID(MessagesGetMessageEditDataRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getMessageEditData#fda68d36: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getMessageEditData#fda68d36: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getMessageEditData#fda68d36: field id: %w", err)
		}
		g.ID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetMessageEditDataRequest.
var (
	_ bin.Encoder = &MessagesGetMessageEditDataRequest{}
	_ bin.Decoder = &MessagesGetMessageEditDataRequest{}
)

// MessagesGetMessageEditData invokes method messages.getMessageEditData#fda68d36 returning error if any.
func (c *Client) MessagesGetMessageEditData(ctx context.Context, request *MessagesGetMessageEditDataRequest) (*MessagesMessageEditData, error) {
	var result MessagesMessageEditData
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
