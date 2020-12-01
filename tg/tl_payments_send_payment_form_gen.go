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

// PaymentsSendPaymentFormRequest represents TL type `payments.sendPaymentForm#2b8879b3`.
type PaymentsSendPaymentFormRequest struct {
	// Flags field of PaymentsSendPaymentFormRequest.
	Flags bin.Fields
	// MsgID field of PaymentsSendPaymentFormRequest.
	MsgID int
	// RequestedInfoID field of PaymentsSendPaymentFormRequest.
	//
	// Use SetRequestedInfoID and GetRequestedInfoID helpers.
	RequestedInfoID string
	// ShippingOptionID field of PaymentsSendPaymentFormRequest.
	//
	// Use SetShippingOptionID and GetShippingOptionID helpers.
	ShippingOptionID string
	// Credentials field of PaymentsSendPaymentFormRequest.
	Credentials InputPaymentCredentialsClass
}

// PaymentsSendPaymentFormRequestTypeID is TL type id of PaymentsSendPaymentFormRequest.
const PaymentsSendPaymentFormRequestTypeID = 0x2b8879b3

// Encode implements bin.Encoder.
func (s *PaymentsSendPaymentFormRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode payments.sendPaymentForm#2b8879b3 as nil")
	}
	b.PutID(PaymentsSendPaymentFormRequestTypeID)
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.sendPaymentForm#2b8879b3: field flags: %w", err)
	}
	b.PutInt(s.MsgID)
	if s.Flags.Has(0) {
		b.PutString(s.RequestedInfoID)
	}
	if s.Flags.Has(1) {
		b.PutString(s.ShippingOptionID)
	}
	if s.Credentials == nil {
		return fmt.Errorf("unable to encode payments.sendPaymentForm#2b8879b3: field credentials is nil")
	}
	if err := s.Credentials.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.sendPaymentForm#2b8879b3: field credentials: %w", err)
	}
	return nil
}

// SetRequestedInfoID sets value of RequestedInfoID conditional field.
func (s *PaymentsSendPaymentFormRequest) SetRequestedInfoID(value string) {
	s.Flags.Set(0)
	s.RequestedInfoID = value
}

// GetRequestedInfoID returns value of RequestedInfoID conditional field and
// boolean which is true if field was set.
func (s *PaymentsSendPaymentFormRequest) GetRequestedInfoID() (value string, ok bool) {
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.RequestedInfoID, true
}

// SetShippingOptionID sets value of ShippingOptionID conditional field.
func (s *PaymentsSendPaymentFormRequest) SetShippingOptionID(value string) {
	s.Flags.Set(1)
	s.ShippingOptionID = value
}

// GetShippingOptionID returns value of ShippingOptionID conditional field and
// boolean which is true if field was set.
func (s *PaymentsSendPaymentFormRequest) GetShippingOptionID() (value string, ok bool) {
	if !s.Flags.Has(1) {
		return value, false
	}
	return s.ShippingOptionID, true
}

// Decode implements bin.Decoder.
func (s *PaymentsSendPaymentFormRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode payments.sendPaymentForm#2b8879b3 to nil")
	}
	if err := b.ConsumeID(PaymentsSendPaymentFormRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.sendPaymentForm#2b8879b3: %w", err)
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.sendPaymentForm#2b8879b3: field flags: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode payments.sendPaymentForm#2b8879b3: field msg_id: %w", err)
		}
		s.MsgID = value
	}
	if s.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode payments.sendPaymentForm#2b8879b3: field requested_info_id: %w", err)
		}
		s.RequestedInfoID = value
	}
	if s.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode payments.sendPaymentForm#2b8879b3: field shipping_option_id: %w", err)
		}
		s.ShippingOptionID = value
	}
	{
		value, err := DecodeInputPaymentCredentials(b)
		if err != nil {
			return fmt.Errorf("unable to decode payments.sendPaymentForm#2b8879b3: field credentials: %w", err)
		}
		s.Credentials = value
	}
	return nil
}

// Ensuring interfaces in compile-time for PaymentsSendPaymentFormRequest.
var (
	_ bin.Encoder = &PaymentsSendPaymentFormRequest{}
	_ bin.Decoder = &PaymentsSendPaymentFormRequest{}
)

// PaymentsSendPaymentForm invokes method payments.sendPaymentForm#2b8879b3 returning error if any.
func (c *Client) PaymentsSendPaymentForm(ctx context.Context, request *PaymentsSendPaymentFormRequest) (PaymentsPaymentResultClass, error) {
	var result PaymentsPaymentResultBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.PaymentResult, nil
}
