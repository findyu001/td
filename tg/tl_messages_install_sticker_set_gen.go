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

// MessagesInstallStickerSetRequest represents TL type `messages.installStickerSet#c78fe460`.
type MessagesInstallStickerSetRequest struct {
	// Stickerset field of MessagesInstallStickerSetRequest.
	Stickerset InputStickerSetClass
	// Archived field of MessagesInstallStickerSetRequest.
	Archived bool
}

// MessagesInstallStickerSetRequestTypeID is TL type id of MessagesInstallStickerSetRequest.
const MessagesInstallStickerSetRequestTypeID = 0xc78fe460

// Encode implements bin.Encoder.
func (i *MessagesInstallStickerSetRequest) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode messages.installStickerSet#c78fe460 as nil")
	}
	b.PutID(MessagesInstallStickerSetRequestTypeID)
	if i.Stickerset == nil {
		return fmt.Errorf("unable to encode messages.installStickerSet#c78fe460: field stickerset is nil")
	}
	if err := i.Stickerset.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.installStickerSet#c78fe460: field stickerset: %w", err)
	}
	b.PutBool(i.Archived)
	return nil
}

// Decode implements bin.Decoder.
func (i *MessagesInstallStickerSetRequest) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode messages.installStickerSet#c78fe460 to nil")
	}
	if err := b.ConsumeID(MessagesInstallStickerSetRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.installStickerSet#c78fe460: %w", err)
	}
	{
		value, err := DecodeInputStickerSet(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.installStickerSet#c78fe460: field stickerset: %w", err)
		}
		i.Stickerset = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode messages.installStickerSet#c78fe460: field archived: %w", err)
		}
		i.Archived = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesInstallStickerSetRequest.
var (
	_ bin.Encoder = &MessagesInstallStickerSetRequest{}
	_ bin.Decoder = &MessagesInstallStickerSetRequest{}
)

// MessagesInstallStickerSet invokes method messages.installStickerSet#c78fe460 returning error if any.
func (c *Client) MessagesInstallStickerSet(ctx context.Context, request *MessagesInstallStickerSetRequest) (MessagesStickerSetInstallResultClass, error) {
	var result MessagesStickerSetInstallResultBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.StickerSetInstallResult, nil
}
