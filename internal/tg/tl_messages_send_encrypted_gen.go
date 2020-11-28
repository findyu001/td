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

// MessagesSendEncryptedRequest represents TL type `messages.sendEncrypted#44fa7a15`.
type MessagesSendEncryptedRequest struct {
	// Flags field of MessagesSendEncryptedRequest.
	Flags bin.Fields
	// Silent field of MessagesSendEncryptedRequest.
	Silent bool
	// Peer field of MessagesSendEncryptedRequest.
	Peer InputEncryptedChat
	// RandomID field of MessagesSendEncryptedRequest.
	RandomID int64
	// Data field of MessagesSendEncryptedRequest.
	Data []byte
}

// MessagesSendEncryptedRequestTypeID is TL type id of MessagesSendEncryptedRequest.
const MessagesSendEncryptedRequestTypeID = 0x44fa7a15

// Encode implements bin.Encoder.
func (s *MessagesSendEncryptedRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sendEncrypted#44fa7a15 as nil")
	}
	b.PutID(MessagesSendEncryptedRequestTypeID)
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendEncrypted#44fa7a15: field flags: %w", err)
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendEncrypted#44fa7a15: field peer: %w", err)
	}
	b.PutLong(s.RandomID)
	b.PutBytes(s.Data)
	return nil
}

// SetSilent sets value of Silent conditional field.
func (s *MessagesSendEncryptedRequest) SetSilent(value bool) {
	if value {
		s.Flags.Set(0)
	} else {
		s.Flags.Unset(0)
	}
}

// Decode implements bin.Decoder.
func (s *MessagesSendEncryptedRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sendEncrypted#44fa7a15 to nil")
	}
	if err := b.ConsumeID(MessagesSendEncryptedRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.sendEncrypted#44fa7a15: %w", err)
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.sendEncrypted#44fa7a15: field flags: %w", err)
		}
	}
	s.Silent = s.Flags.Has(0)
	{
		if err := s.Peer.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.sendEncrypted#44fa7a15: field peer: %w", err)
		}
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendEncrypted#44fa7a15: field random_id: %w", err)
		}
		s.RandomID = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendEncrypted#44fa7a15: field data: %w", err)
		}
		s.Data = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesSendEncryptedRequest.
var (
	_ bin.Encoder = &MessagesSendEncryptedRequest{}
	_ bin.Decoder = &MessagesSendEncryptedRequest{}
)
