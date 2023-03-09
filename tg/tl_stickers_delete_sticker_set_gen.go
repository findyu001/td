// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// StickersDeleteStickerSetRequest represents TL type `stickers.deleteStickerSet#87704394`.
//
// See https://core.telegram.org/method/stickers.deleteStickerSet for reference.
type StickersDeleteStickerSetRequest struct {
	// Stickerset field of StickersDeleteStickerSetRequest.
	Stickerset InputStickerSetClass
}

// StickersDeleteStickerSetRequestTypeID is TL type id of StickersDeleteStickerSetRequest.
const StickersDeleteStickerSetRequestTypeID = 0x87704394

// Ensuring interfaces in compile-time for StickersDeleteStickerSetRequest.
var (
	_ bin.Encoder     = &StickersDeleteStickerSetRequest{}
	_ bin.Decoder     = &StickersDeleteStickerSetRequest{}
	_ bin.BareEncoder = &StickersDeleteStickerSetRequest{}
	_ bin.BareDecoder = &StickersDeleteStickerSetRequest{}
)

func (d *StickersDeleteStickerSetRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Stickerset == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *StickersDeleteStickerSetRequest) String() string {
	if d == nil {
		return "StickersDeleteStickerSetRequest(nil)"
	}
	type Alias StickersDeleteStickerSetRequest
	return fmt.Sprintf("StickersDeleteStickerSetRequest%+v", Alias(*d))
}

// FillFrom fills StickersDeleteStickerSetRequest from given interface.
func (d *StickersDeleteStickerSetRequest) FillFrom(from interface {
	GetStickerset() (value InputStickerSetClass)
}) {
	d.Stickerset = from.GetStickerset()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StickersDeleteStickerSetRequest) TypeID() uint32 {
	return StickersDeleteStickerSetRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*StickersDeleteStickerSetRequest) TypeName() string {
	return "stickers.deleteStickerSet"
}

// TypeInfo returns info about TL type.
func (d *StickersDeleteStickerSetRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stickers.deleteStickerSet",
		ID:   StickersDeleteStickerSetRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Stickerset",
			SchemaName: "stickerset",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *StickersDeleteStickerSetRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode stickers.deleteStickerSet#87704394 as nil")
	}
	b.PutID(StickersDeleteStickerSetRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *StickersDeleteStickerSetRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode stickers.deleteStickerSet#87704394 as nil")
	}
	if d.Stickerset == nil {
		return fmt.Errorf("unable to encode stickers.deleteStickerSet#87704394: field stickerset is nil")
	}
	if err := d.Stickerset.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stickers.deleteStickerSet#87704394: field stickerset: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *StickersDeleteStickerSetRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode stickers.deleteStickerSet#87704394 to nil")
	}
	if err := b.ConsumeID(StickersDeleteStickerSetRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stickers.deleteStickerSet#87704394: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *StickersDeleteStickerSetRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode stickers.deleteStickerSet#87704394 to nil")
	}
	{
		value, err := DecodeInputStickerSet(b)
		if err != nil {
			return fmt.Errorf("unable to decode stickers.deleteStickerSet#87704394: field stickerset: %w", err)
		}
		d.Stickerset = value
	}
	return nil
}

// GetStickerset returns value of Stickerset field.
func (d *StickersDeleteStickerSetRequest) GetStickerset() (value InputStickerSetClass) {
	if d == nil {
		return
	}
	return d.Stickerset
}

// StickersDeleteStickerSet invokes method stickers.deleteStickerSet#87704394 returning error if any.
//
// See https://core.telegram.org/method/stickers.deleteStickerSet for reference.
func (c *Client) StickersDeleteStickerSet(ctx context.Context, stickerset InputStickerSetClass) (bool, error) {
	var result BoolBox

	request := &StickersDeleteStickerSetRequest{
		Stickerset: stickerset,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
