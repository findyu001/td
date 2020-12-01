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

// AccountConfirmPhoneRequest represents TL type `account.confirmPhone#5f2178c3`.
type AccountConfirmPhoneRequest struct {
	// PhoneCodeHash field of AccountConfirmPhoneRequest.
	PhoneCodeHash string
	// PhoneCode field of AccountConfirmPhoneRequest.
	PhoneCode string
}

// AccountConfirmPhoneRequestTypeID is TL type id of AccountConfirmPhoneRequest.
const AccountConfirmPhoneRequestTypeID = 0x5f2178c3

// Encode implements bin.Encoder.
func (c *AccountConfirmPhoneRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode account.confirmPhone#5f2178c3 as nil")
	}
	b.PutID(AccountConfirmPhoneRequestTypeID)
	b.PutString(c.PhoneCodeHash)
	b.PutString(c.PhoneCode)
	return nil
}

// Decode implements bin.Decoder.
func (c *AccountConfirmPhoneRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode account.confirmPhone#5f2178c3 to nil")
	}
	if err := b.ConsumeID(AccountConfirmPhoneRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.confirmPhone#5f2178c3: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.confirmPhone#5f2178c3: field phone_code_hash: %w", err)
		}
		c.PhoneCodeHash = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.confirmPhone#5f2178c3: field phone_code: %w", err)
		}
		c.PhoneCode = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountConfirmPhoneRequest.
var (
	_ bin.Encoder = &AccountConfirmPhoneRequest{}
	_ bin.Decoder = &AccountConfirmPhoneRequest{}
)

// AccountConfirmPhone invokes method account.confirmPhone#5f2178c3 returning error if any.
func (c *Client) AccountConfirmPhone(ctx context.Context, request *AccountConfirmPhoneRequest) (BoolClass, error) {
	var result BoolBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Bool, nil
}
