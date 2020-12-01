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

// MessagesStartBotRequest represents TL type `messages.startBot#e6df7378`.
type MessagesStartBotRequest struct {
	// Bot field of MessagesStartBotRequest.
	Bot InputUserClass
	// Peer field of MessagesStartBotRequest.
	Peer InputPeerClass
	// RandomID field of MessagesStartBotRequest.
	RandomID int64
	// StartParam field of MessagesStartBotRequest.
	StartParam string
}

// MessagesStartBotRequestTypeID is TL type id of MessagesStartBotRequest.
const MessagesStartBotRequestTypeID = 0xe6df7378

// Encode implements bin.Encoder.
func (s *MessagesStartBotRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.startBot#e6df7378 as nil")
	}
	b.PutID(MessagesStartBotRequestTypeID)
	if s.Bot == nil {
		return fmt.Errorf("unable to encode messages.startBot#e6df7378: field bot is nil")
	}
	if err := s.Bot.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.startBot#e6df7378: field bot: %w", err)
	}
	if s.Peer == nil {
		return fmt.Errorf("unable to encode messages.startBot#e6df7378: field peer is nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.startBot#e6df7378: field peer: %w", err)
	}
	b.PutLong(s.RandomID)
	b.PutString(s.StartParam)
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesStartBotRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.startBot#e6df7378 to nil")
	}
	if err := b.ConsumeID(MessagesStartBotRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.startBot#e6df7378: %w", err)
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.startBot#e6df7378: field bot: %w", err)
		}
		s.Bot = value
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.startBot#e6df7378: field peer: %w", err)
		}
		s.Peer = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.startBot#e6df7378: field random_id: %w", err)
		}
		s.RandomID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.startBot#e6df7378: field start_param: %w", err)
		}
		s.StartParam = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesStartBotRequest.
var (
	_ bin.Encoder = &MessagesStartBotRequest{}
	_ bin.Decoder = &MessagesStartBotRequest{}
)

// MessagesStartBot invokes method messages.startBot#e6df7378 returning error if any.
func (c *Client) MessagesStartBot(ctx context.Context, request *MessagesStartBotRequest) (UpdatesClass, error) {
	var result UpdatesBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
