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

// MessagesGetDialogsRequest represents TL type `messages.getDialogs#a0ee3b73`.
type MessagesGetDialogsRequest struct {
	// Flags field of MessagesGetDialogsRequest.
	Flags bin.Fields
	// ExcludePinned field of MessagesGetDialogsRequest.
	ExcludePinned bool
	// FolderID field of MessagesGetDialogsRequest.
	//
	// Use SetFolderID and GetFolderID helpers.
	FolderID int
	// OffsetDate field of MessagesGetDialogsRequest.
	OffsetDate int
	// OffsetID field of MessagesGetDialogsRequest.
	OffsetID int
	// OffsetPeer field of MessagesGetDialogsRequest.
	OffsetPeer InputPeerClass
	// Limit field of MessagesGetDialogsRequest.
	Limit int
	// Hash field of MessagesGetDialogsRequest.
	Hash int
}

// MessagesGetDialogsRequestTypeID is TL type id of MessagesGetDialogsRequest.
const MessagesGetDialogsRequestTypeID = 0xa0ee3b73

// Encode implements bin.Encoder.
func (g *MessagesGetDialogsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getDialogs#a0ee3b73 as nil")
	}
	b.PutID(MessagesGetDialogsRequestTypeID)
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getDialogs#a0ee3b73: field flags: %w", err)
	}
	if g.Flags.Has(1) {
		b.PutInt(g.FolderID)
	}
	b.PutInt(g.OffsetDate)
	b.PutInt(g.OffsetID)
	if g.OffsetPeer == nil {
		return fmt.Errorf("unable to encode messages.getDialogs#a0ee3b73: field offset_peer is nil")
	}
	if err := g.OffsetPeer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getDialogs#a0ee3b73: field offset_peer: %w", err)
	}
	b.PutInt(g.Limit)
	b.PutInt(g.Hash)
	return nil
}

// SetExcludePinned sets value of ExcludePinned conditional field.
func (g *MessagesGetDialogsRequest) SetExcludePinned(value bool) {
	if value {
		g.Flags.Set(0)
	} else {
		g.Flags.Unset(0)
	}
}

// SetFolderID sets value of FolderID conditional field.
func (g *MessagesGetDialogsRequest) SetFolderID(value int) {
	g.Flags.Set(1)
	g.FolderID = value
}

// GetFolderID returns value of FolderID conditional field and
// boolean which is true if field was set.
func (g *MessagesGetDialogsRequest) GetFolderID() (value int, ok bool) {
	if !g.Flags.Has(1) {
		return value, false
	}
	return g.FolderID, true
}

// Decode implements bin.Decoder.
func (g *MessagesGetDialogsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getDialogs#a0ee3b73 to nil")
	}
	if err := b.ConsumeID(MessagesGetDialogsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getDialogs#a0ee3b73: %w", err)
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.getDialogs#a0ee3b73: field flags: %w", err)
		}
	}
	g.ExcludePinned = g.Flags.Has(0)
	if g.Flags.Has(1) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getDialogs#a0ee3b73: field folder_id: %w", err)
		}
		g.FolderID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getDialogs#a0ee3b73: field offset_date: %w", err)
		}
		g.OffsetDate = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getDialogs#a0ee3b73: field offset_id: %w", err)
		}
		g.OffsetID = value
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getDialogs#a0ee3b73: field offset_peer: %w", err)
		}
		g.OffsetPeer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getDialogs#a0ee3b73: field limit: %w", err)
		}
		g.Limit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getDialogs#a0ee3b73: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetDialogsRequest.
var (
	_ bin.Encoder = &MessagesGetDialogsRequest{}
	_ bin.Decoder = &MessagesGetDialogsRequest{}
)

// MessagesGetDialogs invokes method messages.getDialogs#a0ee3b73 returning error if any.
func (c *Client) MessagesGetDialogs(ctx context.Context, request *MessagesGetDialogsRequest) (MessagesDialogsClass, error) {
	var result MessagesDialogsBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Dialogs, nil
}
