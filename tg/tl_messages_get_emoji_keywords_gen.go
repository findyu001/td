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

// MessagesGetEmojiKeywordsRequest represents TL type `messages.getEmojiKeywords#35a0e062`.
type MessagesGetEmojiKeywordsRequest struct {
	// LangCode field of MessagesGetEmojiKeywordsRequest.
	LangCode string
}

// MessagesGetEmojiKeywordsRequestTypeID is TL type id of MessagesGetEmojiKeywordsRequest.
const MessagesGetEmojiKeywordsRequestTypeID = 0x35a0e062

// Encode implements bin.Encoder.
func (g *MessagesGetEmojiKeywordsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getEmojiKeywords#35a0e062 as nil")
	}
	b.PutID(MessagesGetEmojiKeywordsRequestTypeID)
	b.PutString(g.LangCode)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetEmojiKeywordsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getEmojiKeywords#35a0e062 to nil")
	}
	if err := b.ConsumeID(MessagesGetEmojiKeywordsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getEmojiKeywords#35a0e062: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getEmojiKeywords#35a0e062: field lang_code: %w", err)
		}
		g.LangCode = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetEmojiKeywordsRequest.
var (
	_ bin.Encoder = &MessagesGetEmojiKeywordsRequest{}
	_ bin.Decoder = &MessagesGetEmojiKeywordsRequest{}
)

// MessagesGetEmojiKeywords invokes method messages.getEmojiKeywords#35a0e062 returning error if any.
func (c *Client) MessagesGetEmojiKeywords(ctx context.Context, request *MessagesGetEmojiKeywordsRequest) (*EmojiKeywordsDifference, error) {
	var result EmojiKeywordsDifference
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
