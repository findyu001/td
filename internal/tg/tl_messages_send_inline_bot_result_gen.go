// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"

	"github.com/ernado/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// MessagesSendInlineBotResultRequest represents TL type `messages.sendInlineBotResult#220815b0`.
type MessagesSendInlineBotResultRequest struct {
	// Flags field of MessagesSendInlineBotResultRequest.
	Flags bin.Fields
	// Silent field of MessagesSendInlineBotResultRequest.
	Silent bool
	// Background field of MessagesSendInlineBotResultRequest.
	Background bool
	// ClearDraft field of MessagesSendInlineBotResultRequest.
	ClearDraft bool
	// HideVia field of MessagesSendInlineBotResultRequest.
	HideVia bool
	// Peer field of MessagesSendInlineBotResultRequest.
	Peer InputPeerClass
	// ReplyToMsgID field of MessagesSendInlineBotResultRequest.
	//
	// Use SetReplyToMsgID and GetReplyToMsgID helpers.
	ReplyToMsgID int
	// RandomID field of MessagesSendInlineBotResultRequest.
	RandomID int64
	// QueryID field of MessagesSendInlineBotResultRequest.
	QueryID int64
	// ID field of MessagesSendInlineBotResultRequest.
	ID string
	// ScheduleDate field of MessagesSendInlineBotResultRequest.
	//
	// Use SetScheduleDate and GetScheduleDate helpers.
	ScheduleDate int
}

// MessagesSendInlineBotResultRequestTypeID is TL type id of MessagesSendInlineBotResultRequest.
const MessagesSendInlineBotResultRequestTypeID = 0x220815b0

// Encode implements bin.Encoder.
func (s *MessagesSendInlineBotResultRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sendInlineBotResult#220815b0 as nil")
	}
	b.PutID(MessagesSendInlineBotResultRequestTypeID)
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendInlineBotResult#220815b0: field flags: %w", err)
	}
	if s.Peer == nil {
		return fmt.Errorf("unable to encode messages.sendInlineBotResult#220815b0: field peer is nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendInlineBotResult#220815b0: field peer: %w", err)
	}
	if s.Flags.Has(0) {
		b.PutInt(s.ReplyToMsgID)
	}
	b.PutLong(s.RandomID)
	b.PutLong(s.QueryID)
	b.PutString(s.ID)
	if s.Flags.Has(10) {
		b.PutInt(s.ScheduleDate)
	}
	return nil
}

// SetSilent sets value of Silent conditional field.
func (s *MessagesSendInlineBotResultRequest) SetSilent(value bool) {
	if value {
		s.Flags.Set(5)
	} else {
		s.Flags.Unset(5)
	}
}

// SetBackground sets value of Background conditional field.
func (s *MessagesSendInlineBotResultRequest) SetBackground(value bool) {
	if value {
		s.Flags.Set(6)
	} else {
		s.Flags.Unset(6)
	}
}

// SetClearDraft sets value of ClearDraft conditional field.
func (s *MessagesSendInlineBotResultRequest) SetClearDraft(value bool) {
	if value {
		s.Flags.Set(7)
	} else {
		s.Flags.Unset(7)
	}
}

// SetHideVia sets value of HideVia conditional field.
func (s *MessagesSendInlineBotResultRequest) SetHideVia(value bool) {
	if value {
		s.Flags.Set(11)
	} else {
		s.Flags.Unset(11)
	}
}

// SetReplyToMsgID sets value of ReplyToMsgID conditional field.
func (s *MessagesSendInlineBotResultRequest) SetReplyToMsgID(value int) {
	s.Flags.Set(0)
	s.ReplyToMsgID = value
}

// GetReplyToMsgID returns value of ReplyToMsgID conditional field and
// boolean which is true if field was set.
func (s *MessagesSendInlineBotResultRequest) GetReplyToMsgID() (value int, ok bool) {
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.ReplyToMsgID, true
}

// SetScheduleDate sets value of ScheduleDate conditional field.
func (s *MessagesSendInlineBotResultRequest) SetScheduleDate(value int) {
	s.Flags.Set(10)
	s.ScheduleDate = value
}

// GetScheduleDate returns value of ScheduleDate conditional field and
// boolean which is true if field was set.
func (s *MessagesSendInlineBotResultRequest) GetScheduleDate() (value int, ok bool) {
	if !s.Flags.Has(10) {
		return value, false
	}
	return s.ScheduleDate, true
}

// Decode implements bin.Decoder.
func (s *MessagesSendInlineBotResultRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sendInlineBotResult#220815b0 to nil")
	}
	if err := b.ConsumeID(MessagesSendInlineBotResultRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.sendInlineBotResult#220815b0: %w", err)
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.sendInlineBotResult#220815b0: field flags: %w", err)
		}
	}
	s.Silent = s.Flags.Has(5)
	s.Background = s.Flags.Has(6)
	s.ClearDraft = s.Flags.Has(7)
	s.HideVia = s.Flags.Has(11)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendInlineBotResult#220815b0: field peer: %w", err)
		}
		s.Peer = value
	}
	if s.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendInlineBotResult#220815b0: field reply_to_msg_id: %w", err)
		}
		s.ReplyToMsgID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendInlineBotResult#220815b0: field random_id: %w", err)
		}
		s.RandomID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendInlineBotResult#220815b0: field query_id: %w", err)
		}
		s.QueryID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendInlineBotResult#220815b0: field id: %w", err)
		}
		s.ID = value
	}
	if s.Flags.Has(10) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendInlineBotResult#220815b0: field schedule_date: %w", err)
		}
		s.ScheduleDate = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesSendInlineBotResultRequest.
var (
	_ bin.Encoder = &MessagesSendInlineBotResultRequest{}
	_ bin.Decoder = &MessagesSendInlineBotResultRequest{}
)
