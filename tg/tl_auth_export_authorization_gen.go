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

// AuthExportAuthorizationRequest represents TL type `auth.exportAuthorization#e5bfffcd`.
type AuthExportAuthorizationRequest struct {
	// DCID field of AuthExportAuthorizationRequest.
	DCID int
}

// AuthExportAuthorizationRequestTypeID is TL type id of AuthExportAuthorizationRequest.
const AuthExportAuthorizationRequestTypeID = 0xe5bfffcd

// Encode implements bin.Encoder.
func (e *AuthExportAuthorizationRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode auth.exportAuthorization#e5bfffcd as nil")
	}
	b.PutID(AuthExportAuthorizationRequestTypeID)
	b.PutInt(e.DCID)
	return nil
}

// Decode implements bin.Decoder.
func (e *AuthExportAuthorizationRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode auth.exportAuthorization#e5bfffcd to nil")
	}
	if err := b.ConsumeID(AuthExportAuthorizationRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.exportAuthorization#e5bfffcd: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode auth.exportAuthorization#e5bfffcd: field dc_id: %w", err)
		}
		e.DCID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AuthExportAuthorizationRequest.
var (
	_ bin.Encoder = &AuthExportAuthorizationRequest{}
	_ bin.Decoder = &AuthExportAuthorizationRequest{}
)

// AuthExportAuthorization invokes method auth.exportAuthorization#e5bfffcd returning error if any.
func (c *Client) AuthExportAuthorization(ctx context.Context, request *AuthExportAuthorizationRequest) (*AuthExportedAuthorization, error) {
	var result AuthExportedAuthorization
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
