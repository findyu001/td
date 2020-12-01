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

// MessagesStickerSetInstallResultSuccess represents TL type `messages.stickerSetInstallResultSuccess#38641628`.
type MessagesStickerSetInstallResultSuccess struct {
}

// MessagesStickerSetInstallResultSuccessTypeID is TL type id of MessagesStickerSetInstallResultSuccess.
const MessagesStickerSetInstallResultSuccessTypeID = 0x38641628

// Encode implements bin.Encoder.
func (s *MessagesStickerSetInstallResultSuccess) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.stickerSetInstallResultSuccess#38641628 as nil")
	}
	b.PutID(MessagesStickerSetInstallResultSuccessTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesStickerSetInstallResultSuccess) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.stickerSetInstallResultSuccess#38641628 to nil")
	}
	if err := b.ConsumeID(MessagesStickerSetInstallResultSuccessTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.stickerSetInstallResultSuccess#38641628: %w", err)
	}
	return nil
}

// construct implements constructor of MessagesStickerSetInstallResultClass.
func (s MessagesStickerSetInstallResultSuccess) construct() MessagesStickerSetInstallResultClass {
	return &s
}

// Ensuring interfaces in compile-time for MessagesStickerSetInstallResultSuccess.
var (
	_ bin.Encoder = &MessagesStickerSetInstallResultSuccess{}
	_ bin.Decoder = &MessagesStickerSetInstallResultSuccess{}

	_ MessagesStickerSetInstallResultClass = &MessagesStickerSetInstallResultSuccess{}
)

// MessagesStickerSetInstallResultArchive represents TL type `messages.stickerSetInstallResultArchive#35e410a8`.
type MessagesStickerSetInstallResultArchive struct {
	// Sets field of MessagesStickerSetInstallResultArchive.
	Sets []StickerSetCoveredClass
}

// MessagesStickerSetInstallResultArchiveTypeID is TL type id of MessagesStickerSetInstallResultArchive.
const MessagesStickerSetInstallResultArchiveTypeID = 0x35e410a8

// Encode implements bin.Encoder.
func (s *MessagesStickerSetInstallResultArchive) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.stickerSetInstallResultArchive#35e410a8 as nil")
	}
	b.PutID(MessagesStickerSetInstallResultArchiveTypeID)
	b.PutVectorHeader(len(s.Sets))
	for idx, v := range s.Sets {
		if v == nil {
			return fmt.Errorf("unable to encode messages.stickerSetInstallResultArchive#35e410a8: field sets element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.stickerSetInstallResultArchive#35e410a8: field sets element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesStickerSetInstallResultArchive) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.stickerSetInstallResultArchive#35e410a8 to nil")
	}
	if err := b.ConsumeID(MessagesStickerSetInstallResultArchiveTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.stickerSetInstallResultArchive#35e410a8: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.stickerSetInstallResultArchive#35e410a8: field sets: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeStickerSetCovered(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.stickerSetInstallResultArchive#35e410a8: field sets: %w", err)
			}
			s.Sets = append(s.Sets, value)
		}
	}
	return nil
}

// construct implements constructor of MessagesStickerSetInstallResultClass.
func (s MessagesStickerSetInstallResultArchive) construct() MessagesStickerSetInstallResultClass {
	return &s
}

// Ensuring interfaces in compile-time for MessagesStickerSetInstallResultArchive.
var (
	_ bin.Encoder = &MessagesStickerSetInstallResultArchive{}
	_ bin.Decoder = &MessagesStickerSetInstallResultArchive{}

	_ MessagesStickerSetInstallResultClass = &MessagesStickerSetInstallResultArchive{}
)

// MessagesStickerSetInstallResultClass represents messages.StickerSetInstallResult generic type.
//
// Example:
//  g, err := DecodeMessagesStickerSetInstallResult(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *MessagesStickerSetInstallResultSuccess: // messages.stickerSetInstallResultSuccess#38641628
//  case *MessagesStickerSetInstallResultArchive: // messages.stickerSetInstallResultArchive#35e410a8
//  default: panic(v)
//  }
type MessagesStickerSetInstallResultClass interface {
	bin.Encoder
	bin.Decoder
	construct() MessagesStickerSetInstallResultClass
}

// DecodeMessagesStickerSetInstallResult implements binary de-serialization for MessagesStickerSetInstallResultClass.
func DecodeMessagesStickerSetInstallResult(buf *bin.Buffer) (MessagesStickerSetInstallResultClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessagesStickerSetInstallResultSuccessTypeID:
		// Decoding messages.stickerSetInstallResultSuccess#38641628.
		v := MessagesStickerSetInstallResultSuccess{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesStickerSetInstallResultClass: %w", err)
		}
		return &v, nil
	case MessagesStickerSetInstallResultArchiveTypeID:
		// Decoding messages.stickerSetInstallResultArchive#35e410a8.
		v := MessagesStickerSetInstallResultArchive{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesStickerSetInstallResultClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessagesStickerSetInstallResultClass: %w", bin.NewUnexpectedID(id))
	}
}

// MessagesStickerSetInstallResult boxes the MessagesStickerSetInstallResultClass providing a helper.
type MessagesStickerSetInstallResultBox struct {
	StickerSetInstallResult MessagesStickerSetInstallResultClass
}

// Decode implements bin.Decoder for MessagesStickerSetInstallResultBox.
func (b *MessagesStickerSetInstallResultBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessagesStickerSetInstallResultBox to nil")
	}
	v, err := DecodeMessagesStickerSetInstallResult(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.StickerSetInstallResult = v
	return nil
}

// Encode implements bin.Encode for MessagesStickerSetInstallResultBox.
func (b *MessagesStickerSetInstallResultBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.StickerSetInstallResult == nil {
		return fmt.Errorf("unable to encode MessagesStickerSetInstallResultClass as nil")
	}
	return b.StickerSetInstallResult.Encode(buf)
}
