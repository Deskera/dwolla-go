package dwolla

import (
	"encoding/json"
	"log"
)

const (
	// MassPaymentStatusDeferred is when a mass payment is deferred
	MassPaymentStatusDeferred MassPaymentStatus = "deferred"
	// MassPaymentStatusPending is when the mass payment is pending
	MassPaymentStatusPending MassPaymentStatus = "pending"
	// MassPaymentStatusProcessing is when the mass payment is processing
	MassPaymentStatusProcessing MassPaymentStatus = "processing"
	// MassPaymentStatusComplete is when the mass payment is complete
	MassPaymentStatusComplete MassPaymentStatus = "complete"
	// MassPaymentStatusCancelled is when the mass payment is cancelled
	MassPaymentStatusCancelled MassPaymentStatus = "cancelled"
)

const (
	// MassPaymentItemStatusPending is when a mass payment item is pending
	MassPaymentItemStatusPending MassPaymentItemStatus = "pending"
	// MassPaymentItemStatusSuccess is when amass payment item is successful
	MassPaymentItemStatusSuccess MassPaymentItemStatus = "success"
	// MassPaymentItemStatusFailed is when a mass payment item failed
	MassPaymentItemStatusFailed MassPaymentItemStatus = "failed"
)

type massPayment struct {
	authHandler *auth
	baseURL     string
}

// MassPaymentStatus is a mass payment status
type MassPaymentStatus string

// MassPaymentItemStatus is a mass payment item status
type MassPaymentItemStatus string

func MassPaymentHandler(paymentConfig *customer) *customer {
	return paymentConfig
}

func (p *massPayment) InitiateMassPayment(idempotencyKey string, massPaymentReq *MassPayment) (*MassPayment, error) {
	url := p.baseURL + "/mass-payments"

	token, err := p.authHandler.GetToken()
	if err != nil {
		return nil, err
	}

	header := &Header{
		IdempotencyKey: idempotencyKey,
	}

	resp, err := post(url, header, massPaymentReq, token)
	if err != nil {
		return nil, err
	}

	massPaymentLocation := resp.Header.Get(location)
	massPaymentReq.Location = massPaymentLocation

	return massPaymentReq, nil
}

func (p *massPayment) GetMassPaymentById(massPaymentLink string) (*MassPaymentResponse, error) {
	url := massPaymentLink

	token, err := p.authHandler.GetToken()
	if err != nil {
		return nil, err
	}

	resp, err := get(url, token)
	if err != nil {
		return nil, err
	}

	var massPaymentResp MassPaymentResponse
	if err := json.Unmarshal(resp.Body, &massPaymentResp); err != nil {
		return nil, err
	}

	return &massPaymentResp, nil
}

func (p *massPayment) UpdateMassPaymentStatus(massPaymentLink string, status MassPaymentStatus) error {
	url := massPaymentLink

	token, err := p.authHandler.GetToken()
	if err != nil {
		return err
	}

	statusReq := &UpdateMassPayment{
		Status: status,
	}

	resp, err := post(url, nil, statusReq, token)
	if err != nil {
		return err
	}

	log.Println(string(resp.Body))
	return nil

}
