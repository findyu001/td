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

// MessagesMigrateChatRequest represents TL type `messages.migrateChat#15a3b8e3`.
type MessagesMigrateChatRequest struct {
	// ChatID field of MessagesMigrateChatRequest.
	ChatID int
}

// MessagesMigrateChatRequestTypeID is TL type id of MessagesMigrateChatRequest.
const MessagesMigrateChatRequestTypeID = 0x15a3b8e3

// Encode implements bin.Encoder.
func (m *MessagesMigrateChatRequest) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messages.migrateChat#15a3b8e3 as nil")
	}
	b.PutID(MessagesMigrateChatRequestTypeID)
	b.PutInt(m.ChatID)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessagesMigrateChatRequest) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messages.migrateChat#15a3b8e3 to nil")
	}
	if err := b.ConsumeID(MessagesMigrateChatRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.migrateChat#15a3b8e3: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.migrateChat#15a3b8e3: field chat_id: %w", err)
		}
		m.ChatID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesMigrateChatRequest.
var (
	_ bin.Encoder = &MessagesMigrateChatRequest{}
	_ bin.Decoder = &MessagesMigrateChatRequest{}
)
