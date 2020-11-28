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

// ContactsGetStatusesRequest represents TL type `contacts.getStatuses#c4a353ee`.
type ContactsGetStatusesRequest struct {
}

// ContactsGetStatusesRequestTypeID is TL type id of ContactsGetStatusesRequest.
const ContactsGetStatusesRequestTypeID = 0xc4a353ee

// Encode implements bin.Encoder.
func (g *ContactsGetStatusesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode contacts.getStatuses#c4a353ee as nil")
	}
	b.PutID(ContactsGetStatusesRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (g *ContactsGetStatusesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode contacts.getStatuses#c4a353ee to nil")
	}
	if err := b.ConsumeID(ContactsGetStatusesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.getStatuses#c4a353ee: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for ContactsGetStatusesRequest.
var (
	_ bin.Encoder = &ContactsGetStatusesRequest{}
	_ bin.Decoder = &ContactsGetStatusesRequest{}
)
