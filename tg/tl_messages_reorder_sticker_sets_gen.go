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

// MessagesReorderStickerSetsRequest represents TL type `messages.reorderStickerSets#78337739`.
type MessagesReorderStickerSetsRequest struct {
	// Flags field of MessagesReorderStickerSetsRequest.
	Flags bin.Fields
	// Masks field of MessagesReorderStickerSetsRequest.
	Masks bool
	// Order field of MessagesReorderStickerSetsRequest.
	Order []int64
}

// MessagesReorderStickerSetsRequestTypeID is TL type id of MessagesReorderStickerSetsRequest.
const MessagesReorderStickerSetsRequestTypeID = 0x78337739

// Encode implements bin.Encoder.
func (r *MessagesReorderStickerSetsRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.reorderStickerSets#78337739 as nil")
	}
	b.PutID(MessagesReorderStickerSetsRequestTypeID)
	if err := r.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.reorderStickerSets#78337739: field flags: %w", err)
	}
	b.PutVectorHeader(len(r.Order))
	for _, v := range r.Order {
		b.PutLong(v)
	}
	return nil
}

// SetMasks sets value of Masks conditional field.
func (r *MessagesReorderStickerSetsRequest) SetMasks(value bool) {
	if value {
		r.Flags.Set(0)
	} else {
		r.Flags.Unset(0)
	}
}

// Decode implements bin.Decoder.
func (r *MessagesReorderStickerSetsRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.reorderStickerSets#78337739 to nil")
	}
	if err := b.ConsumeID(MessagesReorderStickerSetsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.reorderStickerSets#78337739: %w", err)
	}
	{
		if err := r.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.reorderStickerSets#78337739: field flags: %w", err)
		}
	}
	r.Masks = r.Flags.Has(0)
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.reorderStickerSets#78337739: field order: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode messages.reorderStickerSets#78337739: field order: %w", err)
			}
			r.Order = append(r.Order, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesReorderStickerSetsRequest.
var (
	_ bin.Encoder = &MessagesReorderStickerSetsRequest{}
	_ bin.Decoder = &MessagesReorderStickerSetsRequest{}
)

// MessagesReorderStickerSets invokes method messages.reorderStickerSets#78337739 returning error if any.
func (c *Client) MessagesReorderStickerSets(ctx context.Context, request *MessagesReorderStickerSetsRequest) (BoolClass, error) {
	var result BoolBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Bool, nil
}
