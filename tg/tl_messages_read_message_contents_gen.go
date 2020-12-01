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

// MessagesReadMessageContentsRequest represents TL type `messages.readMessageContents#36a73f77`.
type MessagesReadMessageContentsRequest struct {
	// ID field of MessagesReadMessageContentsRequest.
	ID []int
}

// MessagesReadMessageContentsRequestTypeID is TL type id of MessagesReadMessageContentsRequest.
const MessagesReadMessageContentsRequestTypeID = 0x36a73f77

// Encode implements bin.Encoder.
func (r *MessagesReadMessageContentsRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.readMessageContents#36a73f77 as nil")
	}
	b.PutID(MessagesReadMessageContentsRequestTypeID)
	b.PutVectorHeader(len(r.ID))
	for _, v := range r.ID {
		b.PutInt(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *MessagesReadMessageContentsRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.readMessageContents#36a73f77 to nil")
	}
	if err := b.ConsumeID(MessagesReadMessageContentsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.readMessageContents#36a73f77: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.readMessageContents#36a73f77: field id: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode messages.readMessageContents#36a73f77: field id: %w", err)
			}
			r.ID = append(r.ID, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesReadMessageContentsRequest.
var (
	_ bin.Encoder = &MessagesReadMessageContentsRequest{}
	_ bin.Decoder = &MessagesReadMessageContentsRequest{}
)

// MessagesReadMessageContents invokes method messages.readMessageContents#36a73f77 returning error if any.
func (c *Client) MessagesReadMessageContents(ctx context.Context, request *MessagesReadMessageContentsRequest) (*MessagesAffectedMessages, error) {
	var result MessagesAffectedMessages
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
