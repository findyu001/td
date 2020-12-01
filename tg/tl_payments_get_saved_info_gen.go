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

// PaymentsGetSavedInfoRequest represents TL type `payments.getSavedInfo#227d824b`.
type PaymentsGetSavedInfoRequest struct {
}

// PaymentsGetSavedInfoRequestTypeID is TL type id of PaymentsGetSavedInfoRequest.
const PaymentsGetSavedInfoRequestTypeID = 0x227d824b

// Encode implements bin.Encoder.
func (g *PaymentsGetSavedInfoRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode payments.getSavedInfo#227d824b as nil")
	}
	b.PutID(PaymentsGetSavedInfoRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (g *PaymentsGetSavedInfoRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode payments.getSavedInfo#227d824b to nil")
	}
	if err := b.ConsumeID(PaymentsGetSavedInfoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.getSavedInfo#227d824b: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for PaymentsGetSavedInfoRequest.
var (
	_ bin.Encoder = &PaymentsGetSavedInfoRequest{}
	_ bin.Decoder = &PaymentsGetSavedInfoRequest{}
)

// PaymentsGetSavedInfo invokes method payments.getSavedInfo#227d824b returning error if any.
func (c *Client) PaymentsGetSavedInfo(ctx context.Context, request *PaymentsGetSavedInfoRequest) (*PaymentsSavedInfo, error) {
	var result PaymentsSavedInfo
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
