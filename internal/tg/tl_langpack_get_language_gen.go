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

// LangpackGetLanguageRequest represents TL type `langpack.getLanguage#6a596502`.
type LangpackGetLanguageRequest struct {
	// LangPack field of LangpackGetLanguageRequest.
	LangPack string
	// LangCode field of LangpackGetLanguageRequest.
	LangCode string
}

// LangpackGetLanguageRequestTypeID is TL type id of LangpackGetLanguageRequest.
const LangpackGetLanguageRequestTypeID = 0x6a596502

// Encode implements bin.Encoder.
func (g *LangpackGetLanguageRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode langpack.getLanguage#6a596502 as nil")
	}
	b.PutID(LangpackGetLanguageRequestTypeID)
	b.PutString(g.LangPack)
	b.PutString(g.LangCode)
	return nil
}

// Decode implements bin.Decoder.
func (g *LangpackGetLanguageRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode langpack.getLanguage#6a596502 to nil")
	}
	if err := b.ConsumeID(LangpackGetLanguageRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode langpack.getLanguage#6a596502: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode langpack.getLanguage#6a596502: field lang_pack: %w", err)
		}
		g.LangPack = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode langpack.getLanguage#6a596502: field lang_code: %w", err)
		}
		g.LangCode = value
	}
	return nil
}

// Ensuring interfaces in compile-time for LangpackGetLanguageRequest.
var (
	_ bin.Encoder = &LangpackGetLanguageRequest{}
	_ bin.Decoder = &LangpackGetLanguageRequest{}
)
