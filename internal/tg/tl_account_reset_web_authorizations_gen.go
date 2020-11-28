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

// AccountResetWebAuthorizationsRequest represents TL type `account.resetWebAuthorizations#682d2594`.
type AccountResetWebAuthorizationsRequest struct {
}

// AccountResetWebAuthorizationsRequestTypeID is TL type id of AccountResetWebAuthorizationsRequest.
const AccountResetWebAuthorizationsRequestTypeID = 0x682d2594

// Encode implements bin.Encoder.
func (r *AccountResetWebAuthorizationsRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode account.resetWebAuthorizations#682d2594 as nil")
	}
	b.PutID(AccountResetWebAuthorizationsRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (r *AccountResetWebAuthorizationsRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode account.resetWebAuthorizations#682d2594 to nil")
	}
	if err := b.ConsumeID(AccountResetWebAuthorizationsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.resetWebAuthorizations#682d2594: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountResetWebAuthorizationsRequest.
var (
	_ bin.Encoder = &AccountResetWebAuthorizationsRequest{}
	_ bin.Decoder = &AccountResetWebAuthorizationsRequest{}
)
