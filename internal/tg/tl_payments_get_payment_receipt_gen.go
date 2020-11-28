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

// PaymentsGetPaymentReceiptRequest represents TL type `payments.getPaymentReceipt#a092a980`.
type PaymentsGetPaymentReceiptRequest struct {
	// MsgID field of PaymentsGetPaymentReceiptRequest.
	MsgID int
}

// PaymentsGetPaymentReceiptRequestTypeID is TL type id of PaymentsGetPaymentReceiptRequest.
const PaymentsGetPaymentReceiptRequestTypeID = 0xa092a980

// Encode implements bin.Encoder.
func (g *PaymentsGetPaymentReceiptRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode payments.getPaymentReceipt#a092a980 as nil")
	}
	b.PutID(PaymentsGetPaymentReceiptRequestTypeID)
	b.PutInt(g.MsgID)
	return nil
}

// Decode implements bin.Decoder.
func (g *PaymentsGetPaymentReceiptRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode payments.getPaymentReceipt#a092a980 to nil")
	}
	if err := b.ConsumeID(PaymentsGetPaymentReceiptRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.getPaymentReceipt#a092a980: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode payments.getPaymentReceipt#a092a980: field msg_id: %w", err)
		}
		g.MsgID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for PaymentsGetPaymentReceiptRequest.
var (
	_ bin.Encoder = &PaymentsGetPaymentReceiptRequest{}
	_ bin.Decoder = &PaymentsGetPaymentReceiptRequest{}
)
