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

// MessagesUnpinAllMessagesRequest represents TL type `messages.unpinAllMessages#f025bc8b`.
type MessagesUnpinAllMessagesRequest struct {
	// Peer field of MessagesUnpinAllMessagesRequest.
	Peer InputPeerClass
}

// MessagesUnpinAllMessagesRequestTypeID is TL type id of MessagesUnpinAllMessagesRequest.
const MessagesUnpinAllMessagesRequestTypeID = 0xf025bc8b

// Encode implements bin.Encoder.
func (u *MessagesUnpinAllMessagesRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode messages.unpinAllMessages#f025bc8b as nil")
	}
	b.PutID(MessagesUnpinAllMessagesRequestTypeID)
	if u.Peer == nil {
		return fmt.Errorf("unable to encode messages.unpinAllMessages#f025bc8b: field peer is nil")
	}
	if err := u.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.unpinAllMessages#f025bc8b: field peer: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *MessagesUnpinAllMessagesRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode messages.unpinAllMessages#f025bc8b to nil")
	}
	if err := b.ConsumeID(MessagesUnpinAllMessagesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.unpinAllMessages#f025bc8b: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.unpinAllMessages#f025bc8b: field peer: %w", err)
		}
		u.Peer = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesUnpinAllMessagesRequest.
var (
	_ bin.Encoder = &MessagesUnpinAllMessagesRequest{}
	_ bin.Decoder = &MessagesUnpinAllMessagesRequest{}
)
