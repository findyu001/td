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

// AccountResetWallPapersRequest represents TL type `account.resetWallPapers#bb3b9804`.
type AccountResetWallPapersRequest struct {
}

// AccountResetWallPapersRequestTypeID is TL type id of AccountResetWallPapersRequest.
const AccountResetWallPapersRequestTypeID = 0xbb3b9804

// Encode implements bin.Encoder.
func (r *AccountResetWallPapersRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode account.resetWallPapers#bb3b9804 as nil")
	}
	b.PutID(AccountResetWallPapersRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (r *AccountResetWallPapersRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode account.resetWallPapers#bb3b9804 to nil")
	}
	if err := b.ConsumeID(AccountResetWallPapersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.resetWallPapers#bb3b9804: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountResetWallPapersRequest.
var (
	_ bin.Encoder = &AccountResetWallPapersRequest{}
	_ bin.Decoder = &AccountResetWallPapersRequest{}
)
