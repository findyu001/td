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

// MessagesRequestUrlAuthRequest represents TL type `messages.requestUrlAuth#e33f5613`.
type MessagesRequestUrlAuthRequest struct {
	// Peer field of MessagesRequestUrlAuthRequest.
	Peer InputPeerClass
	// MsgID field of MessagesRequestUrlAuthRequest.
	MsgID int
	// ButtonID field of MessagesRequestUrlAuthRequest.
	ButtonID int
}

// MessagesRequestUrlAuthRequestTypeID is TL type id of MessagesRequestUrlAuthRequest.
const MessagesRequestUrlAuthRequestTypeID = 0xe33f5613

// Encode implements bin.Encoder.
func (r *MessagesRequestUrlAuthRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.requestUrlAuth#e33f5613 as nil")
	}
	b.PutID(MessagesRequestUrlAuthRequestTypeID)
	if r.Peer == nil {
		return fmt.Errorf("unable to encode messages.requestUrlAuth#e33f5613: field peer is nil")
	}
	if err := r.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.requestUrlAuth#e33f5613: field peer: %w", err)
	}
	b.PutInt(r.MsgID)
	b.PutInt(r.ButtonID)
	return nil
}

// Decode implements bin.Decoder.
func (r *MessagesRequestUrlAuthRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.requestUrlAuth#e33f5613 to nil")
	}
	if err := b.ConsumeID(MessagesRequestUrlAuthRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.requestUrlAuth#e33f5613: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.requestUrlAuth#e33f5613: field peer: %w", err)
		}
		r.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.requestUrlAuth#e33f5613: field msg_id: %w", err)
		}
		r.MsgID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.requestUrlAuth#e33f5613: field button_id: %w", err)
		}
		r.ButtonID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesRequestUrlAuthRequest.
var (
	_ bin.Encoder = &MessagesRequestUrlAuthRequest{}
	_ bin.Decoder = &MessagesRequestUrlAuthRequest{}
)
