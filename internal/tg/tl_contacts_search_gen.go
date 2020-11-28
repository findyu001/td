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

// ContactsSearchRequest represents TL type `contacts.search#11f812d8`.
type ContactsSearchRequest struct {
	// Q field of ContactsSearchRequest.
	Q string
	// Limit field of ContactsSearchRequest.
	Limit int
}

// ContactsSearchRequestTypeID is TL type id of ContactsSearchRequest.
const ContactsSearchRequestTypeID = 0x11f812d8

// Encode implements bin.Encoder.
func (s *ContactsSearchRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode contacts.search#11f812d8 as nil")
	}
	b.PutID(ContactsSearchRequestTypeID)
	b.PutString(s.Q)
	b.PutInt(s.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (s *ContactsSearchRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode contacts.search#11f812d8 to nil")
	}
	if err := b.ConsumeID(ContactsSearchRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.search#11f812d8: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.search#11f812d8: field q: %w", err)
		}
		s.Q = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.search#11f812d8: field limit: %w", err)
		}
		s.Limit = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ContactsSearchRequest.
var (
	_ bin.Encoder = &ContactsSearchRequest{}
	_ bin.Decoder = &ContactsSearchRequest{}
)
