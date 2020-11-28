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

// StickersChangeStickerPositionRequest represents TL type `stickers.changeStickerPosition#ffb6d4ca`.
type StickersChangeStickerPositionRequest struct {
	// Sticker field of StickersChangeStickerPositionRequest.
	Sticker InputDocumentClass
	// Position field of StickersChangeStickerPositionRequest.
	Position int
}

// StickersChangeStickerPositionRequestTypeID is TL type id of StickersChangeStickerPositionRequest.
const StickersChangeStickerPositionRequestTypeID = 0xffb6d4ca

// Encode implements bin.Encoder.
func (c *StickersChangeStickerPositionRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode stickers.changeStickerPosition#ffb6d4ca as nil")
	}
	b.PutID(StickersChangeStickerPositionRequestTypeID)
	if c.Sticker == nil {
		return fmt.Errorf("unable to encode stickers.changeStickerPosition#ffb6d4ca: field sticker is nil")
	}
	if err := c.Sticker.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stickers.changeStickerPosition#ffb6d4ca: field sticker: %w", err)
	}
	b.PutInt(c.Position)
	return nil
}

// Decode implements bin.Decoder.
func (c *StickersChangeStickerPositionRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode stickers.changeStickerPosition#ffb6d4ca to nil")
	}
	if err := b.ConsumeID(StickersChangeStickerPositionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stickers.changeStickerPosition#ffb6d4ca: %w", err)
	}
	{
		value, err := DecodeInputDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode stickers.changeStickerPosition#ffb6d4ca: field sticker: %w", err)
		}
		c.Sticker = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stickers.changeStickerPosition#ffb6d4ca: field position: %w", err)
		}
		c.Position = value
	}
	return nil
}

// Ensuring interfaces in compile-time for StickersChangeStickerPositionRequest.
var (
	_ bin.Encoder = &StickersChangeStickerPositionRequest{}
	_ bin.Decoder = &StickersChangeStickerPositionRequest{}
)
