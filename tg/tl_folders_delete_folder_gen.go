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

// FoldersDeleteFolderRequest represents TL type `folders.deleteFolder#1c295881`.
type FoldersDeleteFolderRequest struct {
	// FolderID field of FoldersDeleteFolderRequest.
	FolderID int
}

// FoldersDeleteFolderRequestTypeID is TL type id of FoldersDeleteFolderRequest.
const FoldersDeleteFolderRequestTypeID = 0x1c295881

// Encode implements bin.Encoder.
func (d *FoldersDeleteFolderRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode folders.deleteFolder#1c295881 as nil")
	}
	b.PutID(FoldersDeleteFolderRequestTypeID)
	b.PutInt(d.FolderID)
	return nil
}

// Decode implements bin.Decoder.
func (d *FoldersDeleteFolderRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode folders.deleteFolder#1c295881 to nil")
	}
	if err := b.ConsumeID(FoldersDeleteFolderRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode folders.deleteFolder#1c295881: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode folders.deleteFolder#1c295881: field folder_id: %w", err)
		}
		d.FolderID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for FoldersDeleteFolderRequest.
var (
	_ bin.Encoder = &FoldersDeleteFolderRequest{}
	_ bin.Decoder = &FoldersDeleteFolderRequest{}
)

// FoldersDeleteFolder invokes method folders.deleteFolder#1c295881 returning error if any.
func (c *Client) FoldersDeleteFolder(ctx context.Context, request *FoldersDeleteFolderRequest) (UpdatesClass, error) {
	var result UpdatesBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
