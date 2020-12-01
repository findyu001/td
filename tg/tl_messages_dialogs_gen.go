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

// MessagesDialogs represents TL type `messages.dialogs#15ba6c40`.
type MessagesDialogs struct {
	// Dialogs field of MessagesDialogs.
	Dialogs []DialogClass
	// Messages field of MessagesDialogs.
	Messages []MessageClass
	// Chats field of MessagesDialogs.
	Chats []ChatClass
	// Users field of MessagesDialogs.
	Users []UserClass
}

// MessagesDialogsTypeID is TL type id of MessagesDialogs.
const MessagesDialogsTypeID = 0x15ba6c40

// Encode implements bin.Encoder.
func (d *MessagesDialogs) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode messages.dialogs#15ba6c40 as nil")
	}
	b.PutID(MessagesDialogsTypeID)
	b.PutVectorHeader(len(d.Dialogs))
	for idx, v := range d.Dialogs {
		if v == nil {
			return fmt.Errorf("unable to encode messages.dialogs#15ba6c40: field dialogs element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.dialogs#15ba6c40: field dialogs element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(d.Messages))
	for idx, v := range d.Messages {
		if v == nil {
			return fmt.Errorf("unable to encode messages.dialogs#15ba6c40: field messages element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.dialogs#15ba6c40: field messages element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(d.Chats))
	for idx, v := range d.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode messages.dialogs#15ba6c40: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.dialogs#15ba6c40: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(d.Users))
	for idx, v := range d.Users {
		if v == nil {
			return fmt.Errorf("unable to encode messages.dialogs#15ba6c40: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.dialogs#15ba6c40: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *MessagesDialogs) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode messages.dialogs#15ba6c40 to nil")
	}
	if err := b.ConsumeID(MessagesDialogsTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.dialogs#15ba6c40: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.dialogs#15ba6c40: field dialogs: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeDialog(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.dialogs#15ba6c40: field dialogs: %w", err)
			}
			d.Dialogs = append(d.Dialogs, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.dialogs#15ba6c40: field messages: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessage(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.dialogs#15ba6c40: field messages: %w", err)
			}
			d.Messages = append(d.Messages, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.dialogs#15ba6c40: field chats: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.dialogs#15ba6c40: field chats: %w", err)
			}
			d.Chats = append(d.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.dialogs#15ba6c40: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.dialogs#15ba6c40: field users: %w", err)
			}
			d.Users = append(d.Users, value)
		}
	}
	return nil
}

// construct implements constructor of MessagesDialogsClass.
func (d MessagesDialogs) construct() MessagesDialogsClass { return &d }

// Ensuring interfaces in compile-time for MessagesDialogs.
var (
	_ bin.Encoder = &MessagesDialogs{}
	_ bin.Decoder = &MessagesDialogs{}

	_ MessagesDialogsClass = &MessagesDialogs{}
)

// MessagesDialogsSlice represents TL type `messages.dialogsSlice#71e094f3`.
type MessagesDialogsSlice struct {
	// Count field of MessagesDialogsSlice.
	Count int
	// Dialogs field of MessagesDialogsSlice.
	Dialogs []DialogClass
	// Messages field of MessagesDialogsSlice.
	Messages []MessageClass
	// Chats field of MessagesDialogsSlice.
	Chats []ChatClass
	// Users field of MessagesDialogsSlice.
	Users []UserClass
}

// MessagesDialogsSliceTypeID is TL type id of MessagesDialogsSlice.
const MessagesDialogsSliceTypeID = 0x71e094f3

// Encode implements bin.Encoder.
func (d *MessagesDialogsSlice) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode messages.dialogsSlice#71e094f3 as nil")
	}
	b.PutID(MessagesDialogsSliceTypeID)
	b.PutInt(d.Count)
	b.PutVectorHeader(len(d.Dialogs))
	for idx, v := range d.Dialogs {
		if v == nil {
			return fmt.Errorf("unable to encode messages.dialogsSlice#71e094f3: field dialogs element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.dialogsSlice#71e094f3: field dialogs element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(d.Messages))
	for idx, v := range d.Messages {
		if v == nil {
			return fmt.Errorf("unable to encode messages.dialogsSlice#71e094f3: field messages element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.dialogsSlice#71e094f3: field messages element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(d.Chats))
	for idx, v := range d.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode messages.dialogsSlice#71e094f3: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.dialogsSlice#71e094f3: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(d.Users))
	for idx, v := range d.Users {
		if v == nil {
			return fmt.Errorf("unable to encode messages.dialogsSlice#71e094f3: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.dialogsSlice#71e094f3: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *MessagesDialogsSlice) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode messages.dialogsSlice#71e094f3 to nil")
	}
	if err := b.ConsumeID(MessagesDialogsSliceTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.dialogsSlice#71e094f3: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.dialogsSlice#71e094f3: field count: %w", err)
		}
		d.Count = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.dialogsSlice#71e094f3: field dialogs: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeDialog(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.dialogsSlice#71e094f3: field dialogs: %w", err)
			}
			d.Dialogs = append(d.Dialogs, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.dialogsSlice#71e094f3: field messages: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessage(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.dialogsSlice#71e094f3: field messages: %w", err)
			}
			d.Messages = append(d.Messages, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.dialogsSlice#71e094f3: field chats: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.dialogsSlice#71e094f3: field chats: %w", err)
			}
			d.Chats = append(d.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.dialogsSlice#71e094f3: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.dialogsSlice#71e094f3: field users: %w", err)
			}
			d.Users = append(d.Users, value)
		}
	}
	return nil
}

// construct implements constructor of MessagesDialogsClass.
func (d MessagesDialogsSlice) construct() MessagesDialogsClass { return &d }

// Ensuring interfaces in compile-time for MessagesDialogsSlice.
var (
	_ bin.Encoder = &MessagesDialogsSlice{}
	_ bin.Decoder = &MessagesDialogsSlice{}

	_ MessagesDialogsClass = &MessagesDialogsSlice{}
)

// MessagesDialogsNotModified represents TL type `messages.dialogsNotModified#f0e3e596`.
type MessagesDialogsNotModified struct {
	// Count field of MessagesDialogsNotModified.
	Count int
}

// MessagesDialogsNotModifiedTypeID is TL type id of MessagesDialogsNotModified.
const MessagesDialogsNotModifiedTypeID = 0xf0e3e596

// Encode implements bin.Encoder.
func (d *MessagesDialogsNotModified) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode messages.dialogsNotModified#f0e3e596 as nil")
	}
	b.PutID(MessagesDialogsNotModifiedTypeID)
	b.PutInt(d.Count)
	return nil
}

// Decode implements bin.Decoder.
func (d *MessagesDialogsNotModified) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode messages.dialogsNotModified#f0e3e596 to nil")
	}
	if err := b.ConsumeID(MessagesDialogsNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.dialogsNotModified#f0e3e596: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.dialogsNotModified#f0e3e596: field count: %w", err)
		}
		d.Count = value
	}
	return nil
}

// construct implements constructor of MessagesDialogsClass.
func (d MessagesDialogsNotModified) construct() MessagesDialogsClass { return &d }

// Ensuring interfaces in compile-time for MessagesDialogsNotModified.
var (
	_ bin.Encoder = &MessagesDialogsNotModified{}
	_ bin.Decoder = &MessagesDialogsNotModified{}

	_ MessagesDialogsClass = &MessagesDialogsNotModified{}
)

// MessagesDialogsClass represents messages.Dialogs generic type.
//
// Example:
//  g, err := DecodeMessagesDialogs(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *MessagesDialogs: // messages.dialogs#15ba6c40
//  case *MessagesDialogsSlice: // messages.dialogsSlice#71e094f3
//  case *MessagesDialogsNotModified: // messages.dialogsNotModified#f0e3e596
//  default: panic(v)
//  }
type MessagesDialogsClass interface {
	bin.Encoder
	bin.Decoder
	construct() MessagesDialogsClass
}

// DecodeMessagesDialogs implements binary de-serialization for MessagesDialogsClass.
func DecodeMessagesDialogs(buf *bin.Buffer) (MessagesDialogsClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessagesDialogsTypeID:
		// Decoding messages.dialogs#15ba6c40.
		v := MessagesDialogs{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesDialogsClass: %w", err)
		}
		return &v, nil
	case MessagesDialogsSliceTypeID:
		// Decoding messages.dialogsSlice#71e094f3.
		v := MessagesDialogsSlice{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesDialogsClass: %w", err)
		}
		return &v, nil
	case MessagesDialogsNotModifiedTypeID:
		// Decoding messages.dialogsNotModified#f0e3e596.
		v := MessagesDialogsNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesDialogsClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessagesDialogsClass: %w", bin.NewUnexpectedID(id))
	}
}

// MessagesDialogs boxes the MessagesDialogsClass providing a helper.
type MessagesDialogsBox struct {
	Dialogs MessagesDialogsClass
}

// Decode implements bin.Decoder for MessagesDialogsBox.
func (b *MessagesDialogsBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessagesDialogsBox to nil")
	}
	v, err := DecodeMessagesDialogs(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.Dialogs = v
	return nil
}

// Encode implements bin.Encode for MessagesDialogsBox.
func (b *MessagesDialogsBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.Dialogs == nil {
		return fmt.Errorf("unable to encode MessagesDialogsClass as nil")
	}
	return b.Dialogs.Encode(buf)
}
