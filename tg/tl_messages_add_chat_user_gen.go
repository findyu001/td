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

// MessagesAddChatUserRequest represents TL type `messages.addChatUser#f9a0aa09`.
type MessagesAddChatUserRequest struct {
	// ChatID field of MessagesAddChatUserRequest.
	ChatID int
	// UserID field of MessagesAddChatUserRequest.
	UserID InputUserClass
	// FwdLimit field of MessagesAddChatUserRequest.
	FwdLimit int
}

// MessagesAddChatUserRequestTypeID is TL type id of MessagesAddChatUserRequest.
const MessagesAddChatUserRequestTypeID = 0xf9a0aa09

// Encode implements bin.Encoder.
func (a *MessagesAddChatUserRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode messages.addChatUser#f9a0aa09 as nil")
	}
	b.PutID(MessagesAddChatUserRequestTypeID)
	b.PutInt(a.ChatID)
	if a.UserID == nil {
		return fmt.Errorf("unable to encode messages.addChatUser#f9a0aa09: field user_id is nil")
	}
	if err := a.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.addChatUser#f9a0aa09: field user_id: %w", err)
	}
	b.PutInt(a.FwdLimit)
	return nil
}

// Decode implements bin.Decoder.
func (a *MessagesAddChatUserRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode messages.addChatUser#f9a0aa09 to nil")
	}
	if err := b.ConsumeID(MessagesAddChatUserRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.addChatUser#f9a0aa09: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.addChatUser#f9a0aa09: field chat_id: %w", err)
		}
		a.ChatID = value
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.addChatUser#f9a0aa09: field user_id: %w", err)
		}
		a.UserID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.addChatUser#f9a0aa09: field fwd_limit: %w", err)
		}
		a.FwdLimit = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesAddChatUserRequest.
var (
	_ bin.Encoder = &MessagesAddChatUserRequest{}
	_ bin.Decoder = &MessagesAddChatUserRequest{}
)

// MessagesAddChatUser invokes method messages.addChatUser#f9a0aa09 returning error if any.
func (c *Client) MessagesAddChatUser(ctx context.Context, request *MessagesAddChatUserRequest) (UpdatesClass, error) {
	var result UpdatesBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
