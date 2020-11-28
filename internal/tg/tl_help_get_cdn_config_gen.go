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

// HelpGetCdnConfigRequest represents TL type `help.getCdnConfig#52029342`.
type HelpGetCdnConfigRequest struct {
}

// HelpGetCdnConfigRequestTypeID is TL type id of HelpGetCdnConfigRequest.
const HelpGetCdnConfigRequestTypeID = 0x52029342

// Encode implements bin.Encoder.
func (g *HelpGetCdnConfigRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getCdnConfig#52029342 as nil")
	}
	b.PutID(HelpGetCdnConfigRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (g *HelpGetCdnConfigRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getCdnConfig#52029342 to nil")
	}
	if err := b.ConsumeID(HelpGetCdnConfigRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.getCdnConfig#52029342: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for HelpGetCdnConfigRequest.
var (
	_ bin.Encoder = &HelpGetCdnConfigRequest{}
	_ bin.Decoder = &HelpGetCdnConfigRequest{}
)
