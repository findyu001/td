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

// AccountGetPasswordRequest represents TL type `account.getPassword#548a30f5`.
type AccountGetPasswordRequest struct {
}

// AccountGetPasswordRequestTypeID is TL type id of AccountGetPasswordRequest.
const AccountGetPasswordRequestTypeID = 0x548a30f5

// Encode implements bin.Encoder.
func (g *AccountGetPasswordRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getPassword#548a30f5 as nil")
	}
	b.PutID(AccountGetPasswordRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (g *AccountGetPasswordRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getPassword#548a30f5 to nil")
	}
	if err := b.ConsumeID(AccountGetPasswordRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getPassword#548a30f5: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountGetPasswordRequest.
var (
	_ bin.Encoder = &AccountGetPasswordRequest{}
	_ bin.Decoder = &AccountGetPasswordRequest{}
)
