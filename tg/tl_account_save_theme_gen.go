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

// AccountSaveThemeRequest represents TL type `account.saveTheme#f257106c`.
type AccountSaveThemeRequest struct {
	// Theme field of AccountSaveThemeRequest.
	Theme InputThemeClass
	// Unsave field of AccountSaveThemeRequest.
	Unsave bool
}

// AccountSaveThemeRequestTypeID is TL type id of AccountSaveThemeRequest.
const AccountSaveThemeRequestTypeID = 0xf257106c

// Encode implements bin.Encoder.
func (s *AccountSaveThemeRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.saveTheme#f257106c as nil")
	}
	b.PutID(AccountSaveThemeRequestTypeID)
	if s.Theme == nil {
		return fmt.Errorf("unable to encode account.saveTheme#f257106c: field theme is nil")
	}
	if err := s.Theme.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.saveTheme#f257106c: field theme: %w", err)
	}
	b.PutBool(s.Unsave)
	return nil
}

// Decode implements bin.Decoder.
func (s *AccountSaveThemeRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.saveTheme#f257106c to nil")
	}
	if err := b.ConsumeID(AccountSaveThemeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.saveTheme#f257106c: %w", err)
	}
	{
		value, err := DecodeInputTheme(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.saveTheme#f257106c: field theme: %w", err)
		}
		s.Theme = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode account.saveTheme#f257106c: field unsave: %w", err)
		}
		s.Unsave = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountSaveThemeRequest.
var (
	_ bin.Encoder = &AccountSaveThemeRequest{}
	_ bin.Decoder = &AccountSaveThemeRequest{}
)

// AccountSaveTheme invokes method account.saveTheme#f257106c returning error if any.
func (c *Client) AccountSaveTheme(ctx context.Context, request *AccountSaveThemeRequest) (BoolClass, error) {
	var result BoolBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Bool, nil
}
