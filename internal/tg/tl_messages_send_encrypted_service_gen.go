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

// MessagesSendEncryptedServiceRequest represents TL type `messages.sendEncryptedService#32d439a4`.
type MessagesSendEncryptedServiceRequest struct {
	// Peer field of MessagesSendEncryptedServiceRequest.
	Peer InputEncryptedChat
	// RandomID field of MessagesSendEncryptedServiceRequest.
	RandomID int64
	// Data field of MessagesSendEncryptedServiceRequest.
	Data []byte
}

// MessagesSendEncryptedServiceRequestTypeID is TL type id of MessagesSendEncryptedServiceRequest.
const MessagesSendEncryptedServiceRequestTypeID = 0x32d439a4

// Encode implements bin.Encoder.
func (s *MessagesSendEncryptedServiceRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sendEncryptedService#32d439a4 as nil")
	}
	b.PutID(MessagesSendEncryptedServiceRequestTypeID)
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendEncryptedService#32d439a4: field peer: %w", err)
	}
	b.PutLong(s.RandomID)
	b.PutBytes(s.Data)
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSendEncryptedServiceRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sendEncryptedService#32d439a4 to nil")
	}
	if err := b.ConsumeID(MessagesSendEncryptedServiceRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.sendEncryptedService#32d439a4: %w", err)
	}
	{
		if err := s.Peer.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.sendEncryptedService#32d439a4: field peer: %w", err)
		}
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendEncryptedService#32d439a4: field random_id: %w", err)
		}
		s.RandomID = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendEncryptedService#32d439a4: field data: %w", err)
		}
		s.Data = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesSendEncryptedServiceRequest.
var (
	_ bin.Encoder = &MessagesSendEncryptedServiceRequest{}
	_ bin.Decoder = &MessagesSendEncryptedServiceRequest{}
)
