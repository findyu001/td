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

// ContactsDeleteByPhonesRequest represents TL type `contacts.deleteByPhones#1013fd9e`.
type ContactsDeleteByPhonesRequest struct {
	// Phones field of ContactsDeleteByPhonesRequest.
	Phones []string
}

// ContactsDeleteByPhonesRequestTypeID is TL type id of ContactsDeleteByPhonesRequest.
const ContactsDeleteByPhonesRequestTypeID = 0x1013fd9e

// Encode implements bin.Encoder.
func (d *ContactsDeleteByPhonesRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode contacts.deleteByPhones#1013fd9e as nil")
	}
	b.PutID(ContactsDeleteByPhonesRequestTypeID)
	b.PutVectorHeader(len(d.Phones))
	for _, v := range d.Phones {
		b.PutString(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *ContactsDeleteByPhonesRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode contacts.deleteByPhones#1013fd9e to nil")
	}
	if err := b.ConsumeID(ContactsDeleteByPhonesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.deleteByPhones#1013fd9e: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.deleteByPhones#1013fd9e: field phones: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode contacts.deleteByPhones#1013fd9e: field phones: %w", err)
			}
			d.Phones = append(d.Phones, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for ContactsDeleteByPhonesRequest.
var (
	_ bin.Encoder = &ContactsDeleteByPhonesRequest{}
	_ bin.Decoder = &ContactsDeleteByPhonesRequest{}
)

// ContactsDeleteByPhones invokes method contacts.deleteByPhones#1013fd9e returning error if any.
func (c *Client) ContactsDeleteByPhones(ctx context.Context, request *ContactsDeleteByPhonesRequest) (BoolClass, error) {
	var result BoolBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Bool, nil
}
