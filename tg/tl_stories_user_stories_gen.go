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

// StoriesUserStories represents TL type `stories.userStories#37a6ff5f`.
//
// See https://core.telegram.org/constructor/stories.userStories for reference.
type StoriesUserStories struct {
	// Stories field of StoriesUserStories.
	Stories UserStories
	// Users field of StoriesUserStories.
	Users []UserClass
}

// StoriesUserStoriesTypeID is TL type id of StoriesUserStories.
const StoriesUserStoriesTypeID = 0x37a6ff5f

// Ensuring interfaces in compile-time for StoriesUserStories.
var (
	_ bin.Encoder     = &StoriesUserStories{}
	_ bin.Decoder     = &StoriesUserStories{}
	_ bin.BareEncoder = &StoriesUserStories{}
	_ bin.BareDecoder = &StoriesUserStories{}
)

func (u *StoriesUserStories) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Stories.Zero()) {
		return false
	}
	if !(u.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *StoriesUserStories) String() string {
	if u == nil {
		return "StoriesUserStories(nil)"
	}
	type Alias StoriesUserStories
	return fmt.Sprintf("StoriesUserStories%+v", Alias(*u))
}

// FillFrom fills StoriesUserStories from given interface.
func (u *StoriesUserStories) FillFrom(from interface {
	GetStories() (value UserStories)
	GetUsers() (value []UserClass)
}) {
	u.Stories = from.GetStories()
	u.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoriesUserStories) TypeID() uint32 {
	return StoriesUserStoriesTypeID
}

// TypeName returns name of type in TL schema.
func (*StoriesUserStories) TypeName() string {
	return "stories.userStories"
}

// TypeInfo returns info about TL type.
func (u *StoriesUserStories) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stories.userStories",
		ID:   StoriesUserStoriesTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Stories",
			SchemaName: "stories",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *StoriesUserStories) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode stories.userStories#37a6ff5f as nil")
	}
	b.PutID(StoriesUserStoriesTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *StoriesUserStories) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode stories.userStories#37a6ff5f as nil")
	}
	if err := u.Stories.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stories.userStories#37a6ff5f: field stories: %w", err)
	}
	b.PutVectorHeader(len(u.Users))
	for idx, v := range u.Users {
		if v == nil {
			return fmt.Errorf("unable to encode stories.userStories#37a6ff5f: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode stories.userStories#37a6ff5f: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *StoriesUserStories) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode stories.userStories#37a6ff5f to nil")
	}
	if err := b.ConsumeID(StoriesUserStoriesTypeID); err != nil {
		return fmt.Errorf("unable to decode stories.userStories#37a6ff5f: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *StoriesUserStories) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode stories.userStories#37a6ff5f to nil")
	}
	{
		if err := u.Stories.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stories.userStories#37a6ff5f: field stories: %w", err)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stories.userStories#37a6ff5f: field users: %w", err)
		}

		if headerLen > 0 {
			u.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode stories.userStories#37a6ff5f: field users: %w", err)
			}
			u.Users = append(u.Users, value)
		}
	}
	return nil
}

// GetStories returns value of Stories field.
func (u *StoriesUserStories) GetStories() (value UserStories) {
	if u == nil {
		return
	}
	return u.Stories
}

// GetUsers returns value of Users field.
func (u *StoriesUserStories) GetUsers() (value []UserClass) {
	if u == nil {
		return
	}
	return u.Users
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (u *StoriesUserStories) MapUsers() (value UserClassArray) {
	return UserClassArray(u.Users)
}
