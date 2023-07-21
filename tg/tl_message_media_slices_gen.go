//go:build !no_gotd_slices
// +build !no_gotd_slices

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

// MessageMediaClassArray is adapter for slice of MessageMediaClass.
type MessageMediaClassArray []MessageMediaClass

// Sort sorts slice of MessageMediaClass.
func (s MessageMediaClassArray) Sort(less func(a, b MessageMediaClass) bool) MessageMediaClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessageMediaClass.
func (s MessageMediaClassArray) SortStable(less func(a, b MessageMediaClass) bool) MessageMediaClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessageMediaClass.
func (s MessageMediaClassArray) Retain(keep func(x MessageMediaClass) bool) MessageMediaClassArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s MessageMediaClassArray) First() (v MessageMediaClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessageMediaClassArray) Last() (v MessageMediaClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessageMediaClassArray) PopFirst() (v MessageMediaClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessageMediaClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessageMediaClassArray) Pop() (v MessageMediaClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsMessageMediaPhoto returns copy with only MessageMediaPhoto constructors.
func (s MessageMediaClassArray) AsMessageMediaPhoto() (to MessageMediaPhotoArray) {
	for _, elem := range s {
		value, ok := elem.(*MessageMediaPhoto)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsMessageMediaGeo returns copy with only MessageMediaGeo constructors.
func (s MessageMediaClassArray) AsMessageMediaGeo() (to MessageMediaGeoArray) {
	for _, elem := range s {
		value, ok := elem.(*MessageMediaGeo)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsMessageMediaContact returns copy with only MessageMediaContact constructors.
func (s MessageMediaClassArray) AsMessageMediaContact() (to MessageMediaContactArray) {
	for _, elem := range s {
		value, ok := elem.(*MessageMediaContact)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsMessageMediaDocument returns copy with only MessageMediaDocument constructors.
func (s MessageMediaClassArray) AsMessageMediaDocument() (to MessageMediaDocumentArray) {
	for _, elem := range s {
		value, ok := elem.(*MessageMediaDocument)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsMessageMediaWebPage returns copy with only MessageMediaWebPage constructors.
func (s MessageMediaClassArray) AsMessageMediaWebPage() (to MessageMediaWebPageArray) {
	for _, elem := range s {
		value, ok := elem.(*MessageMediaWebPage)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsMessageMediaVenue returns copy with only MessageMediaVenue constructors.
func (s MessageMediaClassArray) AsMessageMediaVenue() (to MessageMediaVenueArray) {
	for _, elem := range s {
		value, ok := elem.(*MessageMediaVenue)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsMessageMediaGame returns copy with only MessageMediaGame constructors.
func (s MessageMediaClassArray) AsMessageMediaGame() (to MessageMediaGameArray) {
	for _, elem := range s {
		value, ok := elem.(*MessageMediaGame)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsMessageMediaInvoice returns copy with only MessageMediaInvoice constructors.
func (s MessageMediaClassArray) AsMessageMediaInvoice() (to MessageMediaInvoiceArray) {
	for _, elem := range s {
		value, ok := elem.(*MessageMediaInvoice)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsMessageMediaGeoLive returns copy with only MessageMediaGeoLive constructors.
func (s MessageMediaClassArray) AsMessageMediaGeoLive() (to MessageMediaGeoLiveArray) {
	for _, elem := range s {
		value, ok := elem.(*MessageMediaGeoLive)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsMessageMediaPoll returns copy with only MessageMediaPoll constructors.
func (s MessageMediaClassArray) AsMessageMediaPoll() (to MessageMediaPollArray) {
	for _, elem := range s {
		value, ok := elem.(*MessageMediaPoll)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsMessageMediaDice returns copy with only MessageMediaDice constructors.
func (s MessageMediaClassArray) AsMessageMediaDice() (to MessageMediaDiceArray) {
	for _, elem := range s {
		value, ok := elem.(*MessageMediaDice)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsMessageMediaStory returns copy with only MessageMediaStory constructors.
func (s MessageMediaClassArray) AsMessageMediaStory() (to MessageMediaStoryArray) {
	for _, elem := range s {
		value, ok := elem.(*MessageMediaStory)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// MessageMediaPhotoArray is adapter for slice of MessageMediaPhoto.
type MessageMediaPhotoArray []MessageMediaPhoto

// Sort sorts slice of MessageMediaPhoto.
func (s MessageMediaPhotoArray) Sort(less func(a, b MessageMediaPhoto) bool) MessageMediaPhotoArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessageMediaPhoto.
func (s MessageMediaPhotoArray) SortStable(less func(a, b MessageMediaPhoto) bool) MessageMediaPhotoArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessageMediaPhoto.
func (s MessageMediaPhotoArray) Retain(keep func(x MessageMediaPhoto) bool) MessageMediaPhotoArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s MessageMediaPhotoArray) First() (v MessageMediaPhoto, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessageMediaPhotoArray) Last() (v MessageMediaPhoto, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessageMediaPhotoArray) PopFirst() (v MessageMediaPhoto, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessageMediaPhoto
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessageMediaPhotoArray) Pop() (v MessageMediaPhoto, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// MessageMediaGeoArray is adapter for slice of MessageMediaGeo.
type MessageMediaGeoArray []MessageMediaGeo

// Sort sorts slice of MessageMediaGeo.
func (s MessageMediaGeoArray) Sort(less func(a, b MessageMediaGeo) bool) MessageMediaGeoArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessageMediaGeo.
func (s MessageMediaGeoArray) SortStable(less func(a, b MessageMediaGeo) bool) MessageMediaGeoArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessageMediaGeo.
func (s MessageMediaGeoArray) Retain(keep func(x MessageMediaGeo) bool) MessageMediaGeoArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s MessageMediaGeoArray) First() (v MessageMediaGeo, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessageMediaGeoArray) Last() (v MessageMediaGeo, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessageMediaGeoArray) PopFirst() (v MessageMediaGeo, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessageMediaGeo
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessageMediaGeoArray) Pop() (v MessageMediaGeo, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// MessageMediaContactArray is adapter for slice of MessageMediaContact.
type MessageMediaContactArray []MessageMediaContact

// Sort sorts slice of MessageMediaContact.
func (s MessageMediaContactArray) Sort(less func(a, b MessageMediaContact) bool) MessageMediaContactArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessageMediaContact.
func (s MessageMediaContactArray) SortStable(less func(a, b MessageMediaContact) bool) MessageMediaContactArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessageMediaContact.
func (s MessageMediaContactArray) Retain(keep func(x MessageMediaContact) bool) MessageMediaContactArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s MessageMediaContactArray) First() (v MessageMediaContact, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessageMediaContactArray) Last() (v MessageMediaContact, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessageMediaContactArray) PopFirst() (v MessageMediaContact, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessageMediaContact
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessageMediaContactArray) Pop() (v MessageMediaContact, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// MessageMediaDocumentArray is adapter for slice of MessageMediaDocument.
type MessageMediaDocumentArray []MessageMediaDocument

// Sort sorts slice of MessageMediaDocument.
func (s MessageMediaDocumentArray) Sort(less func(a, b MessageMediaDocument) bool) MessageMediaDocumentArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessageMediaDocument.
func (s MessageMediaDocumentArray) SortStable(less func(a, b MessageMediaDocument) bool) MessageMediaDocumentArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessageMediaDocument.
func (s MessageMediaDocumentArray) Retain(keep func(x MessageMediaDocument) bool) MessageMediaDocumentArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s MessageMediaDocumentArray) First() (v MessageMediaDocument, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessageMediaDocumentArray) Last() (v MessageMediaDocument, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessageMediaDocumentArray) PopFirst() (v MessageMediaDocument, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessageMediaDocument
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessageMediaDocumentArray) Pop() (v MessageMediaDocument, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// MessageMediaWebPageArray is adapter for slice of MessageMediaWebPage.
type MessageMediaWebPageArray []MessageMediaWebPage

// Sort sorts slice of MessageMediaWebPage.
func (s MessageMediaWebPageArray) Sort(less func(a, b MessageMediaWebPage) bool) MessageMediaWebPageArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessageMediaWebPage.
func (s MessageMediaWebPageArray) SortStable(less func(a, b MessageMediaWebPage) bool) MessageMediaWebPageArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessageMediaWebPage.
func (s MessageMediaWebPageArray) Retain(keep func(x MessageMediaWebPage) bool) MessageMediaWebPageArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s MessageMediaWebPageArray) First() (v MessageMediaWebPage, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessageMediaWebPageArray) Last() (v MessageMediaWebPage, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessageMediaWebPageArray) PopFirst() (v MessageMediaWebPage, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessageMediaWebPage
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessageMediaWebPageArray) Pop() (v MessageMediaWebPage, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// MessageMediaVenueArray is adapter for slice of MessageMediaVenue.
type MessageMediaVenueArray []MessageMediaVenue

// Sort sorts slice of MessageMediaVenue.
func (s MessageMediaVenueArray) Sort(less func(a, b MessageMediaVenue) bool) MessageMediaVenueArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessageMediaVenue.
func (s MessageMediaVenueArray) SortStable(less func(a, b MessageMediaVenue) bool) MessageMediaVenueArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessageMediaVenue.
func (s MessageMediaVenueArray) Retain(keep func(x MessageMediaVenue) bool) MessageMediaVenueArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s MessageMediaVenueArray) First() (v MessageMediaVenue, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessageMediaVenueArray) Last() (v MessageMediaVenue, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessageMediaVenueArray) PopFirst() (v MessageMediaVenue, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessageMediaVenue
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessageMediaVenueArray) Pop() (v MessageMediaVenue, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// MessageMediaGameArray is adapter for slice of MessageMediaGame.
type MessageMediaGameArray []MessageMediaGame

// Sort sorts slice of MessageMediaGame.
func (s MessageMediaGameArray) Sort(less func(a, b MessageMediaGame) bool) MessageMediaGameArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessageMediaGame.
func (s MessageMediaGameArray) SortStable(less func(a, b MessageMediaGame) bool) MessageMediaGameArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessageMediaGame.
func (s MessageMediaGameArray) Retain(keep func(x MessageMediaGame) bool) MessageMediaGameArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s MessageMediaGameArray) First() (v MessageMediaGame, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessageMediaGameArray) Last() (v MessageMediaGame, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessageMediaGameArray) PopFirst() (v MessageMediaGame, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessageMediaGame
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessageMediaGameArray) Pop() (v MessageMediaGame, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// MessageMediaInvoiceArray is adapter for slice of MessageMediaInvoice.
type MessageMediaInvoiceArray []MessageMediaInvoice

// Sort sorts slice of MessageMediaInvoice.
func (s MessageMediaInvoiceArray) Sort(less func(a, b MessageMediaInvoice) bool) MessageMediaInvoiceArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessageMediaInvoice.
func (s MessageMediaInvoiceArray) SortStable(less func(a, b MessageMediaInvoice) bool) MessageMediaInvoiceArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessageMediaInvoice.
func (s MessageMediaInvoiceArray) Retain(keep func(x MessageMediaInvoice) bool) MessageMediaInvoiceArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s MessageMediaInvoiceArray) First() (v MessageMediaInvoice, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessageMediaInvoiceArray) Last() (v MessageMediaInvoice, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessageMediaInvoiceArray) PopFirst() (v MessageMediaInvoice, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessageMediaInvoice
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessageMediaInvoiceArray) Pop() (v MessageMediaInvoice, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// MessageMediaGeoLiveArray is adapter for slice of MessageMediaGeoLive.
type MessageMediaGeoLiveArray []MessageMediaGeoLive

// Sort sorts slice of MessageMediaGeoLive.
func (s MessageMediaGeoLiveArray) Sort(less func(a, b MessageMediaGeoLive) bool) MessageMediaGeoLiveArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessageMediaGeoLive.
func (s MessageMediaGeoLiveArray) SortStable(less func(a, b MessageMediaGeoLive) bool) MessageMediaGeoLiveArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessageMediaGeoLive.
func (s MessageMediaGeoLiveArray) Retain(keep func(x MessageMediaGeoLive) bool) MessageMediaGeoLiveArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s MessageMediaGeoLiveArray) First() (v MessageMediaGeoLive, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessageMediaGeoLiveArray) Last() (v MessageMediaGeoLive, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessageMediaGeoLiveArray) PopFirst() (v MessageMediaGeoLive, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessageMediaGeoLive
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessageMediaGeoLiveArray) Pop() (v MessageMediaGeoLive, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// MessageMediaPollArray is adapter for slice of MessageMediaPoll.
type MessageMediaPollArray []MessageMediaPoll

// Sort sorts slice of MessageMediaPoll.
func (s MessageMediaPollArray) Sort(less func(a, b MessageMediaPoll) bool) MessageMediaPollArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessageMediaPoll.
func (s MessageMediaPollArray) SortStable(less func(a, b MessageMediaPoll) bool) MessageMediaPollArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessageMediaPoll.
func (s MessageMediaPollArray) Retain(keep func(x MessageMediaPoll) bool) MessageMediaPollArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s MessageMediaPollArray) First() (v MessageMediaPoll, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessageMediaPollArray) Last() (v MessageMediaPoll, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessageMediaPollArray) PopFirst() (v MessageMediaPoll, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessageMediaPoll
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessageMediaPollArray) Pop() (v MessageMediaPoll, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// MessageMediaDiceArray is adapter for slice of MessageMediaDice.
type MessageMediaDiceArray []MessageMediaDice

// Sort sorts slice of MessageMediaDice.
func (s MessageMediaDiceArray) Sort(less func(a, b MessageMediaDice) bool) MessageMediaDiceArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessageMediaDice.
func (s MessageMediaDiceArray) SortStable(less func(a, b MessageMediaDice) bool) MessageMediaDiceArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessageMediaDice.
func (s MessageMediaDiceArray) Retain(keep func(x MessageMediaDice) bool) MessageMediaDiceArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s MessageMediaDiceArray) First() (v MessageMediaDice, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessageMediaDiceArray) Last() (v MessageMediaDice, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessageMediaDiceArray) PopFirst() (v MessageMediaDice, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessageMediaDice
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessageMediaDiceArray) Pop() (v MessageMediaDice, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// MessageMediaStoryArray is adapter for slice of MessageMediaStory.
type MessageMediaStoryArray []MessageMediaStory

// Sort sorts slice of MessageMediaStory.
func (s MessageMediaStoryArray) Sort(less func(a, b MessageMediaStory) bool) MessageMediaStoryArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessageMediaStory.
func (s MessageMediaStoryArray) SortStable(less func(a, b MessageMediaStory) bool) MessageMediaStoryArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessageMediaStory.
func (s MessageMediaStoryArray) Retain(keep func(x MessageMediaStory) bool) MessageMediaStoryArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s MessageMediaStoryArray) First() (v MessageMediaStory, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessageMediaStoryArray) Last() (v MessageMediaStory, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessageMediaStoryArray) PopFirst() (v MessageMediaStory, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessageMediaStory
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessageMediaStoryArray) Pop() (v MessageMediaStory, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of MessageMediaStory by ID.
func (s MessageMediaStoryArray) SortByID() MessageMediaStoryArray {
	return s.Sort(func(a, b MessageMediaStory) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of MessageMediaStory by ID.
func (s MessageMediaStoryArray) SortStableByID() MessageMediaStoryArray {
	return s.SortStable(func(a, b MessageMediaStory) bool {
		return a.GetID() < b.GetID()
	})
}

// FillMap fills constructors to given map.
func (s MessageMediaStoryArray) FillMap(to map[int]MessageMediaStory) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s MessageMediaStoryArray) ToMap() map[int]MessageMediaStory {
	r := make(map[int]MessageMediaStory, len(s))
	s.FillMap(r)
	return r
}
