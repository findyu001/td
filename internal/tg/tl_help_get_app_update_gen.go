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

// HelpGetAppUpdateRequest represents TL type `help.getAppUpdate#522d5a7d`.
type HelpGetAppUpdateRequest struct {
	// Source field of HelpGetAppUpdateRequest.
	Source string
}

// HelpGetAppUpdateRequestTypeID is TL type id of HelpGetAppUpdateRequest.
const HelpGetAppUpdateRequestTypeID = 0x522d5a7d

// Encode implements bin.Encoder.
func (g *HelpGetAppUpdateRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getAppUpdate#522d5a7d as nil")
	}
	b.PutID(HelpGetAppUpdateRequestTypeID)
	b.PutString(g.Source)
	return nil
}

// Decode implements bin.Decoder.
func (g *HelpGetAppUpdateRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getAppUpdate#522d5a7d to nil")
	}
	if err := b.ConsumeID(HelpGetAppUpdateRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.getAppUpdate#522d5a7d: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.getAppUpdate#522d5a7d: field source: %w", err)
		}
		g.Source = value
	}
	return nil
}

// Ensuring interfaces in compile-time for HelpGetAppUpdateRequest.
var (
	_ bin.Encoder = &HelpGetAppUpdateRequest{}
	_ bin.Decoder = &HelpGetAppUpdateRequest{}
)
