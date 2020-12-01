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

// MessagesDeleteChatUserRequest represents TL type `messages.deleteChatUser#e0611f16`.
type MessagesDeleteChatUserRequest struct {
	// ChatID field of MessagesDeleteChatUserRequest.
	ChatID int
	// UserID field of MessagesDeleteChatUserRequest.
	UserID InputUserClass
}

// MessagesDeleteChatUserRequestTypeID is TL type id of MessagesDeleteChatUserRequest.
const MessagesDeleteChatUserRequestTypeID = 0xe0611f16

// Encode implements bin.Encoder.
func (d *MessagesDeleteChatUserRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode messages.deleteChatUser#e0611f16 as nil")
	}
	b.PutID(MessagesDeleteChatUserRequestTypeID)
	b.PutInt(d.ChatID)
	if d.UserID == nil {
		return fmt.Errorf("unable to encode messages.deleteChatUser#e0611f16: field user_id is nil")
	}
	if err := d.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.deleteChatUser#e0611f16: field user_id: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *MessagesDeleteChatUserRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode messages.deleteChatUser#e0611f16 to nil")
	}
	if err := b.ConsumeID(MessagesDeleteChatUserRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.deleteChatUser#e0611f16: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.deleteChatUser#e0611f16: field chat_id: %w", err)
		}
		d.ChatID = value
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.deleteChatUser#e0611f16: field user_id: %w", err)
		}
		d.UserID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesDeleteChatUserRequest.
var (
	_ bin.Encoder = &MessagesDeleteChatUserRequest{}
	_ bin.Decoder = &MessagesDeleteChatUserRequest{}
)

// MessagesDeleteChatUser invokes method messages.deleteChatUser#e0611f16 returning error if any.
func (c *Client) MessagesDeleteChatUser(ctx context.Context, request *MessagesDeleteChatUserRequest) (UpdatesClass, error) {
	var result UpdatesBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
