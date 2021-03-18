package dwolla

import (
	"encoding/json"
	"log"
)

// MassPaymentStatus is a mass payment status
type MassPaymentStatus string

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

// MassPaymentItemStatus is a mass payment item status
type MassPaymentItemStatus string

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

//MassPaymentHandler is used to create/update mass payment requests
func MassPaymentHandler(paymentConfig *customer) *customer {
	return paymentConfig
}

func (p *massPayment) InitiateMassPayment(idempotencyKey string, massPaymentReq *MassPayment) (*MassPayment, *Raw, error) {
	url := p.baseURL + "/mass-payments"

	token, err := p.authHandler.GetToken()
	if err != nil {
		return nil, nil, err
	}

	header := &Header{
		IdempotencyKey: idempotencyKey,
	}

	resp, raw, err := post(url, header, massPaymentReq, token)
	if err != nil {
		return nil, raw, err
	}

	massPaymentLocation := resp.Header.Get(location)
	massPaymentID, err := ExtractIDFromLocation(massPaymentLocation)
	if err != nil {
		return nil, raw, err
	}

	massPaymentReq.Location = massPaymentLocation
	massPaymentReq.ID = massPaymentID

	return massPaymentReq, raw, nil
}

func (p *massPayment) GetMassPaymentByID(massPaymentID string) (*MassPaymentResponse, *Raw, error) {
	url := p.baseURL + "/mass-payments/" + massPaymentID

	token, err := p.authHandler.GetToken()
	if err != nil {
		return nil, nil, err
	}

	resp, raw, err := get(url, token)
	if err != nil {
		return nil, raw, err
	}

	var massPaymentResp MassPaymentResponse
	if err := json.Unmarshal(resp.Body, &massPaymentResp); err != nil {
		return nil, raw, err
	}

	return &massPaymentResp, raw, nil
}

func (p *massPayment) GetMassPaymentItemsByID(massPaymentID string) (*MassPaymentItems, *Raw, error) {
	url := p.baseURL + "/mass-payments/" + massPaymentID + "/items"

	token, err := p.authHandler.GetToken()
	if err != nil {
		return nil, nil, err
	}

	resp, raw, err := get(url, token)
	if err != nil {
		return nil, raw, err
	}

	var massPaymentItems MassPaymentItems
	if err := json.Unmarshal(resp.Body, &massPaymentItems); err != nil {
		return nil, raw, err
	}

	return &massPaymentItems, raw, nil
}

func (p *massPayment) UpdateMassPaymentStatus(massPaymentID string, status MassPaymentStatus) (*Raw, error) {
	url := p.baseURL + "/mass-payments/" + massPaymentID

	token, err := p.authHandler.GetToken()
	if err != nil {
		return nil, err
	}

	statusReq := &UpdateMassPayment{
		Status: status,
	}

	resp, raw, err := post(url, nil, statusReq, token)
	if err != nil {
		return raw, err
	}

	log.Println(string(resp.Body))
	return raw, nil

}
