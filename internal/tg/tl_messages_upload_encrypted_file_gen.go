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

// MessagesUploadEncryptedFileRequest represents TL type `messages.uploadEncryptedFile#5057c497`.
type MessagesUploadEncryptedFileRequest struct {
	// Peer field of MessagesUploadEncryptedFileRequest.
	Peer InputEncryptedChat
	// File field of MessagesUploadEncryptedFileRequest.
	File InputEncryptedFileClass
}

// MessagesUploadEncryptedFileRequestTypeID is TL type id of MessagesUploadEncryptedFileRequest.
const MessagesUploadEncryptedFileRequestTypeID = 0x5057c497

// Encode implements bin.Encoder.
func (u *MessagesUploadEncryptedFileRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode messages.uploadEncryptedFile#5057c497 as nil")
	}
	b.PutID(MessagesUploadEncryptedFileRequestTypeID)
	if err := u.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.uploadEncryptedFile#5057c497: field peer: %w", err)
	}
	if u.File == nil {
		return fmt.Errorf("unable to encode messages.uploadEncryptedFile#5057c497: field file is nil")
	}
	if err := u.File.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.uploadEncryptedFile#5057c497: field file: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *MessagesUploadEncryptedFileRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode messages.uploadEncryptedFile#5057c497 to nil")
	}
	if err := b.ConsumeID(MessagesUploadEncryptedFileRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.uploadEncryptedFile#5057c497: %w", err)
	}
	{
		if err := u.Peer.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.uploadEncryptedFile#5057c497: field peer: %w", err)
		}
	}
	{
		value, err := DecodeInputEncryptedFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.uploadEncryptedFile#5057c497: field file: %w", err)
		}
		u.File = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesUploadEncryptedFileRequest.
var (
	_ bin.Encoder = &MessagesUploadEncryptedFileRequest{}
	_ bin.Decoder = &MessagesUploadEncryptedFileRequest{}
)
