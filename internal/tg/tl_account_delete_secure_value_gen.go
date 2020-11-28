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

// AccountDeleteSecureValueRequest represents TL type `account.deleteSecureValue#b880bc4b`.
type AccountDeleteSecureValueRequest struct {
	// Types field of AccountDeleteSecureValueRequest.
	Types []SecureValueTypeClass
}

// AccountDeleteSecureValueRequestTypeID is TL type id of AccountDeleteSecureValueRequest.
const AccountDeleteSecureValueRequestTypeID = 0xb880bc4b

// Encode implements bin.Encoder.
func (d *AccountDeleteSecureValueRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode account.deleteSecureValue#b880bc4b as nil")
	}
	b.PutID(AccountDeleteSecureValueRequestTypeID)
	b.PutVectorHeader(len(d.Types))
	for idx, v := range d.Types {
		if v == nil {
			return fmt.Errorf("unable to encode account.deleteSecureValue#b880bc4b: field types element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.deleteSecureValue#b880bc4b: field types element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *AccountDeleteSecureValueRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode account.deleteSecureValue#b880bc4b to nil")
	}
	if err := b.ConsumeID(AccountDeleteSecureValueRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.deleteSecureValue#b880bc4b: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.deleteSecureValue#b880bc4b: field types: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeSecureValueType(b)
			if err != nil {
				return fmt.Errorf("unable to decode account.deleteSecureValue#b880bc4b: field types: %w", err)
			}
			d.Types = append(d.Types, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountDeleteSecureValueRequest.
var (
	_ bin.Encoder = &AccountDeleteSecureValueRequest{}
	_ bin.Decoder = &AccountDeleteSecureValueRequest{}
)
