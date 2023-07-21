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

// StoryViews represents TL type `storyViews#d36760cf`.
//
// See https://core.telegram.org/constructor/storyViews for reference.
type StoryViews struct {
	// Flags field of StoryViews.
	Flags bin.Fields
	// ViewsCount field of StoryViews.
	ViewsCount int
	// RecentViewers field of StoryViews.
	//
	// Use SetRecentViewers and GetRecentViewers helpers.
	RecentViewers []int64
}

// StoryViewsTypeID is TL type id of StoryViews.
const StoryViewsTypeID = 0xd36760cf

// Ensuring interfaces in compile-time for StoryViews.
var (
	_ bin.Encoder     = &StoryViews{}
	_ bin.Decoder     = &StoryViews{}
	_ bin.BareEncoder = &StoryViews{}
	_ bin.BareDecoder = &StoryViews{}
)

func (s *StoryViews) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.ViewsCount == 0) {
		return false
	}
	if !(s.RecentViewers == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StoryViews) String() string {
	if s == nil {
		return "StoryViews(nil)"
	}
	type Alias StoryViews
	return fmt.Sprintf("StoryViews%+v", Alias(*s))
}

// FillFrom fills StoryViews from given interface.
func (s *StoryViews) FillFrom(from interface {
	GetViewsCount() (value int)
	GetRecentViewers() (value []int64, ok bool)
}) {
	s.ViewsCount = from.GetViewsCount()
	if val, ok := from.GetRecentViewers(); ok {
		s.RecentViewers = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoryViews) TypeID() uint32 {
	return StoryViewsTypeID
}

// TypeName returns name of type in TL schema.
func (*StoryViews) TypeName() string {
	return "storyViews"
}

// TypeInfo returns info about TL type.
func (s *StoryViews) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "storyViews",
		ID:   StoryViewsTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ViewsCount",
			SchemaName: "views_count",
		},
		{
			Name:       "RecentViewers",
			SchemaName: "recent_viewers",
			Null:       !s.Flags.Has(0),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *StoryViews) SetFlags() {
	if !(s.RecentViewers == nil) {
		s.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (s *StoryViews) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyViews#d36760cf as nil")
	}
	b.PutID(StoryViewsTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StoryViews) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyViews#d36760cf as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode storyViews#d36760cf: field flags: %w", err)
	}
	b.PutInt(s.ViewsCount)
	if s.Flags.Has(0) {
		b.PutVectorHeader(len(s.RecentViewers))
		for _, v := range s.RecentViewers {
			b.PutLong(v)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StoryViews) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyViews#d36760cf to nil")
	}
	if err := b.ConsumeID(StoryViewsTypeID); err != nil {
		return fmt.Errorf("unable to decode storyViews#d36760cf: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StoryViews) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyViews#d36760cf to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode storyViews#d36760cf: field flags: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode storyViews#d36760cf: field views_count: %w", err)
		}
		s.ViewsCount = value
	}
	if s.Flags.Has(0) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode storyViews#d36760cf: field recent_viewers: %w", err)
		}

		if headerLen > 0 {
			s.RecentViewers = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode storyViews#d36760cf: field recent_viewers: %w", err)
			}
			s.RecentViewers = append(s.RecentViewers, value)
		}
	}
	return nil
}

// GetViewsCount returns value of ViewsCount field.
func (s *StoryViews) GetViewsCount() (value int) {
	if s == nil {
		return
	}
	return s.ViewsCount
}

// SetRecentViewers sets value of RecentViewers conditional field.
func (s *StoryViews) SetRecentViewers(value []int64) {
	s.Flags.Set(0)
	s.RecentViewers = value
}

// GetRecentViewers returns value of RecentViewers conditional field and
// boolean which is true if field was set.
func (s *StoryViews) GetRecentViewers() (value []int64, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.RecentViewers, true
}
