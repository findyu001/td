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

// AccountSetContactSignUpNotificationRequest represents TL type `account.setContactSignUpNotification#cff43f61`.
type AccountSetContactSignUpNotificationRequest struct {
	// Silent field of AccountSetContactSignUpNotificationRequest.
	Silent bool
}

// AccountSetContactSignUpNotificationRequestTypeID is TL type id of AccountSetContactSignUpNotificationRequest.
const AccountSetContactSignUpNotificationRequestTypeID = 0xcff43f61

// Encode implements bin.Encoder.
func (s *AccountSetContactSignUpNotificationRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.setContactSignUpNotification#cff43f61 as nil")
	}
	b.PutID(AccountSetContactSignUpNotificationRequestTypeID)
	b.PutBool(s.Silent)
	return nil
}

// Decode implements bin.Decoder.
func (s *AccountSetContactSignUpNotificationRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.setContactSignUpNotification#cff43f61 to nil")
	}
	if err := b.ConsumeID(AccountSetContactSignUpNotificationRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.setContactSignUpNotification#cff43f61: %w", err)
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode account.setContactSignUpNotification#cff43f61: field silent: %w", err)
		}
		s.Silent = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountSetContactSignUpNotificationRequest.
var (
	_ bin.Encoder = &AccountSetContactSignUpNotificationRequest{}
	_ bin.Decoder = &AccountSetContactSignUpNotificationRequest{}
)

// AccountSetContactSignUpNotification invokes method account.setContactSignUpNotification#cff43f61 returning error if any.
func (c *Client) AccountSetContactSignUpNotification(ctx context.Context, request *AccountSetContactSignUpNotificationRequest) (BoolClass, error) {
	var result BoolBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Bool, nil
}
