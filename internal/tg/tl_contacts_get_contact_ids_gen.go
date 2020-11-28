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

// ContactsGetContactIDsRequest represents TL type `contacts.getContactIDs#2caa4a42`.
type ContactsGetContactIDsRequest struct {
	// Hash field of ContactsGetContactIDsRequest.
	Hash int
}

// ContactsGetContactIDsRequestTypeID is TL type id of ContactsGetContactIDsRequest.
const ContactsGetContactIDsRequestTypeID = 0x2caa4a42

// Encode implements bin.Encoder.
func (g *ContactsGetContactIDsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode contacts.getContactIDs#2caa4a42 as nil")
	}
	b.PutID(ContactsGetContactIDsRequestTypeID)
	b.PutInt(g.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (g *ContactsGetContactIDsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode contacts.getContactIDs#2caa4a42 to nil")
	}
	if err := b.ConsumeID(ContactsGetContactIDsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.getContactIDs#2caa4a42: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.getContactIDs#2caa4a42: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ContactsGetContactIDsRequest.
var (
	_ bin.Encoder = &ContactsGetContactIDsRequest{}
	_ bin.Decoder = &ContactsGetContactIDsRequest{}
)
