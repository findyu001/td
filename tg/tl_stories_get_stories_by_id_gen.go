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

// StoriesGetStoriesByIDRequest represents TL type `stories.getStoriesByID#5774ca74`.
//
// See https://core.telegram.org/method/stories.getStoriesByID for reference.
type StoriesGetStoriesByIDRequest struct {
	// Peer field of StoriesGetStoriesByIDRequest.
	Peer InputPeerClass
	// ID field of StoriesGetStoriesByIDRequest.
	ID []int
}

// StoriesGetStoriesByIDRequestTypeID is TL type id of StoriesGetStoriesByIDRequest.
const StoriesGetStoriesByIDRequestTypeID = 0x5774ca74

// Ensuring interfaces in compile-time for StoriesGetStoriesByIDRequest.
var (
	_ bin.Encoder     = &StoriesGetStoriesByIDRequest{}
	_ bin.Decoder     = &StoriesGetStoriesByIDRequest{}
	_ bin.BareEncoder = &StoriesGetStoriesByIDRequest{}
	_ bin.BareDecoder = &StoriesGetStoriesByIDRequest{}
)

func (g *StoriesGetStoriesByIDRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.ID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *StoriesGetStoriesByIDRequest) String() string {
	if g == nil {
		return "StoriesGetStoriesByIDRequest(nil)"
	}
	type Alias StoriesGetStoriesByIDRequest
	return fmt.Sprintf("StoriesGetStoriesByIDRequest%+v", Alias(*g))
}

// FillFrom fills StoriesGetStoriesByIDRequest from given interface.
func (g *StoriesGetStoriesByIDRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetID() (value []int)
}) {
	g.Peer = from.GetPeer()
	g.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoriesGetStoriesByIDRequest) TypeID() uint32 {
	return StoriesGetStoriesByIDRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*StoriesGetStoriesByIDRequest) TypeName() string {
	return "stories.getStoriesByID"
}

// TypeInfo returns info about TL type.
func (g *StoriesGetStoriesByIDRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stories.getStoriesByID",
		ID:   StoriesGetStoriesByIDRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *StoriesGetStoriesByIDRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode stories.getStoriesByID#5774ca74 as nil")
	}
	b.PutID(StoriesGetStoriesByIDRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *StoriesGetStoriesByIDRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode stories.getStoriesByID#5774ca74 as nil")
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode stories.getStoriesByID#5774ca74: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stories.getStoriesByID#5774ca74: field peer: %w", err)
	}
	b.PutVectorHeader(len(g.ID))
	for _, v := range g.ID {
		b.PutInt(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *StoriesGetStoriesByIDRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode stories.getStoriesByID#5774ca74 to nil")
	}
	if err := b.ConsumeID(StoriesGetStoriesByIDRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stories.getStoriesByID#5774ca74: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *StoriesGetStoriesByIDRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode stories.getStoriesByID#5774ca74 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode stories.getStoriesByID#5774ca74: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stories.getStoriesByID#5774ca74: field id: %w", err)
		}

		if headerLen > 0 {
			g.ID = make([]int, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode stories.getStoriesByID#5774ca74: field id: %w", err)
			}
			g.ID = append(g.ID, value)
		}
	}
	return nil
}

// GetPeer returns value of Peer field.
func (g *StoriesGetStoriesByIDRequest) GetPeer() (value InputPeerClass) {
	if g == nil {
		return
	}
	return g.Peer
}

// GetID returns value of ID field.
func (g *StoriesGetStoriesByIDRequest) GetID() (value []int) {
	if g == nil {
		return
	}
	return g.ID
}

// StoriesGetStoriesByID invokes method stories.getStoriesByID#5774ca74 returning error if any.
//
// See https://core.telegram.org/method/stories.getStoriesByID for reference.
func (c *Client) StoriesGetStoriesByID(ctx context.Context, request *StoriesGetStoriesByIDRequest) (*StoriesStories, error) {
	var result StoriesStories

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
