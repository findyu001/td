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

// AuthRequestPasswordRecoveryRequest represents TL type `auth.requestPasswordRecovery#d897bc66`.
type AuthRequestPasswordRecoveryRequest struct {
}

// AuthRequestPasswordRecoveryRequestTypeID is TL type id of AuthRequestPasswordRecoveryRequest.
const AuthRequestPasswordRecoveryRequestTypeID = 0xd897bc66

// Encode implements bin.Encoder.
func (r *AuthRequestPasswordRecoveryRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode auth.requestPasswordRecovery#d897bc66 as nil")
	}
	b.PutID(AuthRequestPasswordRecoveryRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (r *AuthRequestPasswordRecoveryRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode auth.requestPasswordRecovery#d897bc66 to nil")
	}
	if err := b.ConsumeID(AuthRequestPasswordRecoveryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.requestPasswordRecovery#d897bc66: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for AuthRequestPasswordRecoveryRequest.
var (
	_ bin.Encoder = &AuthRequestPasswordRecoveryRequest{}
	_ bin.Decoder = &AuthRequestPasswordRecoveryRequest{}
)
