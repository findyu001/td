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

// AccountGetNotifySettingsRequest represents TL type `account.getNotifySettings#12b3ad31`.
type AccountGetNotifySettingsRequest struct {
	// Peer field of AccountGetNotifySettingsRequest.
	Peer InputNotifyPeerClass
}

// AccountGetNotifySettingsRequestTypeID is TL type id of AccountGetNotifySettingsRequest.
const AccountGetNotifySettingsRequestTypeID = 0x12b3ad31

// Encode implements bin.Encoder.
func (g *AccountGetNotifySettingsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getNotifySettings#12b3ad31 as nil")
	}
	b.PutID(AccountGetNotifySettingsRequestTypeID)
	if g.Peer == nil {
		return fmt.Errorf("unable to encode account.getNotifySettings#12b3ad31: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.getNotifySettings#12b3ad31: field peer: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *AccountGetNotifySettingsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getNotifySettings#12b3ad31 to nil")
	}
	if err := b.ConsumeID(AccountGetNotifySettingsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getNotifySettings#12b3ad31: %w", err)
	}
	{
		value, err := DecodeInputNotifyPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.getNotifySettings#12b3ad31: field peer: %w", err)
		}
		g.Peer = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountGetNotifySettingsRequest.
var (
	_ bin.Encoder = &AccountGetNotifySettingsRequest{}
	_ bin.Decoder = &AccountGetNotifySettingsRequest{}
)
