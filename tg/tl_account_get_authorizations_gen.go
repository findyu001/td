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

// AccountGetAuthorizationsRequest represents TL type `account.getAuthorizations#e320c158`.
type AccountGetAuthorizationsRequest struct {
}

// AccountGetAuthorizationsRequestTypeID is TL type id of AccountGetAuthorizationsRequest.
const AccountGetAuthorizationsRequestTypeID = 0xe320c158

// Encode implements bin.Encoder.
func (g *AccountGetAuthorizationsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getAuthorizations#e320c158 as nil")
	}
	b.PutID(AccountGetAuthorizationsRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (g *AccountGetAuthorizationsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getAuthorizations#e320c158 to nil")
	}
	if err := b.ConsumeID(AccountGetAuthorizationsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getAuthorizations#e320c158: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountGetAuthorizationsRequest.
var (
	_ bin.Encoder = &AccountGetAuthorizationsRequest{}
	_ bin.Decoder = &AccountGetAuthorizationsRequest{}
)

// AccountGetAuthorizations invokes method account.getAuthorizations#e320c158 returning error if any.
func (c *Client) AccountGetAuthorizations(ctx context.Context, request *AccountGetAuthorizationsRequest) (*AccountAuthorizations, error) {
	var result AccountAuthorizations
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
