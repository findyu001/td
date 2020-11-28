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

// ContactsResetSavedRequest represents TL type `contacts.resetSaved#879537f1`.
type ContactsResetSavedRequest struct {
}

// ContactsResetSavedRequestTypeID is TL type id of ContactsResetSavedRequest.
const ContactsResetSavedRequestTypeID = 0x879537f1

// Encode implements bin.Encoder.
func (r *ContactsResetSavedRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode contacts.resetSaved#879537f1 as nil")
	}
	b.PutID(ContactsResetSavedRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (r *ContactsResetSavedRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode contacts.resetSaved#879537f1 to nil")
	}
	if err := b.ConsumeID(ContactsResetSavedRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.resetSaved#879537f1: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for ContactsResetSavedRequest.
var (
	_ bin.Encoder = &ContactsResetSavedRequest{}
	_ bin.Decoder = &ContactsResetSavedRequest{}
)
